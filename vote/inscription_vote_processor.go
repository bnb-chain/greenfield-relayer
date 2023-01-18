package vote

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/avast/retry-go/v4"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	"github.com/tendermint/tendermint/votepool"
	"gorm.io/gorm"
)

type InscriptionVoteProcessor struct {
	votePoolExecutor    *VotePoolExecutor
	daoManager          *dao.DaoManager
	config              *config.Config
	signer              *VoteSigner
	inscriptionExecutor *executor.InscriptionExecutor
	blsPublicKey        []byte
}

func NewInscriptionVoteProcessor(cfg *config.Config, dao *dao.DaoManager, signer *VoteSigner, inscriptionExecutor *executor.InscriptionExecutor,
	votePoolExecutor *VotePoolExecutor,
) *InscriptionVoteProcessor {
	return &InscriptionVoteProcessor{
		config:              cfg,
		daoManager:          dao,
		signer:              signer,
		inscriptionExecutor: inscriptionExecutor,
		votePoolExecutor:    votePoolExecutor,
		blsPublicKey:        util.GetBlsPubKeyFromPrivKeyStr(cfg.VotePoolConfig.BlsPrivateKey),
	}
}

// SignAndBroadcast Will sign using the bls private key, broadcast the vote to votepool
func (p *InscriptionVoteProcessor) SignAndBroadcast() {
	for {
		err := p.signAndBroadcast()
		if err != nil {
			time.Sleep(RetryInterval)
		}
	}
}

func (p *InscriptionVoteProcessor) signAndBroadcast() error {
	latestHeight, err := p.inscriptionExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
		return err
	}

	leastSavedTxHeight, err := p.daoManager.InscriptionDao.GetLeastSavedTransactionHeight()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get least saved tx, error: %s", err.Error())
		return err
	}
	if leastSavedTxHeight+p.config.InscriptionConfig.NumberOfBlocksForFinality > latestHeight {
		return nil
	}
	txs, err := p.daoManager.InscriptionDao.GetTransactionsByStatusAndHeight(db.Saved, leastSavedTxHeight)
	if err != nil {
		relayercommon.Logger.Errorf("failed to get transactions at height %d from db, error: %s", leastSavedTxHeight, err.Error())
		return err
	}

	if len(txs) == 0 {
		return nil
	}

	// for every tx, we are going to sign it and broadcast vote of it.
	for _, tx := range txs {
		nextDeliverySequence, err := p.inscriptionExecutor.GetNextDeliverySequenceForChannel(relayercommon.ChannelId(tx.ChannelId))
		if err != nil {
			return err
		}
		if tx.Sequence < nextDeliverySequence {
			relayercommon.Logger.Infof("sequence %d for channel %d has already been filled ", tx.Sequence, tx.ChannelId)
			continue
		}
		aggregatedPayload, err := p.aggregatePayloadForTx(tx)
		if err != nil {
			return err
		}
		v := p.constructVoteAndSign(aggregatedPayload)

		// TODO remove testing purpose code
		bs2 := common.Hex2Bytes("2969268e6722a8e16579e7a3380f83a2dd0b15478a2994cb0ac6480e1aead999")
		secretKey2, err := blst.SecretKeyFromBytes(bs2)
		if err != nil {
			panic(err)
		}
		eh := p.getEventHash(aggregatedPayload)
		pubKey2 := secretKey2.PublicKey()
		sign2 := secretKey2.Sign(eh).Marshal()

		mockVoteFromRelayer2 := &votepool.Vote{
			PubKey:    pubKey2.Marshal(),
			Signature: sign2,
			EventType: 1,
			EventHash: eh,
		}

		// TODO remove testing purpose code
		bs3 := common.Hex2Bytes("6f235c2c0d91ecdf961f4409061a785d456b9bc4b398e2a0940378397772cb0b")
		secretKey3, err := blst.SecretKeyFromBytes(bs3)
		if err != nil {
			panic(err)
		}
		pubKey3 := secretKey3.PublicKey()
		sign3 := secretKey3.Sign(eh).Marshal()

		mockVoteFromRelayer3 := &votepool.Vote{
			PubKey:    pubKey3.Marshal(),
			Signature: sign3,
			EventType: 1,
			EventHash: eh,
		}

		// broadcast v
		if err = retry.Do(func() error {
			err = p.votePoolExecutor.BroadcastVote(mockVoteFromRelayer3)
			err = p.votePoolExecutor.BroadcastVote(mockVoteFromRelayer2)
			err = p.votePoolExecutor.BroadcastVote(v)
			if err != nil {
				return fmt.Errorf("failed to submit vote for event with channel id %d and sequence %d", tx.ChannelId, tx.Sequence)
			}
			return nil
		}, retry.Context(context.Background()), relayercommon.RtyAttem, relayercommon.RtyDelay, relayercommon.RtyErr); err != nil {
			return err
		}

		// After vote submitted to vote pool, persist vote Data and update the status of tx to 'SELF_VOTED'.
		err = p.daoManager.InscriptionDao.DB.Transaction(func(dbTx *gorm.DB) error {
			err = p.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, db.SelfVoted)
			if err != nil {
				return err
			}
			err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, tx.ChannelId, tx.Sequence, aggregatedPayload))
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *InscriptionVoteProcessor) CollectVotes() {
	for {
		err := p.collectVotes()
		if err != nil {
			time.Sleep(RetryInterval)
		}
	}
}

