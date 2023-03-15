package assembler

import (
	"encoding/hex"
	"fmt"
	"github.com/bnb-chain/greenfield-relayer/db/model"
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
	ticker := time.NewTicker(InturnRelayerAssembleInterval)
	for range ticker.C {
		if err := a.process(channelId); err != nil {
			logging.Logger.Errorf("encounter error when relaying packages, err=%s ", err.Error())
			time.Sleep(common.RetryInterval)
		}
	}
}

func (a *BSCAssembler) process(channelId types.ChannelId) error {
	inturnRelayer, err := a.greenfieldExecutor.GetInturnRelayer()
	if err != nil {
		return err
	}
	isInturnRelyer := inturnRelayer.BlsPubKey == a.blsPubKey
	var startSequence uint64
	if isInturnRelyer {
		seq, err := a.daoManager.SequenceDao.GetByChannelId(uint8(channelId))
		if err != nil {
			return err
		}
		startSequence = uint64(seq.Sequence)

		// in-turn relayer get the start sequence from chain first time, it starts to relay after the sequence gets updated
		timeDiff := time.Now().Unix() - int64(inturnRelayer.RelayInterval.Start)
		if timeDiff < GNFDSequenceUpdateWaitingTime {
			time.Sleep(time.Duration(timeDiff))
			startSequence, err = a.bscExecutor.GetNextDeliveryOracleSequenceWithRetry()
			if err != nil {
				return err
			}
			if err = a.daoManager.SequenceDao.Upsert(uint8(channelId), startSequence); err != nil {
				return err
			}
		}
		logging.Logger.Debug("relay as inturn relayer")
	} else {
		// non-inturn relayer retries every 10 second, gets the sequence from chain
		time.Sleep(GNFDSequenceUpdateWaitingTime * time.Second)
		startSequence, err = a.bscExecutor.GetNextDeliveryOracleSequenceWithRetry()
		if err != nil {
			return err
		}
		logging.Logger.Debug("relay as non-inturn relayer")
		if err := a.daoManager.BSCDao.UpdateBatchPackagesStatusToDelivered(startSequence); err != nil {
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
		if err := a.processPkgs(pkgs, uint8(channelId), i, nonce); err != nil {
			return err
		}
		logging.Logger.Infof("relayed packages with oracle sequence %d ", i)
		// update next delivery sequence in DB
		if err = a.daoManager.SequenceDao.Upsert(uint8(channelId), i+1); err != nil {
			return err
		}
		nonce++
	}
	return nil
}

func (a *BSCAssembler) processPkgs(pkgs []*model.BscRelayPackage, channelId uint8, sequence uint64, nonce uint64) error {
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
	if err = a.daoManager.BSCDao.UpdateBatchPackagesStatusAndClaimedTxHash(pkgIds, db.Delivered, txHash); err != nil {
		logging.Logger.Errorf("failed to update packages to 'Delivered', error=%s", err.Error())
		return err
	}
	return nil
}
