package assembler

import (
	"encoding/hex"
	"time"

	"github.com/bnb-chain/inscription-relayer/db/model"

	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type InscriptionAssembler struct {
	config              *config.Config
	bscExecutor         *executor.BSCExecutor
	inscriptionExecutor *executor.InscriptionExecutor
	daoManager          *dao.DaoManager
	votePoolExecutor    *vote.VotePoolExecutor
}

func NewInscriptionAssembler(cfg *config.Config, executor *executor.InscriptionExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor, votePoolExecutor *vote.VotePoolExecutor) *InscriptionAssembler {
	return &InscriptionAssembler{
		config:              cfg,
		inscriptionExecutor: executor,
		daoManager:          dao,
		bscExecutor:         bscExecutor,
		votePoolExecutor:    votePoolExecutor,
	}
}

// AssembleTransactionAndSend assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *InscriptionAssembler) AssembleTransactionAndSend() {
	for _, c := range a.getMonitorChannels() {
		go a.assembleTransactionAndSendForChannel(common.ChannelId(c))
	}
}

func (a *InscriptionAssembler) assembleTransactionAndSendForChannel(channelId common.ChannelId) {
	for {
		err := a.process(channelId)
		if err != nil {
			common.Logger.Errorf("encounter error when relaying tx, err=%s ", err.Error())
			time.Sleep(common.RetryInterval)
		}
	}
}

func (a *InscriptionAssembler) process(channelId common.ChannelId) error {
	nextSequence, err := a.inscriptionExecutor.GetNextDeliverySequenceForChannel(channelId)
	if err != nil {
		return err
	}

	tx, err := a.daoManager.InscriptionDao.GetTransactionByChannelIdAndSequenceAndStatus(channelId, nextSequence, db.AllVoted)
	if err != nil {
		return err
	}
	if (*tx == model.InscriptionRelayTransaction{}) {
		return nil
	}
	// Get votes result for a tx, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(tx.ChannelId, tx.Sequence)
	if err != nil {
		common.Logger.Errorf("failed to get votes for event with channel id %d and sequence %d", tx.ChannelId, tx.Sequence)
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

	relayerPubKey := util.GetBlsPubKeyFromPrivKeyStr(a.votePoolExecutor.GetBlsPrivateKey())
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerBlsPubKeys)
	firstInturnRelayerIdx := int(tx.TxTime) % len(relayerBlsPubKeys)
	txRelayStartTime := tx.TxTime + InscriptionRelayingDelayTime
	common.Logger.Infof("tx will be relayed starting at %d", txRelayStartTime)

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
		curRelayerRelayingStartTime = txRelayStartTime + FirstInTurnRelayerRelayingWindow + int64(indexDiff-1)*InTurnRelayerRelayingWindow
	}
	common.Logger.Infof("current relayer starts relaying from %d", curRelayerRelayingStartTime)

	filled := make(chan struct{})
	errC := make(chan error)
	go a.validateSequenceFilled(filled, errC, nextSequence, channelId)

	for {
		select {
		case err = <-errC:
			return err
		case <-filled:
			if err = a.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, db.Delivered); err != nil {
				common.Logger.Errorf("failed to update tx status %s", tx)
				return err
			}
			return nil
		default:
			if time.Now().Unix() >= curRelayerRelayingStartTime {
				common.Logger.Infof("relaying transaction with channel id %d and sequence %d", tx.ChannelId, tx.Sequence)
				txHash, err := a.bscExecutor.CallBuildInSystemContract(aggregatedSignature, util.BitSetToBigInt(valBitSet), votes[0].ClaimPayload)
				if err != nil {
					return err
				}
				common.Logger.Infof("delivered transaction to BSC with txHash %s", txHash.String())
				err = a.daoManager.InscriptionDao.UpdateTransactionStatusAndClaimedTxHash(tx.Id, db.Delivered, txHash.String())
				if err != nil {
					common.Logger.Errorf("failed to update Tx with channel id %d and sequence %d to status 'filled'", tx.ChannelId, tx.Sequence)
					return err
				}
				return nil
			}
		}
	}
}

func (a *InscriptionAssembler) validateSequenceFilled(filled chan struct{}, errC chan error, sequence uint64, channelID common.ChannelId) {
	ticker := time.NewTicker(common.RetryInterval)
	defer ticker.Stop()
	for {
		nextDeliverySequence, err := a.inscriptionExecutor.GetNextDeliverySequenceForChannel(channelID)
		if err != nil {
			errC <- err
		}
		if sequence < nextDeliverySequence {
			common.Logger.Infof("sequence %d for channel %d has already been filled ", sequence, channelID)
			filled <- struct{}{}
		}
		<-ticker.C
	}
}

func (a *InscriptionAssembler) getMonitorChannels() []uint8 {
	return a.config.InscriptionConfig.MonitorChannelList
}
