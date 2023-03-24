package assembler

import (
	"bytes"
	"encoding/hex"
	"fmt"
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

type GreenfieldAssembler struct {
	mutex                            sync.RWMutex
	config                           *config.Config
	bscExecutor                      *executor.BSCExecutor
	greenfieldExecutor               *executor.GreenfieldExecutor
	daoManager                       *dao.DaoManager
	blsPubKey                        []byte
	hasRetrievedSequenceByChannelMap map[types.ChannelId]bool // flag for in-turn relayer that if it has requested the sequence from chain during its interval
	metricService                    *metric.MetricService
}

func NewGreenfieldAssembler(cfg *config.Config, executor *executor.GreenfieldExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor,
	ms *metric.MetricService) *GreenfieldAssembler {
	channels := cfg.GreenfieldConfig.MonitorChannelList
	retrievedSequenceByChannelMap := make(map[types.ChannelId]bool)
	for _, c := range channels {
		retrievedSequenceByChannelMap[types.ChannelId(c)] = false
	}
	return &GreenfieldAssembler{
		config:                           cfg,
		greenfieldExecutor:               executor,
		daoManager:                       dao,
		bscExecutor:                      bscExecutor,
		blsPubKey:                        util.BlsPubKeyFromPrivKeyStr(cfg.GreenfieldConfig.BlsPrivateKey),
		hasRetrievedSequenceByChannelMap: retrievedSequenceByChannelMap,
		metricService:                    ms,
	}
}

// AssembleTransactionsLoop assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *GreenfieldAssembler) AssembleTransactionsLoop() {
	ticker := time.NewTicker(common.RetryInterval)
	for range ticker.C {
		inturnRelayer, err := a.bscExecutor.GetInturnRelayer()
		if err != nil {
			logging.Logger.Errorf("encounter error when retrieving in-turn relayer from chain, err=%s ", err.Error())
			continue
		}
		wg := new(sync.WaitGroup)
		errChan := make(chan error)
		for _, c := range a.getMonitorChannels() {
			wg.Add(1)
			go a.assembleTransactionAndSendForChannel(types.ChannelId(c), inturnRelayer, errChan, wg)
		}
		wg.Wait()
	}
}

func (a *GreenfieldAssembler) assembleTransactionAndSendForChannel(channelId types.ChannelId, inturnRelayer *types.InturnRelayer, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	err := a.process(channelId, inturnRelayer)
	if err != nil {
		errChan <- err
	}
}

func (a *GreenfieldAssembler) process(channelId types.ChannelId, inturnRelayer *types.InturnRelayer) error {
	var startSequence uint64
	inturnRelayerPubkey, err := hex.DecodeString(inturnRelayer.BlsPublicKey)
	if err != nil {
		return err
	}
	isInturnRelyer := bytes.Equal(a.blsPubKey, inturnRelayerPubkey)
	a.metricService.SetBSCInturnRelayerMetrics(isInturnRelyer, inturnRelayer.Start, inturnRelayer.End)
	if isInturnRelyer {
		// get next delivered sequence from DB
		seq, err := a.daoManager.SequenceDao.GetByChannelId(uint8(channelId))
		if err != nil {
			return err
		}
		startSequence = uint64(seq.Sequence)
		// in-turn relayer get the start sequence from chain once during its interval
		if !a.hasRetrievedSequenceByChannelMap[channelId] {
			now := time.Now().Unix()
			timeDiff := now - int64(inturnRelayer.Start)
			if timeDiff < a.config.RelayConfig.BSCSequenceUpdateLatency {
				if timeDiff < 0 {
					return fmt.Errorf("blockchain time and relayer time is not consistent, now %d should be after %d", now, inturnRelayer.Start)
				}
				return nil
			}
			startSequence, err = a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
			if err != nil {
				return err
			}
			if err = a.daoManager.SequenceDao.Upsert(uint8(channelId), startSequence); err != nil {
				return err
			}
			a.mutex.Lock()
			a.hasRetrievedSequenceByChannelMap[channelId] = true
			a.mutex.Unlock()
		}
		a.metricService.SetNextSequenceForChannelFromDB(uint8(channelId), startSequence)
		seqFromChain, err := a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
		if err != nil {
			return err
		}
		a.metricService.SetNextSequenceForChannelFromChain(uint8(channelId), seqFromChain)
	} else {
		a.mutex.Lock()
		a.hasRetrievedSequenceByChannelMap[channelId] = false
		a.mutex.Unlock()

		time.Sleep(time.Duration(a.config.RelayConfig.BSCSequenceUpdateLatency) * time.Second)
		startSequence, err := a.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(channelId)
		if err != nil {
			return err
		}
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
	a.metricService.SetGnfdProcessedBlockHeight(tx.Height)

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
