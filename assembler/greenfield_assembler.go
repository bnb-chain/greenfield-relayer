package assembler

import (
	"bytes"
	"encoding/hex"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
	"time"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/metric"
	"github.com/bnb-chain/greenfield-relayer/types"
	"github.com/bnb-chain/greenfield-relayer/util"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type AlertKey struct {
	channel types.ChannelId
	seq     uint64
}

type GreenfieldAssembler struct {
	mutex                          sync.RWMutex
	config                         *config.Config
	bscExecutor                    *executor.BSCExecutor
	greenfieldExecutor             *executor.GreenfieldExecutor
	daoManager                     *dao.DaoManager
	blsPubKey                      []byte
	inturnRelayerSequenceStatusMap map[types.ChannelId]*types.SequenceStatus // flag for in-turn relayer that if it has requested the sequence from chain during its interval
	relayerNonceStatus             *types.NonceStatus
	metricService                  *metric.MetricService
	alertSetMutex                  sync.RWMutex
	alertSet                       map[AlertKey]struct{}
}

func NewGreenfieldAssembler(cfg *config.Config, executor *executor.GreenfieldExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor,
	ms *metric.MetricService) *GreenfieldAssembler {
	channels := cfg.GreenfieldConfig.MonitorChannelList
	inturnRelayerSequenceStatusMap := make(map[types.ChannelId]*types.SequenceStatus)

	for _, c := range channels {
		inturnRelayerSequenceStatusMap[types.ChannelId(c)] = &types.SequenceStatus{}
	}
	return &GreenfieldAssembler{
		config:                         cfg,
		greenfieldExecutor:             executor,
		daoManager:                     dao,
		bscExecutor:                    bscExecutor,
		blsPubKey:                      executor.BlsPubKey,
		inturnRelayerSequenceStatusMap: inturnRelayerSequenceStatusMap,
		relayerNonceStatus:             &types.NonceStatus{},
		metricService:                  ms,
		alertSet:                       make(map[AlertKey]struct{}, 0),
	}
}

// AssembleTransactionsLoop assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *GreenfieldAssembler) AssembleTransactionsLoop() {
	ticker := time.NewTicker(common.AssembleInterval)
	for range ticker.C {
		inturnRelayer, err := a.bscExecutor.GetInturnRelayer()
		if err != nil {
			logging.Logger.Errorf("encounter error when retrieving in-turn relayer from chain, err=%s ", err.Error())
			continue
		}
		inturnRelayerPubkey, err := hex.DecodeString(inturnRelayer.BlsPublicKey)
		if err != nil {
			logging.Logger.Errorf("encounter error when decode in-turn relayer key, err=%s ", err.Error())
			continue
		}
		isInturnRelyer := bytes.Equal(a.blsPubKey, inturnRelayerPubkey)
		a.metricService.SetBSCInturnRelayerMetrics(isInturnRelyer, inturnRelayer.Start, inturnRelayer.End)

		if (isInturnRelyer && !a.relayerNonceStatus.HasRetrieved) || !isInturnRelyer {
			nonce, err := a.bscExecutor.GetNonce()
			if err != nil {
				logging.Logger.Errorf("encounter error when get relayer nonce, err=%s ", err.Error())
				continue
			}
			a.relayerNonceStatus.Nonce = nonce
		}

		wg := new(sync.WaitGroup)
		for _, c := range a.getMonitorChannels() {
			wg.Add(1)
			go a.assembleTransactionAndSendForChannel(types.ChannelId(c), inturnRelayer, isInturnRelyer, wg)
		}
		wg.Wait()
	}
}

func (a *GreenfieldAssembler) assembleTransactionAndSendForChannel(channelId types.ChannelId, inturnRelayer *types.InturnRelayer, isInturnRelyer bool, wg *sync.WaitGroup) {
	defer wg.Done()
	err := a.process(channelId, inturnRelayer, isInturnRelyer)
	if err != nil {
		logging.Logger.Errorf("encounter error, err=%s", err.Error())
	}
}