func (p *InscriptionVoteProcessor) collectVotes() error {
	txs, err := p.daoManager.InscriptionDao.GetTransactionsByStatus(db.SelfVoted)
	if err != nil {
		relayercommon.Logger.Errorf("failed to get voted transactions from db, error: %s", err.Error())
		return err
	}
	for _, tx := range txs {
		err := p.prepareEnoughValidVotesForTx(tx)
		if err != nil {
			return err
		}
		err = p.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, db.AllVoted)
		if err != nil {
			return err
		}
	}
	return nil
}

// prepareEnoughValidVotesForTx fetches and validate votes result, store in vote table
func (p *InscriptionVoteProcessor) prepareEnoughValidVotesForTx(tx *model.InscriptionRelayTransaction) error {
	// TODO Remove this after testing
	// validators, err := p.inscriptionExecutor.QueryCachedLatestValidators()

	localVote, err := p.daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(tx.ChannelId, tx.Sequence, hex.EncodeToString(p.blsPublicKey))
	if err != nil {
		return err
	}

	// TODO switch to query bsc valicators once ready
	// validators, err := p.inscriptionExecutor.bscExecutor.QueryCachedLatestValidators()
	validators, err := p.inscriptionExecutor.QueryCachedLatestValidators()
	if err != nil {
		return err
	}

	err = p.queryMoreThanTwoThirdVotesForTx(localVote, validators, tx.Id)
	if err != nil {
		return err
	}
	return nil
}

// queryMoreThanTwoThirdVotesForTx queries votes from votePool   // TODO switch to BSC validators once ready
func (p *InscriptionVoteProcessor) queryMoreThanTwoThirdVotesForTx(localVote *model.Vote, validators []stakingtypes.Validator, txId int64) error {
	triedTimes := 0
	validVotesTotalCount := 1 // assume local vote is valid
	channelId := localVote.ChannelId
	seq := localVote.Sequence
	ticker := time.NewTicker(2 * time.Second)
	for {
		<-ticker.C
		triedTimes++
		// skip current tx if reach the max retry.
		if triedTimes >= QueryVotepoolMaxRetryTimes {
			if err := p.daoManager.InscriptionDao.UpdateTransactionStatus(txId, 5); err != nil {
				relayercommon.Logger.Errorf("failed to update tx id=%d status", txId)
				return err
			}
			return nil
		}

		queriedVotes, err := p.votePoolExecutor.QueryVotes(localVote.EventHash, votepool.ToBscCrossChainEvent)
		if err != nil {
			relayercommon.Logger.Errorf("encounter error when query votes. will retry.")
			return err
		}
		validVotesCountPerReq := len(queriedVotes)
		if validVotesCountPerReq == 0 {
			continue
		}

		isLocalVoteIncluded := false

		for _, v := range queriedVotes {
			if !p.isVotePubKeyValid(v, validators) {
				relayercommon.Logger.Errorf("vote's pub-key %s does not belong to any validator", hex.EncodeToString(v.PubKey[:]))
				validVotesCountPerReq--
				continue
			}

			if err := VerifySignature(v, localVote.EventHash); err != nil {
				relayercommon.Logger.Errorf("verify vote's signature failed,  err=%s", err)
				validVotesCountPerReq--
				continue
			}

			// it is local vote
			if bytes.Equal(v.PubKey[:], p.blsPublicKey) {
				isLocalVoteIncluded = true
				validVotesCountPerReq--
				continue
			}

			// check duplicate, the vote might have been saved in previous request.
			exist, err := p.daoManager.VoteDao.IsVoteExist(channelId, seq, hex.EncodeToString(v.PubKey[:]))
			if err != nil {
				return err
			}
			if exist {
				validVotesCountPerReq--
				continue
			}
			// a vote result persisted into DB should be valid, unique.
			err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, channelId, seq, localVote.ClaimPayload))
			if err != nil {
				return err
			}
		}

		validVotesTotalCount += validVotesCountPerReq

		if validVotesTotalCount > len(validators)*2/3 {
			return nil
		}
		if !isLocalVoteIncluded {
			v, err := DtoToEntity(localVote)
			if err != nil {
				return err
			}
			err = p.votePoolExecutor.BroadcastVote(v)
			if err != nil {
				return err
			}
		}
		continue
	}
}

