package assembler

import (
	"bytes"
	"encoding/hex"
	"fmt"
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
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type BSCAssembler struct {
	config               *config.Config
	greenfieldExecutor   *executor.GreenfieldExecutor
	bscExecutor          *executor.BSCExecutor
	daoManager           *dao.DaoManager
	blsPubKey            []byte
	hasRetrievedSequence bool
	metricService        *metric.MetricService
}

func NewBSCAssembler(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager, greenfieldExecutor *executor.GreenfieldExecutor, ms *metric.MetricService) *BSCAssembler {
	return &BSCAssembler{
		config:             cfg,
		bscExecutor:        executor,
		daoManager:         dao,
		greenfieldExecutor: greenfieldExecutor,
		blsPubKey:          greenfieldExecutor.BlsPubKey,
		metricService:      ms,
	}
}

// AssemblePackagesAndClaimLoop assemble packages and then claim in Greenfield
func (a *BSCAssembler) AssemblePackagesAndClaimLoop() {
	a.assemblePackagesAndClaimForOracleChannel(common.OracleChannelId)
}

func (a *BSCAssembler) assemblePackagesAndClaimForOracleChannel(channelId types.ChannelId) {
	ticker := time.NewTicker(common.AssembleInterval)
	for range ticker.C {
		if err := a.process(channelId); err != nil {
			logging.Logger.Errorf("encounter error when relaying packages, err=%s ", err.Error())
		}
	}
}

func (a *BSCAssembler) process(channelId types.ChannelId) error {
	inturnRelayer, err := a.greenfieldExecutor.GetInturnRelayer()
	if err != nil {
		return err
	}
	inturnRelayerPubkey, err := hex.DecodeString(inturnRelayer.BlsPubKey)
	if err != nil {
		return err
	}
	isInturnRelyer := bytes.Equal(a.blsPubKey, inturnRelayerPubkey)
	a.metricService.SetGnfdInturnRelayerMetrics(isInturnRelyer, inturnRelayer.RelayInterval.Start, inturnRelayer.RelayInterval.End)
	var startSequence uint64

	if isInturnRelyer {
		seq, err := a.daoManager.SequenceDao.GetByChannelId(uint8(channelId))
		if err != nil {
			return err
		}
		startSequence = uint64(seq.Sequence)

		if !a.hasRetrievedSequence {
			// in-turn relayer get the start sequence from chain first time, it starts to relay after the sequence gets updated
			now := time.Now().Unix()
			timeDiff := now - int64(inturnRelayer.RelayInterval.Start)

			if timeDiff < a.config.RelayConfig.GreenfieldSequenceUpdateLatency {
				if timeDiff < 0 {
					return fmt.Errorf("blockchain time and relayer time is not consistent, now %d should be after %d", now, inturnRelayer.RelayInterval.Start)
				}
				return nil
			}
			startSequence, err = a.bscExecutor.GetNextDeliveryOracleSequenceWithRetry()
			if err != nil {
				return err
			}
			if err = a.daoManager.SequenceDao.Upsert(uint8(channelId), startSequence); err != nil {
				return err
			}
		}
		a.metricService.SetNextSequenceForChannelFromDB(uint8(channelId), startSequence)
		seqFromChain, err := a.bscExecutor.GetNextDeliveryOracleSequenceWithRetry()
		if err != nil {
			return err
		}
		a.metricService.SetNextSequenceForChannelFromChain(uint8(channelId), seqFromChain)
	} else {
		a.hasRetrievedSequence = false
		// non-inturn relayer retries every 10 second, gets the sequence from chain
		time.Sleep(time.Duration(a.config.RelayConfig.GreenfieldSequenceUpdateLatency) * time.Second)
		startSequence, err = a.bscExecutor.GetNextDeliveryOracleSequenceWithRetry()
		if err != nil {
			return err
		}
	}
	endSequence, err := a.daoManager.BSCDao.GetLatestOracleSequenceByStatus(db.AllVoted)
	if err != nil {
		return err
	}
	if endSequence == -1 {
		return nil
	}
	nonce, err := a.greenfieldExecutor.GetNonce()
	if err != nil {
		return err
	}

	for i := startSequence; i <= uint64(endSequence); i++ {
		pkgs, err := a.daoManager.BSCDao.GetPackagesByOracleSequence(i)
		if err != nil {
			return err
		}
		if len(pkgs) == 0 {
			return nil
		}
		status := pkgs[0].Status
		pkgTime := pkgs[0].TxTime

		if status != db.AllVoted && status != db.Delivered {
			return fmt.Errorf("packages with oracle sequence %d does not get enough votes yet", i)
		}

		// non-inturn relayer can not relay tx within the timeout of in-turn relayer
		if !isInturnRelyer && time.Now().Unix() < pkgTime+a.config.RelayConfig.BSCToGreenfieldInturnRelayerTimeout {
			return nil
		}
		if err := a.processPkgs(pkgs, uint8(channelId), i, nonce, isInturnRelyer); err != nil {
			return err
		}
		logging.Logger.Infof("relayed packages with oracle sequence %d ", i)

		nonce++
	}
	return nil
}

func (a *BSCAssembler) processPkgs(pkgs []*model.BscRelayPackage, channelId uint8, sequence uint64, nonce uint64, isInturnRelyer bool) error {
	// Get votes result for a packages, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(channelId, sequence)
	if err != nil {
		logging.Logger.Errorf("failed to get votes result for packages for channel %d and sequence %d", channelId, sequence)
		return err
	}

	validators, err := a.greenfieldExecutor.QueryCachedLatestValidators()
	if err != nil {
		return err
	}

	aggregatedSignature, valBitSet, err := vote.AggregateSignatureAndValidatorBitSet(votes, validators)
	if err != nil {
		return err
	}

	txHash, err := a.greenfieldExecutor.ClaimPackages(votes[0].ClaimPayload, aggregatedSignature, valBitSet.Bytes(), pkgs[0].TxTime, sequence, nonce)
	if err != nil {
		return err
	}
	logging.Logger.Infof("claimed transaction with txHash %s", txHash)

	var pkgIds []int64
	for _, p := range pkgs {
		pkgIds = append(pkgIds, p.Id)
	}
	a.metricService.SetBSCProcessedBlockHeight(pkgs[0].Height)
	if !isInturnRelyer {
		if err = a.daoManager.BSCDao.UpdateBatchPackagesClaimedTxHash(pkgIds, txHash); err != nil {
			return err
		}
		return nil
	}

	if err = a.daoManager.BSCDao.UpdateBatchPackagesStatusAndClaimedTxHash(pkgIds, db.Delivered, txHash); err != nil {
		logging.Logger.Errorf("failed to update packages to 'Delivered', error=%s", err.Error())
		return err
	}
	if err = a.daoManager.SequenceDao.Upsert(channelId, sequence+1); err != nil {
		return err
	}
	return nil
}