func (a *GreenfieldAssembler) process(channelId types.ChannelId, inturnRelayer *types.InturnRelayer, isInturnRelyer bool) error {
	var (
		startSeq    uint64
		endSequence int64
	)

	if isInturnRelyer {
		if !a.inturnRelayerSequenceStatusMap[channelId].HasRetrieved {
			now := time.Now().Unix()
			timeDiff := now - int64(inturnRelayer.Start)
			if timeDiff < a.config.RelayConfig.BSCSequenceUpdateLatency {
				if timeDiff < 0 {
					return fmt.Errorf("blockchain time and relayer time is not consistent, now %d should be after %d", now, inturnRelayer.Start)
				}
				return nil
			}
			inTurnRelayerStartSeq, err := a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
			if err != nil {
				return fmt.Errorf("faield to get next delivery sequence for channel %d, err=%s", channelId, err.Error())
			}
			a.mutex.Lock()
			a.inturnRelayerSequenceStatusMap[channelId].HasRetrieved = true
			a.inturnRelayerSequenceStatusMap[channelId].NextDeliverySeq = inTurnRelayerStartSeq
			a.mutex.Unlock()
		}
		startSeq = a.inturnRelayerSequenceStatusMap[channelId].NextDeliverySeq
	} else {
		a.mutex.Lock()
		a.inturnRelayerSequenceStatusMap[channelId].HasRetrieved = false
		a.mutex.Unlock()
		time.Sleep(time.Duration(a.config.RelayConfig.BSCSequenceUpdateLatency) * time.Second)
		var err error
		startSeq, err = a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
		if err != nil {
			return fmt.Errorf("faield to get next delivery sequence for channel %d, err=%s", channelId, err.Error())
		}
	}

	err := a.updateMetrics(channelId, startSeq)
	if err != nil {
		return err
	}

	if isInturnRelyer {
		endSequence, err = a.daoManager.GreenfieldDao.GetLatestSequenceByChannelIdAndStatus(channelId, db.AllVoted)
		if err != nil {
			return fmt.Errorf("faield to get latest sequence from DB, err=%s", err.Error())
		}
		if endSequence == -1 {
			return nil
		}
	} else {
		endSeq, err := a.greenfieldExecutor.GetNextSendSequenceForChannelWithRetry(a.getDestChainId(), channelId)
		if err != nil {
			return fmt.Errorf("faield to get next send sequence, err=%s", err.Error())
		}
		endSequence = int64(endSeq)
	}

	logging.Logger.Debugf("channel %d start seq and end enq are %d and %d", channelId, startSeq, endSequence)

	// if the start seq larger than the largest alerts' related tx's seq, then clear all alerts because tx are delivered
	a.alertSetMutex.Lock()
	if len(a.alertSet) > 0 {
		var maxTxSeqOfAlert uint64
		for k := range a.alertSet {
			if k.seq > maxTxSeqOfAlert {
				maxTxSeqOfAlert = k.seq
			}
		}
		if startSeq > maxTxSeqOfAlert {
			a.metricService.SetHasTxDelay(false)
			for k := range a.alertSet {
				if k.channel == channelId {
					delete(a.alertSet, k)
				}
			}
		}
	}
	a.alertSetMutex.Unlock()
	for i := startSeq; i <= uint64(endSequence); i++ {
		tx, err := a.daoManager.GreenfieldDao.GetTransactionByChannelIdAndSequence(channelId, i)
		if err != nil {
			return fmt.Errorf("faield to get transaction by cahnnel id %d and sequence %d from DB, err=%s", channelId, i, err.Error())
		}
		if (*tx == model.GreenfieldRelayTransaction{}) {
			return nil
		}

		if time.Since(time.Unix(tx.TxTime, 0)).Seconds() > common.TxDelayAlertThreshHold {
			a.metricService.SetHasTxDelay(true)
			key := AlertKey{
				channel: channelId,
				seq:     i,
			}
			a.alertSetMutex.Lock()
			a.alertSet[key] = struct{}{}
			a.alertSetMutex.Unlock()
		}

		if tx.Status != db.AllVoted && tx.Status != db.Delivered {
			return fmt.Errorf("tx with channel id %d and sequence %d does not get enough votes yet", tx.ChannelId, tx.Sequence)
		}
		if !isInturnRelyer && time.Now().Unix() < tx.TxTime+a.config.RelayConfig.GreenfieldToBSCInturnRelayerTimeout {
			return nil
		}
		if err := a.processTx(tx, a.relayerNonceStatus.Nonce, isInturnRelyer); err != nil {
			return err
		}
		logging.Logger.Infof("relayed tx with channel id %d and sequence %d ", tx.ChannelId, tx.Sequence)
		a.mutex.Lock()
		a.relayerNonceStatus.Nonce++
		a.mutex.Unlock()
	}
	return nil
}

