package vote

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/avast/retry-go/v4"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/jinzhu/gorm"
	"github.com/tendermint/tendermint/votepool"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	"inscription-relayer/listener"
	"inscription-relayer/util"
	"time"
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
	votePoolExecutor *VotePoolExecutor) (*InscriptionVoteProcessor, error) {
	pubKey, err := util.GetBlsPubKeyFromPrivKeyStr(cfg.VotePoolConfig.BlsPrivateKey)
	if err != nil {
		return nil, err
	}
	return &InscriptionVoteProcessor{
		config:              cfg,
		daoManager:          dao,
		signer:              signer,
		inscriptionExecutor: inscriptionExecutor,
		votePoolExecutor:    votePoolExecutor,
		blsPublicKey:        pubKey,
	}, nil
}

// SignAndBroadcast Will sign using the bls private key, broadcast the vote to votepool
func (p *InscriptionVoteProcessor) SignAndBroadcast() error {

	for {
		latestHeight, err := p.inscriptionExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
			time.Sleep(listener.GetBlockHeightRetryInterval)
			continue
		}

		leastSavedTxHeight, err := p.daoManager.InscriptionDao.GetLeastSavedTxHeight()
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		if leastSavedTxHeight+p.config.InscriptionConfig.NumberOfBlocksForFinality > latestHeight {
			time.Sleep(2 * time.Second)
			continue
		}
		txs, err := p.daoManager.InscriptionDao.GetTransactionsByStatusAndHeight(model.SAVED, leastSavedTxHeight)
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get transactions at height %d from db, error: %s", leastSavedTxHeight, err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		if len(txs) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}

		// for every tx, we are going to sign it and broadcast vote of it.
		for _, tx := range txs {
			v, err := p.constructVoteAndSign(tx)
			if err != nil {
				return err
			}

			//broadcast v
			if err = retry.Do(func() error {
				err = p.votePoolExecutor.BroadcastVote(v)
				if err != nil {
					return fmt.Errorf("failed to submit vote for event with txhash: %s", tx.TxHash)
				}
				return nil
			}, retry.Context(context.Background()), relayercommon.RtyAttem, relayercommon.RtyDelay, relayercommon.RtyErr); err != nil {
				return err
			}

			//After vote submitted to vote pool, persist vote Data and update the status of tx to 'VOTED'.
			err = p.daoManager.InscriptionDao.DB.Transaction(func(dbTx *gorm.DB) error {
				err = p.daoManager.InscriptionDao.UpdateTxStatus(tx.Id, model.VOTED)
				if err != nil {
					return err
				}
				err = p.daoManager.VoteDao.SaveVote(FromEntityToDto(v, tx.ChannelId, tx.Sequence))
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}
}

func (p *InscriptionVoteProcessor) CollectVotes() error {

	for {
		txs, err := p.daoManager.InscriptionDao.GetTransactionsByStatus(model.VOTED)
		if err != nil {
			relayercommon.Logger.Errorf("failed to get voted transactions from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}
		for _, tx := range txs {
			err := p.prepareEnoughValidVotesForTx(tx)
			if err != nil {
				return err
			}

			err = p.daoManager.InscriptionDao.UpdateTxStatus(tx.Id, model.VOTED_ALL)
			if err != nil {
				return err
			}
		}
	}
}

// prepareEnoughValidVotesForTx will prepare fetch and validate votes result, store in votes
func (p *InscriptionVoteProcessor) prepareEnoughValidVotesForTx(tx *model.InscriptionRelayTransaction) error {

	validators, err := p.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return err
	}
	//Query from votePool until there are more than 2/3 votes
	err = p.queryMoreThanTwoThirdVotesForTx(tx, validators)

	if err != nil {
		return err
	}
	return nil
}

// queryMoreThanTwoThirdVotesForTx query votes from votePool
func (p *InscriptionVoteProcessor) queryMoreThanTwoThirdVotesForTx(tx *model.InscriptionRelayTransaction, validators []stakingtypes.Validator) error {

	validVotesTotalCount := 1 // assume local vote is valid
	channelId := tx.ChannelId
	seq := tx.Sequence
	localVote, err := p.constructVoteAndSign(tx)
	if err != nil {
		return err
	}
	for {
		queriedVotes, err := p.votePoolExecutor.QueryVotes(localVote.EventHash, votepool.ToBscCrossChainEvent)
		if err != nil {
			continue
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

			if err := Verify(v, localVote.EventHash); err != nil {
				relayercommon.Logger.Errorf("verify vote's signature failed,  err=%s", err)
				validVotesCountPerReq--
				continue
			}

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
			err = p.daoManager.VoteDao.SaveVote(FromEntityToDto(v, channelId, seq))
			if err != nil {
				return err
			}
		}

		validVotesTotalCount += validVotesCountPerReq
		if validVotesTotalCount < len(validators)*2/3 {
			if !isLocalVoteIncluded {
				err := p.votePoolExecutor.BroadcastVote(localVote)
				if err != nil {
					return err
				}
			}
			continue
		}
		return nil
	}
}

func (p *InscriptionVoteProcessor) constructVoteAndSign(tx *model.InscriptionRelayTransaction) (*votepool.Vote, error) {
	var v votepool.Vote
	v.EventType = votepool.ToBscCrossChainEvent
	eventHash, err := p.getEventHash(tx)
	if err != nil {
		return nil, err
	}
	p.signer.SignVote(&v, eventHash)
	return &v, nil
}

func (p *InscriptionVoteProcessor) getEventHash(tx *model.InscriptionRelayTransaction) ([]byte, error) {
	b, err := rlp.EncodeToBytes(tx.PayLoad)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(b).Bytes(), nil
}

func (p *InscriptionVoteProcessor) isVotePubKeyValid(v *votepool.Vote, validators []stakingtypes.Validator) bool {
	for _, validator := range validators {
		if bytes.Equal(v.PubKey[:], validator.RelayerBlsKey[:]) {
			return true
		}
	}
	return false
}
