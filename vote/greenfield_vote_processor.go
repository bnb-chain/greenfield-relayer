package vote

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/cometbft/cometbft/votepool"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"

	rcommon "github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/types"
	"github.com/bnb-chain/greenfield-relayer/util"
)

type GreenfieldVoteProcessor struct {
	daoManager         *dao.DaoManager
	config             *config.Config
	signer             *VoteSigner
	greenfieldExecutor *executor.GreenfieldExecutor
	blsPublicKey       []byte
}

func NewGreenfieldVoteProcessor(cfg *config.Config, dao *dao.DaoManager, signer *VoteSigner,
	greenfieldExecutor *executor.GreenfieldExecutor) *GreenfieldVoteProcessor {
	return &GreenfieldVoteProcessor{
		config:             cfg,
		daoManager:         dao,
		signer:             signer,
		greenfieldExecutor: greenfieldExecutor,
		blsPublicKey:       greenfieldExecutor.BlsPubKey,
	}
}

// SignAndBroadcastLoop signs tx using the relayer's bls private key, then broadcasts the vote to Greenfield votepool
func (p *GreenfieldVoteProcessor) SignAndBroadcastLoop() {
	ticker := time.NewTicker(time.Duration(p.config.VotePoolConfig.BroadcastIntervalInMillisecond) * time.Millisecond)
	for range ticker.C {
		if err := p.signAndBroadcast(); err != nil {
			logging.Logger.Errorf("encounter error, err: %s", err.Error())
		}
	}
}

