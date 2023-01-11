package assembler

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	"github.com/bnb-chain/inscription-relayer/vote"
	ethcommon "github.com/ethereum/go-ethereum/common"
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
			time.Sleep(RetryInterval)
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
		common.Logger.Errorf("failed to get AllVoted tx with channel id %d and sequence : %d", channelId, nextSequence)
		return err
	}
	if (*tx == model.InscriptionRelayTransaction{}) {
		return nil
	}
	// Get votes result for a tx, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(tx.ChannelId, tx.Sequence)
	if err != nil {
		common.Logger.Errorf("failed to get votes result for tx : %s", tx.TxHash)
		return err
	}

	validators, err := a.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return err
	}
	aggregatedSignature, valBitSet, err := vote.AggregateSignatureAndValidatorBitSet(votes, validators)
	if err != nil {
		return err
	}

	relayerBlsPubKeys, err := a.votePoolExecutor.GetValidatorsBlsPublicKey()
	if err != nil {
		return err
	}

	relayerPubKey := util.GetBlsPubKeyFromPrivKeyStr(a.getBlsPrivateKey())
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerBlsPubKeys)
	inturnRelayerIdx := int(tx.TxTime) % len(relayerBlsPubKeys)
	inturnRelayerRelayingTime := tx.TxTime + RelayWindowInSecond
	common.Logger.Infof("In-turn relayer relaying time is %d", inturnRelayerRelayingTime)

	var indexDiff int
	if relayerIdx >= inturnRelayerIdx {
		indexDiff = relayerIdx - inturnRelayerIdx
	} else {
		indexDiff = len(relayerBlsPubKeys) - (inturnRelayerIdx - relayerIdx)
	}
	curRelayerRelayingStartTime := inturnRelayerRelayingTime + int64(indexDiff*RelayIntervalBetweenRelayersInSecond)
	common.Logger.Infof("current relayer starts relaying from %d", curRelayerRelayingStartTime)

	// Keep pooling the next delivery sequence from dest chain until relaying time meets, or interrupt when seq is filled
	filled := make(chan struct{})
	defer close(filled)
	errC := make(chan error)
	defer close(errC)
	go a.validateSequenceFilled(filled, errC, nextSequence, channelId)

	for {
		select {
		case err = <-errC:
			return err
		case <-filled:
			if err = a.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, db.Filled); err != nil {
				common.Logger.Errorf("failed to update tx status %s", tx)
				return err
			}
			return nil
		default:
			if time.Now().Unix() >= curRelayerRelayingStartTime {
				common.Logger.Infof("relaying transaction %s", tx.TxHash)
				nonce, err := a.bscExecutor.GetClient().PendingNonceAt(context.Background(), a.bscExecutor.TxSender)
				if err != nil {
					return err
				}
				txHash, err := a.bscExecutor.CallBuildInSystemContract(int8(channelId), aggregatedSignature, nextSequence, util.BitSetToBigInt(valBitSet), ethcommon.Hex2Bytes(tx.PayLoad), nonce)
				if err != nil {
					return err
				}
				common.Logger.Infof("delivery transaction to BSC with txHash %s", txHash.String())
				err = a.daoManager.InscriptionDao.UpdateTransactionStatusAndClaimTxHash(tx.Id, db.Filled, txHash.String())
				if err != nil {
					common.Logger.Errorf("failed to update Tx status %d", tx.Id)
					return err
				}
			}
		}
	}
}

func (a *InscriptionAssembler) validateSequenceFilled(filled chan struct{}, errC chan error, sequence uint64, channelID common.ChannelId) {
	ticker := time.NewTicker(RetryInterval)
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

func (a *InscriptionAssembler) getBlsPrivateKey() string {
	return a.config.VotePoolConfig.BlsPrivateKey
}

func (a *InscriptionAssembler) getMonitorChannels() []uint8 {
	return a.config.InscriptionConfig.MonitorChannelList
}
