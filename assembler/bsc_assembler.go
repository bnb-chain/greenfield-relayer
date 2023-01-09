package assembler

import (
	"encoding/hex"
	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	"github.com/bnb-chain/inscription-relayer/vote"
	"time"
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
	pkgs, err := a.daoManager.BSCDao.GetAllVotedPackages(channelId, nextSequence)
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
	//Get votes result for a packages, which are already validated and qualified to aggregate sig
	votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(uint8(channelId), nextSequence)
	if err != nil {
		common.Logger.Errorf("failed to get votes result for packages for channel %d and sequence %d", channelId, nextSequence)
		return err
	}

	//TODO switch to query validators from BSC lightcleint
	validators, err := a.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return err
	}
	aggregatedSignature, valBitSet, err := vote.AggregatedSignatureAndValidatorBitSet(votes, validators)

	if err != nil {
		return err
	}

	relayerPubKeys, err := a.votePoolExecutor.GetValidatorsBlsPublicKey()
	if err != nil {
		return err
	}

	relayerPubKey := util.GetBlsPubKeyFromPrivKeyStr(a.getBlsPrivateKey())

	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerPubKeys)
	inturnRelayerIdx := int(pkgs[0].TxTime) % len(relayerPubKeys)
	inturnRelayerRelayingTime := pkgs[0].TxTime + RelayWindowInSecond

	common.Logger.Infof("In-turn relayer relaying time is %d", inturnRelayerRelayingTime)

	var indexDiff int
	if relayerIdx >= inturnRelayerIdx {
		indexDiff = relayerIdx - inturnRelayerIdx
	} else {
		indexDiff = len(relayerPubKeys) - (inturnRelayerIdx - relayerIdx)
	}
	curRelayerRelayingTime := inturnRelayerRelayingTime + int64(indexDiff*RelayIntervalBetweenRelayersInSecond)
	common.Logger.Infof("Current relayer relaying time is %d", curRelayerRelayingTime)

	// Keep pooling the next delivery sequence from dest chain until relaying time meets, or interrupt when seq is filled

	isAlreadyFilled, err := a.validateSequenceFilled(curRelayerRelayingTime, nextSequence, channelId)
	if err != nil {
		return err
	}
	// if the sequence is already filled, update packages status to FILLED in DB
	if isAlreadyFilled {
		if err = a.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, model.FILLED); err != nil {
			common.Logger.Errorf("failed to update packages status %s", pkgIds)
			return err
		}
		return nil
	}
	txHash, err := a.inscriptionExecutor.ClaimPackages(votes[0].EventHash, aggregatedSignature, valBitSet.Bytes())
	if err != nil {
		return err
	}
	common.Logger.Infof("claimed transaction with txHash %s", txHash)
	err = a.daoManager.BSCDao.UpdateBatchPackagesStatusAndClaimTxHash(pkgIds, model.FILLED, txHash)
	if err != nil {
		common.Logger.Errorf("failed to update packages error %s", err.Error())
		return err
	}
	return nil
}

func (a *BSCAssembler) validateSequenceFilled(curRelayerRelayingTime int64, sequence uint64, channelID common.ChannelId) (bool, error) {
	for time.Now().Unix() < curRelayerRelayingTime {
		nextDeliverySequence, err := a.inscriptionExecutor.GetNextDeliverySequenceForChannel(channelID)
		if err != nil {
			return false, err
		}
		if sequence <= nextDeliverySequence-1 {
			common.Logger.Infof("sequence %d for channel %d has already been filled ", sequence, channelID)
			return true, nil
		}
	}
	return false, nil
}

func (a *BSCAssembler) getBlsPrivateKey() string {
	return a.config.VotePoolConfig.BlsPrivateKey
}