func (a *GreenfieldAssembler) processTx(tx *model.GreenfieldRelayTransaction, nonce uint64, isInturnRelyer bool) error {
	// Get votes result for a tx, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(tx.ChannelId, tx.Sequence)
	if err != nil {
		return fmt.Errorf("failed to get votes for event with channel id %d and sequence %d", tx.ChannelId, tx.Sequence)
	}

	validators, err := a.bscExecutor.QueryCachedLatestValidators()
	if err != nil {
		return fmt.Errorf("failed to query cached validators, err=%s", err.Error())
	}
	aggregatedSignature, valBitSet, err := vote.AggregateSignatureAndValidatorBitSet(votes, validators)
	if err != nil {
		return fmt.Errorf("failed to aggregate signature, err=%s", err.Error())
	}

	txHash, err := a.bscExecutor.CallBuildInSystemContract(aggregatedSignature, util.BitSetToBigInt(valBitSet), votes[0].ClaimPayload, nonce)
	if err != nil {
		return fmt.Errorf("failed to submit tx to BSC, txHash=%s, err=%s", txHash, err.Error())
	}

	logging.Logger.Infof("relayed transaction with channel id %d and sequence %d, txHash=%s", tx.ChannelId, tx.Sequence, txHash)
	a.metricService.SetGnfdProcessedBlockHeight(tx.Height)

	// update next delivery sequence in DB for inturn relayer, for non-inturn relayer, there is enough time for
	// sequence update, so they can track next start seq from chain
	if !isInturnRelyer {
		if err = a.daoManager.GreenfieldDao.UpdateTransactionClaimedTxHash(tx.Id, txHash.String()); err != nil {
			return err
		}
		return fmt.Errorf("failed to update transaciton status, err=%s", err.Error())
	}

	if err = a.daoManager.GreenfieldDao.UpdateTransactionStatusAndClaimedTxHash(tx.Id, db.Delivered, txHash.String()); err != nil {
		return fmt.Errorf("failed to update transaciton status, err=%s", err.Error())
	}
	a.mutex.Lock()
	a.inturnRelayerSequenceStatusMap[types.ChannelId(tx.ChannelId)].NextDeliverySeq = tx.Sequence + 1
	a.mutex.Unlock()
	return nil
}

func (a *GreenfieldAssembler) getMonitorChannels() []uint8 {
	return a.config.GreenfieldConfig.MonitorChannelList
}

func (a *GreenfieldAssembler) updateMetrics(channelId types.ChannelId, nextDeliverySeq uint64) error {
	a.metricService.SetNextReceiveSequenceForChannel(uint8(channelId), nextDeliverySeq)
	nextSendSeq, err := a.greenfieldExecutor.GetNextSendSequenceForChannelWithRetry(a.getDestChainId(), channelId)
	if err != nil {
		return err
	}
	a.metricService.SetNextSendSequenceForChannel(uint8(channelId), nextSendSeq)
	return nil
}

func (a *GreenfieldAssembler) getDestChainId() sdk.ChainID {
	return sdk.ChainID(a.config.BSCConfig.ChainId)
}
