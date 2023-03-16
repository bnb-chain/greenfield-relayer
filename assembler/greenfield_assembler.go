package assembler

import (
	"encoding/hex"
	"fmt"
	"github.com/bnb-chain/greenfield-relayer/common"
	"time"

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
	blsPubKey          string
}

func NewGreenfieldAssembler(cfg *config.Config, executor *executor.GreenfieldExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor) *GreenfieldAssembler {
	return &GreenfieldAssembler{
		config:             cfg,
		greenfieldExecutor: executor,
		daoManager:         dao,
		bscExecutor:        bscExecutor,
		blsPubKey:          hex.EncodeToString(util.BlsPubKeyFromPrivKeyStr(cfg.GreenfieldConfig.BlsPrivateKey)),
	}
}

// AssembleTransactionsLoop assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *GreenfieldAssembler) AssembleTransactionsLoop() {
	for _, c := range a.getMonitorChannels() {
		go a.assembleTransactionAndSendForChannel(types.ChannelId(c))
	}
}

func (a *GreenfieldAssembler) assembleTransactionAndSendForChannel(channelId types.ChannelId) {
	ticker := time.NewTicker(common.RetryInterval)
	for range ticker.C {
		if err := a.process(channelId); err != nil {
			logging.Logger.Errorf("encounter error when relaying tx, err=%s ", err.Error())
		}
	}
}

func (a *GreenfieldAssembler) process(channelId types.ChannelId) error {
	logging.Logger.Infof("current time is %d", time.Now().Unix())
	inturnRelayer, err := a.bscExecutor.GetInturnRelayer()
	if err != nil {
		return err
	}
	isInturnRelyer := inturnRelayer.BlsPublicKey == a.blsPubKey
	var startSequence uint64
	if isInturnRelyer {
		// get next delivered sequence from DB
		seq, err := a.daoManager.SequenceDao.GetByChannelId(uint8(channelId))
		if err != nil {
			return err
		}
		startSequence = uint64(seq.Sequence)

		// in-turn relayer get the start sequence from chain first time, it starts to relay after the  sequence
		// get updated
		now := time.Now().Unix()
		timeDiff := now - int64(inturnRelayer.Start)

		if timeDiff < a.config.RelayConfig.BSCSequenceUpdateLatency {
			if timeDiff < 0 {
				return fmt.Errorf("blockchain time and relayer time is not consistent, now %d should be after %d", now, inturnRelayer.Start)
			}
			time.Sleep(time.Duration(timeDiff) * time.Second)
			startSequence, err = a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
			if err != nil {
				return err
			}
			if err = a.daoManager.SequenceDao.Upsert(uint8(channelId), startSequence); err != nil {
				return err
			}
		}
		logging.Logger.Debug("gnfd relay as in-turn relayer")
	} else {
		time.Sleep(time.Duration(a.config.RelayConfig.BSCSequenceUpdateLatency) * time.Second)
		startSequence, err = a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
		if err != nil {
			return err
		}
		logging.Logger.Debug("gnfd relay as out-turn relayer")
		if err := a.daoManager.GreenfieldDao.UpdateBatchTransactionStatusToDelivered(startSequence); err != nil {
			return err
		}
	}

	endSequence, err := a.daoManager.GreenfieldDao.GetLatestSequenceByChannelIdAndStatus(channelId, db.AllVoted)
	if err != nil {
		return err
	}
	if endSequence == -1 {
		return nil
	}
	logging.Logger.Debugf("channel %d start seq is %d, end seq is %d ", channelId, startSequence, endSequence)
	nonce, err := a.bscExecutor.GetNonce()
	if err != nil {
		return err
	}

	for i := startSequence; i <= uint64(endSequence); i++ {
		tx, err := a.daoManager.GreenfieldDao.GetTransactionByChannelIdAndSequence(channelId, i)
		if err != nil {
			return err
		}
		if (*tx == model.GreenfieldRelayTransaction{}) {
			return nil
		}
		if tx.Status != db.AllVoted && tx.Status != db.Delivered {
			return fmt.Errorf("tx with channel id %d and sequence %d does not get enough votes yet", tx.ChannelId, tx.Sequence)
		}

		if !isInturnRelyer && time.Now().Unix() < tx.TxTime+a.config.RelayConfig.GreenfieldToBSCInturnRelayerTimeout {
			return nil
		}

		if err := a.processTx(tx, nonce, isInturnRelyer); err != nil {
			return err
		}
		logging.Logger.Infof("relayed tx with channel id %d and sequence %d ", tx.ChannelId, tx.Sequence)
		nonce++
	}
	return nil
}

func (a *GreenfieldAssembler) processTx(tx *model.GreenfieldRelayTransaction, nonce uint64, isInturnRelyer bool) error {
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

	txHash, err := a.bscExecutor.CallBuildInSystemContract(aggregatedSignature, util.BitSetToBigInt(valBitSet), votes[0].ClaimPayload, nonce)
	if err != nil {
		return err
	}
	logging.Logger.Infof("relayed transaction with channel id %d and sequence %d, get txHash %s", tx.ChannelId, tx.Sequence, txHash)

	// update next delivery sequence in DB for inturn relayer, for non-inturn relayer, there is enough time for
	// sequence update, so they can track next start seq from chain
	if !isInturnRelyer {
		if err = a.daoManager.GreenfieldDao.UpdateTransactionClaimedTxHash(tx.Id, txHash.String()); err != nil {
			return err
		}
		return nil
	}

	if err = a.daoManager.GreenfieldDao.UpdateTransactionStatusAndClaimedTxHash(tx.Id, db.Delivered, txHash.String()); err != nil {
		return err
	}
	if err = a.daoManager.SequenceDao.Upsert(tx.ChannelId, tx.Sequence+1); err != nil {
		return err
	}

	return nil
}

func (a *GreenfieldAssembler) getMonitorChannels() []uint8 {
	return a.config.GreenfieldConfig.MonitorChannelList
}
