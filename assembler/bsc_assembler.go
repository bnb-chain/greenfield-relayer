package assembler

import (
	"encoding/hex"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"time"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/util"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type BSCAssembler struct {
	config             *config.Config
	greenfieldExecutor *executor.GreenfieldExecutor
	bscExecutor        *executor.BSCExecutor
	daoManager         *dao.DaoManager
	votePoolExecutor   *vote.VotePoolExecutor
}

func NewBSCAssembler(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager, votePoolExecutor *vote.VotePoolExecutor, greenfieldExecutor *executor.GreenfieldExecutor) *BSCAssembler {
	return &BSCAssembler{
		config:             cfg,
		bscExecutor:        executor,
		daoManager:         dao,
		votePoolExecutor:   votePoolExecutor,
		greenfieldExecutor: greenfieldExecutor,
	}
}

// AssemblePackagesAndClaim assemble packages and then claim in Greenfield
func (a *BSCAssembler) AssemblePackagesAndClaim() {
	a.assemblePackagesAndClaimForOracleChannel(common.OracleChannelId)
}

func (a *BSCAssembler) assemblePackagesAndClaimForOracleChannel(channelId common.ChannelId) {
	for {
		err := a.process(channelId)
		if err != nil {
			logging.Logger.Errorf("encounter error when relaying packages, err=%s ", err.Error())
			time.Sleep(common.RetryInterval)
		}
	}
}

func (a *BSCAssembler) process(channelId common.ChannelId) error {
	nextSequence, err := a.bscExecutor.GetNextDeliveryOracleSequence()
	if err != nil {
		return err
	}
	var pkgIds []int64
	pkgs, err := a.daoManager.BSCDao.GetAllVotedPackages(nextSequence)
	if err != nil {
		logging.Logger.Errorf("failed to get all validator voted tx with channel id %d and sequence : %d", channelId, nextSequence)
		return err
	}
	if len(pkgs) == 0 {
		return nil
	}
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

	// packages for same oracle sequence share a timestamp
	pkgTs := pkgs[0].TxTime

	relayerPubKey := util.GetBlsPubKeyFromPrivKeyStr(a.getBlsPrivateKey())
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerPubKeys)
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
				err = a.daoManager.BSCDao.UpdateBatchPackagesStatusAndClaimedTxHash(pkgIds, db.Delivered, txHash)
				if err != nil {
					logging.Logger.Errorf("failed to update packages to 'Delivered', error=%s", err.Error())
					return err
				}
				return nil
			}
		}
	}
}

func (a *BSCAssembler) validateSequenceFilled(filled chan struct{}, errC chan error, sequence uint64) {
	ticker := time.NewTicker(common.RetryInterval)
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

func (a *BSCAssembler) getBlsPrivateKey() string {
	return a.config.VotePoolConfig.BlsPrivateKey
}