func (p *GreenfieldVoteProcessor) signAndBroadcast() error {
	latestHeight, err := p.greenfieldExecutor.GetLatestBlockHeight()
	if err != nil {
		logging.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
		return err
	}

	leastSavedTxHeight, err := p.daoManager.GreenfieldDao.GetLeastSavedTransactionHeight()
	if err != nil {
		logging.Logger.Errorf("failed to get least saved tx height, error: %s", err.Error())
		return err
	}
	if leastSavedTxHeight+p.config.GreenfieldConfig.NumberOfBlocksForFinality > latestHeight {
		return nil
	}
	txs, err := p.daoManager.GreenfieldDao.GetTransactionsByStatusWithLimit(db.Saved, p.config.VotePoolConfig.VotesBatchMaxSizePerInterval)
	if err != nil {
		logging.Logger.Errorf("failed to get transactions from db, error: %s", err.Error())
		return err
	}
	if len(txs) == 0 {
		return nil
	}
	// for every tx, we are going to sign it and broadcast vote of it.
	for _, tx := range txs {

		// in case there is chance that reprocessing same transactions(caused by DB data loss) or processing outdated
		// transactions from block( when relayer need to catch up others), this ensures relayer will skip to next transaction directly
		isFilled, err := p.isTxSequenceFilled(tx)
		if err != nil {
			return err
		}
		if isFilled {
			if err = p.daoManager.GreenfieldDao.UpdateTransactionStatus(tx.Id, db.Delivered); err != nil {
				return err
			}
			logging.Logger.Infof("sequence %d for channel %d has already been filled ", tx.Sequence, tx.ChannelId)
			continue
		}

		aggregatedPayload, err := p.aggregatePayloadForTx(tx)
		if err != nil {
			return err
		}
		v := p.constructVoteAndSign(aggregatedPayload)

		// broadcast v
		if err = retry.Do(func() error {
			logging.Logger.Debugf("broadcasting vote with c %d and seq %d", tx.ChannelId, tx.Sequence)

			err = p.greenfieldExecutor.BroadcastVote(v)
			if err != nil {
				return fmt.Errorf("failed to submit vote for event with channel id %d and sequence %d, err=%s", tx.ChannelId, tx.Sequence, err.Error())
			}
			return nil
		}, retry.Context(context.Background()), rcommon.RtyAttem, rcommon.RtyDelay, rcommon.RtyErr); err != nil {
			return err
		}

		// After vote submitted to vote pool, persist vote Data and update the status of tx to 'SELF_VOTED'.
		err = p.daoManager.GreenfieldDao.DB.Transaction(func(dbTx *gorm.DB) error {
			if e := dao.UpdateTransactionStatus(dbTx, tx.Id, db.SelfVoted); e != nil {
				return e
			}
			exist, e := dao.IsVoteExist(dbTx, tx.ChannelId, tx.Sequence, hex.EncodeToString(v.PubKey[:]))
			if e != nil {
				return e
			}
			if !exist {
				if e = dao.SaveVote(dbTx, EntityToDto(v, tx.ChannelId, tx.Sequence, aggregatedPayload)); e != nil {
					return e
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *GreenfieldVoteProcessor) CollectVotesLoop() {
	ticker := time.NewTicker(time.Duration(p.config.VotePoolConfig.QueryIntervalInMillisecond) * time.Millisecond)
	for range ticker.C {
		if err := p.collectVotes(); err != nil {
			logging.Logger.Errorf("encounter error, err: %s", err.Error())
		}
	}
}

func (p *GreenfieldVoteProcessor) collectVotes() error {
	txs, err := p.daoManager.GreenfieldDao.GetTransactionsByStatusWithLimit(db.SelfVoted, p.config.VotePoolConfig.VotesBatchMaxSizePerInterval)
	if err != nil {
		logging.Logger.Errorf("failed to get voted transactions from db, error: %s", err.Error())
		return err
	}
	wg := new(sync.WaitGroup)
	errCh := make(chan error)
	waitCh := make(chan struct{})
	go func() {
		for _, tx := range txs {
			wg.Add(1)
			go p.collectVoteForTx(tx, errCh, wg)
		}
		wg.Wait()
		close(waitCh)
	}()
	for {
		select {
		case err := <-errCh:
			return err
		case <-waitCh:
			return nil
		}
	}
}

func (p *GreenfieldVoteProcessor) collectVoteForTx(tx *model.GreenfieldRelayTransaction, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	isFilled, err := p.isTxSequenceFilled(tx)
	if err != nil {
		errChan <- err
		return
	}
	if isFilled {
		if err = p.daoManager.GreenfieldDao.UpdateTransactionStatus(tx.Id, db.Delivered); err != nil {
			errChan <- err
			return
		}
		logging.Logger.Infof("sequence %d for channel %d has already been filled ", tx.Sequence, tx.ChannelId)
		return
	}

	if err = p.prepareEnoughValidVotesForTx(tx); err != nil {
		errChan <- err
		return
	}
	if err = p.daoManager.GreenfieldDao.UpdateTransactionStatus(tx.Id, db.AllVoted); err != nil {
		errChan <- err
		return
	}
}

// prepareEnoughValidVotesForTx fetches and validate votes result, store in vote table
func (p *GreenfieldVoteProcessor) prepareEnoughValidVotesForTx(tx *model.GreenfieldRelayTransaction) error {
	localVote, err := p.daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(tx.ChannelId, tx.Sequence, hex.EncodeToString(p.blsPublicKey))
	if err != nil {
		return err
	}

	validators, err := p.greenfieldExecutor.BscExecutor.QueryCachedLatestValidators()
	if err != nil {
		return err
	}

	count, err := p.daoManager.VoteDao.GetVotesCountByChannelIdAndSequence(tx.ChannelId, tx.Sequence)
	if err != nil {
		return err
	}
	if count > int64(len(validators))*2/3 {
		return nil
	}

	if err = p.queryMoreThanTwoThirdVotesForTx(localVote, validators); err != nil {
		return err
	}
	return nil
}

// queryMoreThanTwoThirdVotesForTx queries votes from votePool
func (p *GreenfieldVoteProcessor) queryMoreThanTwoThirdVotesForTx(localVote *model.Vote, validators []types.Validator) error {
	triedTimes := 0
	validVotesTotalCount := 1 // assume local vote is valid
	channelId := localVote.ChannelId
	seq := localVote.Sequence
	ticker := time.NewTicker(VotePoolQueryRetryInterval)

	for range ticker.C {
		triedTimes++
		if triedTimes > QueryVotepoolMaxRetryTimes {
			return errors.New("exceed max retry")
		}

		logging.Logger.Debugf("query vote for c %d and s %d", channelId, seq)
		queriedVotes, err := p.greenfieldExecutor.QueryVotesByEventHashAndType(localVote.EventHash, votepool.ToBscCrossChainEvent)
		if err != nil {
			return err
		}
		validVotesCountPerReq := len(queriedVotes)
		if validVotesCountPerReq == 0 {
			if err := p.reBroadcastVote(localVote); err != nil {
				return err
			}
			continue
		}
		isLocalVoteIncluded := false

		for _, v := range queriedVotes {

			if !p.isVotePubKeyValid(v, validators) {
				validVotesCountPerReq--
				continue
			}

			if err := VerifySignature(v, localVote.EventHash); err != nil {
				validVotesCountPerReq--
				continue
			}

			// check if it is local vote
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
			if err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, channelId, seq, localVote.ClaimPayload)); err != nil {
				return err
			}
		}

		validVotesTotalCount += validVotesCountPerReq

		if validVotesTotalCount > len(validators)*2/3 {
			return nil
		}

		if !isLocalVoteIncluded {
			if err := p.reBroadcastVote(localVote); err != nil {
				return err
			}
		}
		continue
	}
	return nil
}

func (p *GreenfieldVoteProcessor) constructVoteAndSign(aggregatedPayload []byte) *votepool.Vote {
	var v votepool.Vote
	v.EventType = votepool.ToBscCrossChainEvent
	v.EventHash = p.getEventHash(aggregatedPayload)
	p.signer.SignVote(&v)
	return &v
}

func (p *GreenfieldVoteProcessor) getEventHash(aggregatedPayload []byte) []byte {
	return crypto.Keccak256Hash(aggregatedPayload).Bytes()
}

func (p *GreenfieldVoteProcessor) isVotePubKeyValid(v *votepool.Vote, validators []types.Validator) bool {
	for _, validator := range validators {
		if bytes.Equal(v.PubKey[:], validator.BlsPublicKey[:]) {
			return true
		}
	}
	return false
}

// aggregatePayloadForTx aggregate required fields by concatenating their bytes, this will be used as payload when
// calling BSC smart contract, and also used to generate eventHash for broadcasting vote
func (p *GreenfieldVoteProcessor) aggregatePayloadForTx(tx *model.GreenfieldRelayTransaction) ([]byte, error) {
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
		logging.Logger.Errorf("failed to convert tx relayerFee %s from string to big.Int", tx.RelayerFee)
		return nil, err
	}
	aggregatedPayload = append(aggregatedPayload, relayerFeeBts...)

	if tx.PackageType == uint32(sdk.SynCrossChainPackageType) {
		ackRelayerFeeBts, err := p.txFeeToBytes(tx.AckRelayerFee)
		if err != nil {
			logging.Logger.Errorf("failed to convert tx ackRelayerFee %s from string to big.Int", tx.AckRelayerFee)
			return nil, err
		}
		aggregatedPayload = append(aggregatedPayload, ackRelayerFeeBts...)
	}
	aggregatedPayload = append(aggregatedPayload, common.Hex2Bytes(tx.PayLoad)...)
	return aggregatedPayload, nil
}

func (p *GreenfieldVoteProcessor) txFeeToBytes(txFee string) ([]byte, error) {
	fee, ok := new(big.Int).SetString(txFee, 10)
	if !ok {
		return nil, errors.New("failed to convert tx fee")
	}
	feeBytes := make([]byte, 32)
	fee.FillBytes(feeBytes)
	return feeBytes, nil
}

func (p *GreenfieldVoteProcessor) isTxSequenceFilled(tx *model.GreenfieldRelayTransaction) (bool, error) {
	nextDeliverySequence, err := p.greenfieldExecutor.GetNextDeliverySequenceForChannelWithRetry(types.ChannelId(tx.ChannelId))
	if err != nil {
		return false, err
	}
	return tx.Sequence < nextDeliverySequence, nil
}

func (p *GreenfieldVoteProcessor) reBroadcastVote(localVote *model.Vote) error {
	logging.Logger.Debugf("broadcasting vote with c %d and seq %d", localVote.ChannelId, localVote.Sequence)

	v, err := DtoToEntity(localVote)
	if err != nil {
		return err
	}
	return p.greenfieldExecutor.BroadcastVote(v)
}
