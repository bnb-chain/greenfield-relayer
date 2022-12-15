package assembler

import (
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
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
	votePoolExecutor    *executor.VotePoolExecutor
}

func NewBSCAssembler(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager, votePoolExecutor *executor.VotePoolExecutor, inscriptionExecutor *executor.InscriptionExecutor) *BSCAssembler {
	return &BSCAssembler{
		config:              cfg,
		bscExecutor:         executor,
		daoManager:          dao,
		votePoolExecutor:    votePoolExecutor,
		inscriptionExecutor: inscriptionExecutor,
	}
}

// AssemblePackagesAndClaim assemble packages and then claim in Inscription
func (a *BSCAssembler) AssemblePackagesAndClaim() error {
	for _, c := range InscriptionMonitorChannels {
		go func(c relayercommon.ChannelId) {
			err := a.assemblePackagesAndClaimForChannel(c)
			if err != nil {
				return
			}
		}(c)
	}
	return nil
}

func (a *BSCAssembler) assemblePackagesAndClaimForChannel(channelId relayercommon.ChannelId) error {
	for {
		nextSequence, err := a.bscExecutor.GetNextDeliverySequenceForChannel(channelId)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		var pkgIds []int64
		pkgs, err := a.daoManager.BSCDao.GetAllVotedPackages(channelId, nextSequence)

		for _, p := range pkgs {
			pkgIds = append(pkgIds, p.Id)
		}
		if err != nil {
			relayercommon.Logger.Errorf("failed to get all validator voted tx with channel id %d and sequence : %d", channelId, nextSequence)
			time.Sleep(1 * time.Second)
			continue
		}
		//Get votes result for a packages, which are already validated and qualified to aggregate sig
		votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(uint8(channelId), nextSequence)
		if err != nil {
			relayercommon.Logger.Errorf("failed to get votes result for packages for channel %d and sequence %d", channelId, nextSequence)
			time.Sleep(2 * time.Second)
			continue
		}
		voteData, err := a.daoManager.VoteDao.GetVoteDataByChannelAndSequence(uint8(channelId), nextSequence)
		if err != nil {
			relayercommon.Logger.Errorf("failed to get votes result for packages for channel %d and sequence %d", channelId, nextSequence)
			time.Sleep(2 * time.Second)
			continue
		}
		aggregatedSignature, err := vote.AggregatedSignature(votes)
		if err != nil {
			return nil
		}

		msgClaim := &MsgClaim{
			FromAddress:    a.bscExecutor.TxSender.String(),
			ChainId:        uint16(channelId),
			Sequence:       nextSequence,
			TimeStamp:      uint64(pkgs[0].CreatedAt),
			Payload:        voteData.EventHash,
			VoteAddressSet: []uint64{1, 2},
			AggSignature:   aggregatedSignature,
		}

		relayerPubkeys, err := a.votePoolExecutor.GetValidatorsPublicKey()
		if err != nil {
			return err
		}

		relayerPubKey := a.config.VotePoolConfig.BlsPublicKey
		relayerIdx := util.IndexOf(relayerPubKey, relayerPubkeys)
		inturnRelayerIdx := int(msgClaim.TimeStamp) % len(relayerPubkeys)
		inturnRelayerRelayingTime := int64(msgClaim.TimeStamp) + RelayingWindowInSecond
		relayercommon.Logger.Infof("In-turn relayer relaying time is %d", inturnRelayerRelayingTime)

		var indexDiff int
		if relayerIdx >= inturnRelayerIdx {
			indexDiff = relayerIdx - inturnRelayerIdx
		} else {
			indexDiff = len(relayerPubkeys) - (inturnRelayerIdx - relayerIdx)
		}
		curRelayerRelayingTime := inturnRelayerRelayingTime + int64(indexDiff*3)
		relayercommon.Logger.Infof("Current relayer realying time is %d", curRelayerRelayingTime)

		// Keep pooling the next delivery sequence from dest chain until relaying time meets, or interrupt when seq is filled
		c := make(chan struct{})
		go a.validateSequenceFilled(c, curRelayerRelayingTime, nextSequence, channelId)
		isAlreadyFilled := false
		for range c {
			isAlreadyFilled = true
		}
		// if the sequence is already filled, update packages status to FILLED in DB
		if isAlreadyFilled {
			if err = a.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, vote.FILLED); err != nil {
				relayercommon.Logger.Errorf("failed to update packages status %s", pkgIds)
				return err
			}
			continue
		}

		relayercommon.Logger.Infof("claim transaction ")

		//TODO validate claim response
		_, err = a.inscriptionExecutor.ClaimPackages(msgClaim)
		if err != nil {
			return err
		}
		//TODO update related packages status by the response
		err = a.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, vote.FILLED)
		if err != nil {
			relayercommon.Logger.Errorf("failed to update packages status %s", pkgIds)
			continue
		}
	}
}

func (a *BSCAssembler) validateSequenceFilled(c chan struct{}, curRelayerRelayingTime int64, sequence uint64, channelID relayercommon.ChannelId) {
	for time.Now().Unix() < curRelayerRelayingTime {
		nextDeliverySequence, _ := a.bscExecutor.GetNextDeliverySequenceForChannel(channelID)
		if sequence <= nextDeliverySequence-1 {
			relayercommon.Logger.Infof("sequence %d for channel %d has already been filled ", sequence, channelID)
			c <- struct{}{}
			close(c)
			return
		}
		time.Sleep(1 * time.Second)
	}
	close(c)
}
