package assembler

import (
	"encoding/hex"
	"time"

	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type BSCAssembler struct {
	config              *config.Config
	inscriptionExecutor *executor.InscriptionExecutor
	bscExecutor         *executor.BSCExecutor
	daoManager          *dao.DaoManager
	votePoolExecutor    *vote.VotePoolExecutor
}

func NewBSCAssembler(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager, votePoolExecutor *vote.VotePoolExecutor, inscriptionExecutor *executor.InscriptionExecutor) *BSCAssembler {
	return &BSCAssembler{
		config:              cfg,
		bscExecutor:         executor,
		daoManager:          dao,
		votePoolExecutor:    votePoolExecutor,
		inscriptionExecutor: inscriptionExecutor,
	}
}

// AssemblePackagesAndClaim assemble packages and then claim in Inscription
func (a *BSCAssembler) AssemblePackagesAndClaim() {
	a.assemblePackagesAndClaimForOracleChannel(common.OracleChannelId)
}

func (a *BSCAssembler) assemblePackagesAndClaimForOracleChannel(channelId common.ChannelId) {
	for {
		err := a.process(channelId)
		if err != nil {
			common.Logger.Errorf("encounter error when relaying packages, err=%s ", err.Error())
			time.Sleep(RetryInterval)
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
	if len(pkgs) == 0 {
		return nil
	}
	for _, p := range pkgs {
		pkgIds = append(pkgIds, p.Id)
	}
	if err != nil {
		common.Logger.Errorf("failed to get all validator voted tx with channel id %d and sequence : %d", channelId, nextSequence)
		return err
	}
	// Get votes result for a packages, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(uint8(channelId), nextSequence)
	if err != nil {
		common.Logger.Errorf("failed to get votes result for packages for channel %d and sequence %d", channelId, nextSequence)
		return err
	}

	// TODO switch to query validators from BSC lightcleint
	validators, err := a.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return err
	}
	aggregatedSignature, valBitSet, err := vote.AggregateSignatureAndValidatorBitSet(votes, validators)
	if err != nil {
		return err
	}

	relayerPubKeys, err := a.votePoolExecutor.GetValidatorsBlsPublicKey()
	if err != nil {
		return err
	}

	// packages for same oracle sequence share a timestamp
	pkgTs := pkgs[0].TxTime

	relayerPubKey := util.GetBlsPubKeyFromPrivKeyStr(a.getBlsPrivateKey())
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerPubKeys)
	inturnRelayerIdx := int(pkgTs) % len(relayerPubKeys)
	inturnRelayerRelayingTime := pkgTs + RelayWindowInSecond

	common.Logger.Infof("in-turn relayer relaying time is %d", inturnRelayerRelayingTime)

	var indexDiff int
	if relayerIdx >= inturnRelayerIdx {
		indexDiff = relayerIdx - inturnRelayerIdx
	} else {
		indexDiff = len(relayerPubKeys) - (inturnRelayerIdx - relayerIdx)
	}
	curRelayerRelayingStartTime := inturnRelayerRelayingTime + int64(indexDiff*RelayIntervalBetweenRelayersInSecond)
	common.Logger.Infof("current relayer starts relaying from %d", curRelayerRelayingStartTime)

	// Keep pooling the next delivery sequence from dest chain until relaying time meets, or interrupt when seq is filled
	filled := make(chan struct{})
	defer close(filled)
	errC := make(chan error)
	defer close(errC)
	go a.validateSequenceFilled(filled, errC, nextSequence)

	for {
		select {
		case err = <-errC:
			return err
		case <-filled:
			if err = a.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.Filled); err != nil {
				common.Logger.Errorf("failed to update packages status %s", pkgIds)
				return err
			}
			return nil
		default:
			if time.Now().Unix() >= curRelayerRelayingStartTime {
				txHash, err := a.inscriptionExecutor.ClaimPackages(votes[0].Payload, aggregatedSignature, valBitSet.Bytes(), pkgTs)
				if err != nil {
					return err
				}
				common.Logger.Infof("claimed transaction with txHash %s", txHash)
				err = a.daoManager.BSCDao.UpdateBatchPackagesStatusAndClaimedTxHash(pkgIds, db.Filled, txHash)
				if err != nil {
					common.Logger.Errorf("failed to update packages error %s", err.Error())
					return err
				}
			}
		}
	}
}

func (a *BSCAssembler) validateSequenceFilled(filled chan struct{}, errC chan error, sequence uint64) {
	ticker := time.NewTicker(RetryInterval)
	defer ticker.Stop()
	for {
		nextDeliverySequence, err := a.bscExecutor.GetNextDeliveryOracleSequence()
		if err != nil {
			errC <- err
		}
		if sequence < nextDeliverySequence {
			common.Logger.Infof("Oracle sequence %d for has already been filled ", sequence)
			filled <- struct{}{}
		}
		<-ticker.C
	}
}

func (a *BSCAssembler) getBlsPrivateKey() string {
	return a.config.VotePoolConfig.BlsPrivateKey
}