func (p *InscriptionVoteProcessor) constructVoteAndSign(aggregatedPayload []byte) *votepool.Vote {
	var v votepool.Vote
	v.EventType = votepool.ToBscCrossChainEvent
	v.EventHash = p.getEventHash(aggregatedPayload)
	p.signer.SignVote(&v)
	return &v
}

func (p *InscriptionVoteProcessor) getEventHash(aggregatedPayload []byte) []byte {
	return crypto.Keccak256Hash(aggregatedPayload).Bytes()
}

// TODO switch to BSC validators once ready
func (p *InscriptionVoteProcessor) isVotePubKeyValid(v *votepool.Vote, validators []stakingtypes.Validator) bool {
	for _, validator := range validators {
		if bytes.Equal(v.PubKey[:], validator.RelayerBlsKey[:]) {
			return true
		}
	}
	return false
}

// aggregatePayloadForTx aggregate required fields by concatenating their bytes, this will be used as payload when calling BSC smart contract

func (p *InscriptionVoteProcessor) aggregatePayloadForTx(tx *model.InscriptionRelayTransaction) ([]byte, error) {
	var aggregatedPayload []byte

	aggregatedPayload = append(aggregatedPayload, util.Uint16ToBytes(uint16(tx.SrcChainId))...)
	aggregatedPayload = append(aggregatedPayload, util.Uint16ToBytes(uint16(tx.DestChainId))...)
	aggregatedPayload = append(aggregatedPayload, tx.ChannelId)
	aggregatedPayload = append(aggregatedPayload, util.Uint64ToBytes(tx.Sequence)...)
	aggregatedPayload = append(aggregatedPayload, uint8(tx.PackageType))
	aggregatedPayload = append(aggregatedPayload, util.Uint64ToBytes(uint64(tx.TxTime))...)

	// relayerfee big.Int
	relayerFeeBts, err := p.txFeeToBytes(tx.RelayerFee)
	if err != nil {
		relayercommon.Logger.Errorf("failed to convert tx relayerFee %s from string to big.Int", tx.AckRelayerFee)
		return nil, err
	}
	aggregatedPayload = append(aggregatedPayload, relayerFeeBts...)

	if tx.PackageType == uint32(sdk.SynCrossChainPackageType) {
		ackRelayerFeeBts, err := p.txFeeToBytes(tx.AckRelayerFee)
		if err != nil {
			relayercommon.Logger.Errorf("failed to convert tx ackRelayerFee %s from string to big.Int", tx.AckRelayerFee)
			return nil, err
		}
		aggregatedPayload = append(aggregatedPayload, ackRelayerFeeBts...)
	}
	aggregatedPayload = append(aggregatedPayload, common.Hex2Bytes(tx.PayLoad)...)
	return aggregatedPayload, nil
}

func (p *InscriptionVoteProcessor) txFeeToBytes(txFee string) ([]byte, error) {
	fee, ok := new(big.Int).SetString(txFee, 10)
	if !ok {
		return nil, errors.New("failed to convert tx fee")
	}
	feeBytes := make([]byte, 32)
	fee.FillBytes(feeBytes)
	return feeBytes, nil
}
