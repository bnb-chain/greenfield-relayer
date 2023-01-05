package assembler

import (
	"context"
	"encoding/hex"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	"inscription-relayer/util"
	"inscription-relayer/vote"
	"math/big"
	"time"
)

type InscriptionAssembler struct {
	config              *config.Config
	bscExecutor         *executor.BSCExecutor
	inscriptionExecutor *executor.InscriptionExecutor
	daoManager          *dao.DaoManager
	votePoolExecutor    *vote.VotePoolExecutor
}

func NewInscriptionAssembler(cfg *config.Config, executor *executor.InscriptionExecutor, dao *dao.DaoManager, bscExecutor *executor.BSCExecutor, votePoolExecutor *vote.VotePoolExecutor) *InscriptionAssembler {
	return &InscriptionAssembler{
		config:              cfg,
		inscriptionExecutor: executor,
		daoManager:          dao,
		bscExecutor:         bscExecutor,
		votePoolExecutor:    votePoolExecutor,
	}
}

// AssembleTransactionAndSend assemble a tx by gathering votes signature and then call the build-in smart-contract
func (a *InscriptionAssembler) AssembleTransactionAndSend() {
	for _, c := range a.config.InscriptionConfig.MonitorChannelList {
		go func(c common.ChannelId) {
			err := a.assembleTransactionAndSendForChannel(c)
			if err != nil {
				panic(fmt.Sprintf("failed to assembleTransactionAndSendForChannel for channel %d, err %s", c, err.Error()))
			}
		}(common.ChannelId(c))
	}
}

func (a *InscriptionAssembler) assembleTransactionAndSendForChannel(channelId common.ChannelId) error {
	for {
		nextSequence, err := a.inscriptionExecutor.GetNextDeliverySequenceForChannel(channelId)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		tx, err := a.daoManager.InscriptionDao.GetTransactionByChannelIdAndSequence(channelId, nextSequence)
		if tx.Status != model.VOTED_ALL {
			common.Logger.Infof("there are not enough votes collected for tx yet. txHash=%s, current status is %d", tx.TxHash, tx.Status)
			time.Sleep(2 * time.Second)
			continue
		}
		if err != nil {
			common.Logger.Errorf("failed to get all voted tx with channel id %d and sequence : %d", channelId, nextSequence)
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

		if err != nil {
			common.Logger.Errorf("failed to get votedata result for tx : %s", tx.TxHash)
			time.Sleep(2 * time.Second)
			continue
		}

		validators, err := a.inscriptionExecutor.QueryLatestValidators()
		aggregatedSignature, votedAddressSet, err := vote.AggregatedSignatureAndValidatorBitSet(votes, validators)
		if err != nil {
			return err
		}
		validatorBitset := big.NewInt(int64(votedAddressSet))

		relayerBlsPubKeys, err := a.votePoolExecutor.GetValidatorsBlsPublicKey()
		if err != nil {
			return err
		}

		relayerPubKey, err := util.GetBlsPubKeyFromPrivKeyStr(a.config.VotePoolConfig.BlsPrivateKey)
		if err != nil {
			return err
		}
		relayerIdx := util.IndexOf(hex.EncodeToString(relayerPubKey), relayerBlsPubKeys)
		inturnRelayerIdx := int(tx.TxTime) % len(relayerBlsPubKeys)
		inturnRelayerRelayingTime := tx.TxTime + RelayingWindowInSecond
		common.Logger.Infof("In-turn relayer relaying time is %d", inturnRelayerRelayingTime)

		var indexDiff int
		if relayerIdx >= inturnRelayerIdx {
			indexDiff = relayerIdx - inturnRelayerIdx
		} else {
			indexDiff = len(relayerBlsPubKeys) - (inturnRelayerIdx - relayerIdx)
		}
		curRelayerRelayingTime := inturnRelayerRelayingTime + int64(indexDiff*3)
		common.Logger.Infof("Current relayer relaying time is %d", curRelayerRelayingTime)

		// Keep pooling the next delivery sequence from dest chain until relaying time meets, or interrupt when seq is filled
		isAlreadyFilled, err := a.validateSequenceFilled(curRelayerRelayingTime, nextSequence, channelId)
		if err != nil {
			return err
		}

		if isAlreadyFilled {
			if err = a.daoManager.InscriptionDao.UpdateTxStatus(tx.Id, model.FILLED); err != nil {
				common.Logger.Errorf("failed to update tx status %s", tx)
				return err
			}
			continue
		}

		common.Logger.Infof("relaying transaction %s", tx.TxHash)
		nonce, err := a.bscExecutor.GetClient().PendingNonceAt(context.Background(), a.bscExecutor.TxSender)
		txHash, err := a.bscExecutor.CallBuildInSystemContract(int8(channelId), aggregatedSignature, nextSequence, validatorBitset, ethcommon.Hex2Bytes(tx.PayLoad), nonce)
		if err != nil {
			//return err
		}
		err = a.daoManager.InscriptionDao.UpdateTxStatus(tx.Id, model.FILLED)
		if err != nil {
			common.Logger.Errorf("failed to update Tx status %d", tx.Id)
			return err
		}

		common.Logger.Infof("relayed transaction %s, the txHash returned from BSC is %s", tx.TxHash, txHash)
	}
}

func (a *InscriptionAssembler) validateSequenceFilled(curRelayerRelayingTime int64, sequence uint64, channelID common.ChannelId) (bool, error) {
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
