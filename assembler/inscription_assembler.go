package assembler

import (
	"inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
	"inscription-relayer/executor"
	"inscription-relayer/util"
	"inscription-relayer/vote"
	"time"
)

type InscriptionAssembler struct {
	config              *config.Config
	bscExecutor         *executor.BSCExecutor
	inscriptionExecutor *executor.InscriptionExecutor
	daoManager          *dao.DaoManager
	votePoolExecutor    *executor.VotePoolExecutor
}

func NewInscriptionAssembler(cfg *config.Config, executor *executor.InscriptionExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor) *InscriptionAssembler {
	return &InscriptionAssembler{
		config:              cfg,
		inscriptionExecutor: executor,
		daoManager:          dao,
		bscExecutor:         bscExecutor,
	}
}

// AssembleTransactionAndSend assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *InscriptionAssembler) AssembleTransactionAndSend() error {
	for _, c := range BSCMonitorChannels {
		go func(c common.ChannelId) {
			err := a.assembleTransactionAndSendForChannel(c)
			if err != nil {
				return
			}
		}(c)
	}
	return nil
}

func (a *InscriptionAssembler) assembleTransactionAndSendForChannel(channelId common.ChannelId) error {
	for {
		nextSequence, err := a.inscriptionExecutor.GetNextDeliverySequenceForChannel(channelId)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		tx, err := a.daoManager.InscriptionDao.GetTransactionByChannelIdAndSequenceWithStatusAllVoted(channelId, nextSequence)
		if err != nil {
			common.Logger.Errorf("failed to get all validator voted tx with channel id %d and sequence : %d", channelId, nextSequence)
			time.Sleep(2 * time.Second)
			continue
		}

		//Get votes result for a tx, which are already validated and qualified to aggregate sig
		votes, err := a.daoManager.VoteDao.GetVotesByChannelIdAndSequence(tx.ChannelId, tx.Sequence)
		if err != nil {
			common.Logger.Errorf("failed to get votes result for tx : %s", tx.TxHash)
			time.Sleep(2 * time.Second)
			continue
		}

		voteData, err := a.daoManager.VoteDao.GetVoteDataByChannelAndSequence(uint8(channelId), nextSequence)
		if err != nil {
			common.Logger.Errorf("failed to get votedata result for tx : %s", tx.TxHash)
			time.Sleep(2 * time.Second)
			continue
		}

		aggregatedSignature, err := vote.AggregatedSignature(votes)
		if err != nil {
			return nil
		}

		relayerPubkeys, err := a.votePoolExecutor.GetValidatorsPublicKey()
		if err != nil {
			return err
		}

		relayerPubKey := a.config.VotePoolConfig.BlsPublicKey
		relayerIdx := util.IndexOf(relayerPubKey, relayerPubkeys)
		inturnRelayerIdx := int(tx.CreatedAt) % len(relayerPubkeys)
		inturnRelayerRelayingTime := tx.CreatedAt + RelayingWindowInSecond
		common.Logger.Infof("In-turn relayer relaying time is %d", inturnRelayerRelayingTime)

		var indexDiff int
		if relayerIdx >= inturnRelayerIdx {
			indexDiff = relayerIdx - inturnRelayerIdx
		} else {
			indexDiff = len(relayerPubkeys) - (inturnRelayerIdx - relayerIdx)
		}
		curRelayerRelayingTime := inturnRelayerRelayingTime + int64(indexDiff*3) //
		common.Logger.Infof("Current relayer relaying time is %d", curRelayerRelayingTime)

		// Keep pooling the next delivery sequence from dest chain until relaying time meets, or interrupt when seq is filled
		c := make(chan struct{})
		go a.validateSequenceFilled(c, curRelayerRelayingTime, nextSequence, channelId)
		isAlreadyFilled := false
		for range c {
			isAlreadyFilled = true
		}
		// if the sequence is already filled, will update the transaction to FILLED in DB
		if isAlreadyFilled {
			if err = a.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, vote.FILLED); err != nil {
				common.Logger.Errorf("failed to update tx status %s", tx)
				return err
			}
			continue
		}

		common.Logger.Infof("relaying transaction %s", tx.TxHash)

		//TODO
		// call function handlePackage(bytes calldata payload, bytes calldata signature, uint256 validatorSet, uint64 packageSequence, uint8 channelId)
		// a.bscExecutor.CallBuildInSystemContract([]bytes(tx.PayLoad), aggregatedSignature, validatorSet, nextSequence, channelId)
		println(voteData)
		println(aggregatedSignature)

	}
}

func (a *InscriptionAssembler) validateSequenceFilled(c chan struct{}, curRelayerRelayingTime int64, sequence uint64, channelID common.ChannelId) {
	for time.Now().Unix() < curRelayerRelayingTime {
		nextDeliverySequence, _ := a.inscriptionExecutor.GetNextDeliverySequenceForChannel(channelID)
		if sequence <= nextDeliverySequence-1 {
			common.Logger.Infof("sequence %d for channel %d has already been filled ", sequence, channelID)
			c <- struct{}{}
			close(c)
			return
		}
		time.Sleep(1 * time.Second)
	}
	close(c)
}
