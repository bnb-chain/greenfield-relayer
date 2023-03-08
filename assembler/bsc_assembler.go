package assembler

import (
	"encoding/hex"
	"errors"
	"time"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/types"
	"github.com/bnb-chain/greenfield-relayer/util"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type BSCAssembler struct {
	config             *config.Config
	greenfieldExecutor *executor.GreenfieldExecutor
	bscExecutor        *executor.BSCExecutor
	daoManager         *dao.DaoManager
	blsPubKey          string
}

func NewBSCAssembler(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager, greenfieldExecutor *executor.GreenfieldExecutor) *BSCAssembler {
	return &BSCAssembler{
		config:             cfg,
		bscExecutor:        executor,
		daoManager:         dao,
		greenfieldExecutor: greenfieldExecutor,
		blsPubKey:          hex.EncodeToString(util.BlsPubKeyFromPrivKeyStr(cfg.GreenfieldConfig.BlsPrivateKey)),
	}
}

// AssemblePackagesAndClaimLoop assemble packages and then claim in Greenfield
func (a *BSCAssembler) AssemblePackagesAndClaimLoop() {
	a.assemblePackagesAndClaimForOracleChannel(common.OracleChannelId)
}

func (a *BSCAssembler) assemblePackagesAndClaimForOracleChannel(channelId types.ChannelId) {
	for {
		if err := a.process(channelId); err != nil {
			logging.Logger.Errorf("encounter error when relaying packages, err=%s ", err.Error())
			time.Sleep(common.RetryInterval)
		}
	}
}

func (a *BSCAssembler) process(channelId types.ChannelId) error {
	nextSequence, err := a.bscExecutor.GetNextDeliveryOracleSequence()
	if err != nil {
		return err
	}

	pkgs, err := a.daoManager.BSCDao.GetPackagesByOracleSequenceAndStatus(nextSequence, db.AllVoted)
	if err != nil {
		logging.Logger.Errorf("failed to get all validator voted tx with channel id %d and sequence : %d", channelId, nextSequence)
		return err
	}
	if len(pkgs) == 0 {
		return nil
	}

	var pkgIds []int64
	for _, p := range pkgs {
		pkgIds = append(pkgIds, p.Id)
	}
	// Get votes result for a packages, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(uint8(channelId), nextSequence)
	if err != nil {
		logging.Logger.Errorf("failed to get votes result for packages for channel %d and sequence %d", channelId, nextSequence)
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

	relayerPubKeys, err := a.greenfieldExecutor.GetValidatorsBlsPublicKey()
	if err != nil {
		return err
	}

	// packages for same oracle sequence share one timestamp
	pkgTs := pkgs[0].TxTime

	relayerIdx := util.IndexOf(a.blsPubKey, relayerPubKeys)
	if relayerIdx == -1 {
		return errors.New(" relayer's bls pub key not found. ")
	}

	firstInturnRelayerIdx := int(pkgTs) % len(relayerPubKeys)
	packagesRelayStartTime := pkgTs + a.config.RelayConfig.BSCToGreenfieldRelayingDelayTime
	logging.Logger.Infof("packages will be relayed starting at %d", packagesRelayStartTime)

	var indexDiff int
	if relayerIdx >= firstInturnRelayerIdx {
		indexDiff = relayerIdx - firstInturnRelayerIdx
	} else {
		indexDiff = len(relayerPubKeys) - (firstInturnRelayerIdx - relayerIdx)
	}

	curRelayerRelayingStartTime := int64(0)
	if indexDiff == 0 {
		curRelayerRelayingStartTime = packagesRelayStartTime
	} else {
		curRelayerRelayingStartTime = packagesRelayStartTime + a.config.RelayConfig.FirstInTurnRelayerRelayingWindow + int64(indexDiff-1)*a.config.RelayConfig.InTurnRelayerRelayingWindow
	}
	logging.Logger.Infof("current relayer starts relaying from %d", curRelayerRelayingStartTime)

	filled := make(chan struct{})
	errC := make(chan error)
	ticker := time.NewTicker(common.RetryInterval)

	go a.validateSequenceFilled(filled, errC, nextSequence)

	for {
		select {
		case err = <-errC:
			return err
		case <-filled:
			if err = a.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.Delivered); err != nil {
				logging.Logger.Errorf("failed to update packages status to 'Delivered', package Ids =%v", pkgIds)
				return err
			}
			return nil
		case <-ticker.C:
			if time.Now().Unix() >= curRelayerRelayingStartTime {
				txHash, err := a.greenfieldExecutor.ClaimPackages(votes[0].ClaimPayload, aggregatedSignature, valBitSet.Bytes(), pkgTs)
				if err != nil {
					return err
				}
				logging.Logger.Infof("claimed transaction with txHash %s", txHash)
				if err = a.daoManager.BSCDao.UpdateBatchPackagesClaimedTxHash(pkgIds, txHash); err != nil {
					return err
				}
				time.Sleep(executor.GnfdSequenceUpdateLatency)
				return nil
			}
		}
	}
}

func (a *BSCAssembler) validateSequenceFilled(filled chan struct{}, errC chan error, sequence uint64) {
	ticker := time.NewTicker(common.RetrieveSequenceInterval)
	defer ticker.Stop()
	for range ticker.C {
		nextDeliverySequence, err := a.bscExecutor.GetNextDeliveryOracleSequence()
		if err != nil {
			errC <- err
		}
		if sequence < nextDeliverySequence {
			logging.Logger.Infof("Oracle sequence %d has been filled ", sequence)
			filled <- struct{}{}
		}
	}
}
