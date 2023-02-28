package assembler

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/types"
	"github.com/bnb-chain/greenfield-relayer/util"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type GreenfieldAssembler struct {
	config             *config.Config
	bscExecutor        *executor.BSCExecutor
	greenfieldExecutor *executor.GreenfieldExecutor
	daoManager         *dao.DaoManager
}

func NewGreenfieldAssembler(cfg *config.Config, executor *executor.GreenfieldExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor) *GreenfieldAssembler {
	return &GreenfieldAssembler{
		config:             cfg,
		greenfieldExecutor: executor,
		daoManager:         dao,
		bscExecutor:        bscExecutor,
	}
}

// AssembleTransactionsLoop assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *GreenfieldAssembler) AssembleTransactionsLoop() {
	for _, c := range a.getMonitorChannels() {
		go a.assembleTransactionAndSendForChannel(types.ChannelId(c))
	}
}

func (a *GreenfieldAssembler) assembleTransactionAndSendForChannel(channelId types.ChannelId) {
	for {
		if err := a.process(channelId); err != nil {
			logging.Logger.Errorf("encounter error when relaying tx, err=%s ", err.Error())
			time.Sleep(common.RetryInterval)
		}
	}
}

func (a *GreenfieldAssembler) process(channelId types.ChannelId) error {
	nextSequence, err := a.greenfieldExecutor.GetNextDeliverySequenceForChannel(channelId)
	if err != nil {
		return err
	}

	tx, err := a.daoManager.GreenfieldDao.GetTransactionByChannelIdAndSequence(channelId, nextSequence)
	if err != nil {
		return err
	}
	if (*tx == model.GreenfieldRelayTransaction{}) {
		return nil
	}
	if tx.Status != db.AllVoted && tx.Status != db.Delivered {
		return nil
	}
	// Get votes result for a tx, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(tx.ChannelId, tx.Sequence)
	if err != nil {
		logging.Logger.Errorf("failed to get votes for event with channel id %d and sequence %d", tx.ChannelId, tx.Sequence)
		return err
	}
	validators, err := a.bscExecutor.QueryCachedLatestValidators()
	if err != nil {
		return err
	}
	aggregatedSignature, valBitSet, err := vote.AggregateSignatureAndValidatorBitSet(votes, validators)
	if err != nil {
		return err
	}

	relayerBlsPubKeys, err := a.bscExecutor.GetValidatorsBlsPublicKey()
	if err != nil {
		return err
	}

	relayerPubKey := util.BlsPubKeyFromPrivKeyStr(a.greenfieldExecutor.GetBlsPrivateKey())
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerBlsPubKeys)
	if relayerIdx == -1 {
		return errors.New(" not a relayer. ")
	}

	firstInturnRelayerIdx := int(tx.TxTime) % len(relayerBlsPubKeys)
	txRelayStartTime := tx.TxTime + a.config.RelayConfig.GreenfieldToBSCRelayingDelayTime
	logging.Logger.Infof("tx will be relayed starting at %d", txRelayStartTime)

	var indexDiff int
	if relayerIdx >= firstInturnRelayerIdx {
		indexDiff = relayerIdx - firstInturnRelayerIdx
	} else {
		indexDiff = len(relayerBlsPubKeys) - (firstInturnRelayerIdx - relayerIdx)
	}
	curRelayerRelayingStartTime := int64(0)
	if indexDiff == 0 {
		curRelayerRelayingStartTime = txRelayStartTime
	} else {
		curRelayerRelayingStartTime = txRelayStartTime + a.config.RelayConfig.FirstInTurnRelayerRelayingWindow + int64(indexDiff-1)*a.config.RelayConfig.InTurnRelayerRelayingWindow
	}
	logging.Logger.Infof("current relayer starts relaying from %d", curRelayerRelayingStartTime)

	filled := make(chan struct{})
	errC := make(chan error)
	ticker := time.NewTicker(common.RetryInterval)

	go a.validateSequenceFilled(filled, errC, nextSequence, channelId)

	for {
		select {
		case err = <-errC:
			return err
		case <-filled:
			if err = a.daoManager.GreenfieldDao.UpdateTransactionStatus(tx.Id, db.Delivered); err != nil {
				logging.Logger.Errorf("failed to update Tx with channel id %d and sequence %d to status 'Delivered', error=%s", tx.ChannelId, tx.Sequence, err.Error())
				return err
			}
			return nil
		case <-ticker.C:
			if time.Now().Unix() >= curRelayerRelayingStartTime {
				logging.Logger.Infof("relaying transaction with channel id %d and sequence %d", tx.ChannelId, tx.Sequence)
				txHash, err := a.bscExecutor.CallBuildInSystemContract(aggregatedSignature, util.BitSetToBigInt(valBitSet), votes[0].ClaimPayload)
				if err != nil {
					return err
				}
				logging.Logger.Infof("delivered transaction to BSC with txHash %s", txHash.String())

				// `Delivered` does not mean tx is successful even there is txHash returned, so need to wait a bit and use sequence to validate, otherwise retry
				err = a.daoManager.GreenfieldDao.UpdateTransactionStatusAndClaimedTxHash(tx.Id, db.Delivered, txHash.String())
				if err != nil {
					logging.Logger.Errorf("failed to update Tx with channel id %d and sequence %d to status 'Delivered', error=%s", tx.ChannelId, tx.Sequence, err.Error())
					return err
				}
				time.Sleep(executor.BSCSequenceUpdateLatency)
				return nil
			}
		}
	}
}

func (a *GreenfieldAssembler) validateSequenceFilled(filled chan struct{}, errC chan error, sequence uint64, channelID types.ChannelId) {
	ticker := time.NewTicker(common.RetryInterval)
	defer ticker.Stop()
	for range ticker.C {
		nextDeliverySequence, err := a.greenfieldExecutor.GetNextDeliverySequenceForChannel(channelID)
		if err != nil {
			errC <- err
		}
		if sequence < nextDeliverySequence {
			logging.Logger.Infof("sequence %d for channel %d has already been filled ", sequence, channelID)
			filled <- struct{}{}
		}
	}
}

func (a *GreenfieldAssembler) getMonitorChannels() []uint8 {
	return a.config.GreenfieldConfig.MonitorChannelList
}
