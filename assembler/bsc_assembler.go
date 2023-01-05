package assembler

import (
	"encoding/hex"
	"github.com/willf/bitset"
	"inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	"inscription-relayer/util"
	"inscription-relayer/vote"
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
	err := a.assemblePackagesAndClaimForOracleChannel(common.OracleChannelId)
	if err != nil {
		panic(err)
	}
}

func (a *BSCAssembler) assemblePackagesAndClaimForOracleChannel(channelId common.ChannelId) error {
	for {
		err := a.process(channelId)
		if err != nil {
			time.Sleep(1 * time.Second)
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
	aggregatedSignature, votedAddressSet, err := vote.AggregatedSignatureAndValidatorBitSet(votes, validators)
	valBitset := bitset.From([]uint64{votedAddressSet})
	if err != nil {
		return nil
	}

	relayerPubKeys, err := a.votePoolExecutor.GetValidatorsBlsPublicKey()
	if err != nil {
		return err
	}

	relayerPubKey, err := util.GetBlsPubKeyFromPrivKeyStr(a.config.VotePoolConfig.BlsPrivateKey)
	if err != nil {
		return err
	}
	relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerPubKeys)
	inturnRelayerIdx := int(pkgs[0].TxTime) % len(relayerPubKeys)
	inturnRelayerRelayingTime := pkgs[0].TxTime + RelayingWindowInSecond

	common.Logger.Infof("In-turn relayer relaying time is %d", inturnRelayerRelayingTime)

	var indexDiff int
	if relayerIdx >= inturnRelayerIdx {
		indexDiff = relayerIdx - inturnRelayerIdx
	} else {
		indexDiff = len(relayerPubKeys) - (inturnRelayerIdx - relayerIdx)
	}
	curRelayerRelayingTime := inturnRelayerRelayingTime + int64(indexDiff*3)
	common.Logger.Infof("Current relayer realying time is %d", curRelayerRelayingTime)

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
	txHash, err := a.inscriptionExecutor.ClaimPackages(votes[0].EventHash, aggregatedSignature, valBitset.Bytes())
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
