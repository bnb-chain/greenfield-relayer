package vote

import (
	"bytes"
	"context"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	"inscription-relayer/listener"
	"time"
)

type InscriptionVoteProcessor struct {
	votePoolExecutor    *executor.VotePoolExecutor
	daoManager          *dao.DaoManager
	config              *config.Config
	signer              *VoteSigner
	inscriptionExecutor *executor.InscriptionExecutor
}

func NewInscriptionVoteProcessor(cfg *config.Config, dao *dao.DaoManager, signer *VoteSigner, inscriptionExecutor *executor.InscriptionExecutor,
	votePoolExecutor *executor.VotePoolExecutor) *InscriptionVoteProcessor {
	return &InscriptionVoteProcessor{
		config:              cfg,
		daoManager:          dao,
		signer:              signer,
		inscriptionExecutor: inscriptionExecutor,
		votePoolExecutor:    votePoolExecutor,
	}
}

// SignAndBroadcast Will sign using the bls private key, broadcast the vote to votepool
func (p *InscriptionVoteProcessor) SignAndBroadcast() error {
	for {
		latestHeight, err := p.inscriptionExecutor.GetLatestBlockHeightWithRetry()

		if err != nil {
			relayercommon.Logger.Errorf("Failed to get latest block height, error: %s", err.Error())
			time.Sleep(listener.GetBlockHeightRetryInterval)
			continue
		}

		latestVotedTxHeight, err := p.daoManager.InscriptionDao.GetLatestVotedTransactionHeight()
		if err != nil {
			return err
		}
		if latestVotedTxHeight+p.config.InscriptionConfig.NumberOfBlocksForFinality > latestHeight {
			time.Sleep(2 * time.Second)
			continue
		}
		txs, err := p.daoManager.InscriptionDao.GetUnVotedTransactionsAtHeight(latestVotedTxHeight + 1)
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get transactions from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		if len(txs) == 0 {
			relayercommon.Logger.Errorf("Current block txs have been voted, block height: %d", latestVotedTxHeight+1)
			time.Sleep(2 * time.Second)
			continue
		}

		// for every tx, we are going to sign it and broadcast vote of it. 1 tx to 1 vote
		for _, tx := range txs {
			v, err := p.constructVoteAndSign(tx)
			if err != nil {
				return err
			}

			//broadcast v
			if err = retry.Do(func() error {
				err = p.votePoolExecutor.SubmitVote(v)
				if err != nil {
					return fmt.Errorf("failed to submit vote for event with txhash: %s", tx.TxHash)
				}
				return err
			}, retry.Context(context.Background()), relayercommon.RtyAttem, relayercommon.RtyDelay, relayercommon.RtyErr); err != nil {
				return err
			}

			//After vote submitted to vote pool, persist vote Data and update the status of tx to 'VOTED'.
			dbTx := p.daoManager.InscriptionDao.DB.Begin()

			err = p.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, VOTED)
			if err != nil {
				dbTx.Rollback()
				return err
			}
			voteData, err := p.constructVoteData(tx)
			if err != nil {
				return err
			}
			err = p.daoManager.VoteDao.SaveVoteData(voteData.ToDbModel())
			if err != nil {
				dbTx.Rollback()
				return err
			}
			return dbTx.Commit().Error
		}
	}
}

func (p *InscriptionVoteProcessor) CollectVotes() error {

	for {
		var err error
		txs, err := p.daoManager.InscriptionDao.GetVotedTransactions()
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get voted transactions from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}
		for _, tx := range txs {
			voteData, err := p.daoManager.VoteDao.GetVoteDataByChannelAndSequence(tx.ChannelId, tx.Sequence)
			if err != nil {
				return err
			}
			//this will keep query more 2/3 votes, and verify sig, if verifycation fails, will re-broadcast votes and go over again.
			//if less than 2/3, will

			var votes []*Vote
			p.prepareEnoughValidVotesForTransaction(votes, voteData, tx)

			modelVotes := make([]*model.Vote, 0, len(votes))

			for _, v := range votes {
				modelVote := v.ToDbModel(tx.ChannelId, tx.Sequence)
				modelVotes = append(modelVotes, modelVote)
			}
			dbTx := p.daoManager.VoteDao.DB.Begin()

			//Update tx status in DB to ALL_VOTED
			err = p.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, VOTED_ALL)
			if err != nil {
				dbTx.Rollback()
				return err
			}
			//Persist votes result for a tx to DB
			err = p.daoManager.VoteDao.SaveBatchVotes(modelVotes)
			if err != nil {
				dbTx.Rollback()
				return err
			}
			return dbTx.Commit().Error
		}
	}
}

// prepareEnoughValidVotesForTransaction will prepare fetch and validate votes result, store in votes
func (p *InscriptionVoteProcessor) prepareEnoughValidVotesForTransaction(votes []*Vote, voteData *model.VoteData, tx *model.InscriptionRelayTransaction) {
	for {
		validators, _ := p.votePoolExecutor.QueryValidators()
		validatorsSize := len(validators)
		//Query from votePool until there are more than 2/3 votes
		votes, _ = p.queryMoreThanTwoThirdVotesForTransaction(tx, validatorsSize)

		reQueryFromVotePool := false
		eventHash := voteData.EventHash
		voteSize := len(votes)
		relayercommon.Logger.Infof("Query %d votes from votepool for transaciton %s.", voteSize, tx.TxHash)
		//Verify all votes received,
		//1 verify sig.
		//2. verify pub key exists
		validVotesSize := voteSize

		for _, v := range votes {
			var err error
			isValidPubKey := false
			if err = v.Verify(eventHash); err != nil {
				relayercommon.Logger.Errorf("failed to verify vote, vote pub key is %s ", v.PubKey)
				validVotesSize--
				continue
			}
			for _, validator := range validators {
				// will check current vote pub key belong to any validator, if not found, means current vote is not valid
				if bytes.Equal(v.PubKey[:], validator.PubKey.Bytes()[:]) {
					isValidPubKey = true
				}
			}
			if !isValidPubKey {
				validVotesSize--
				relayercommon.Logger.Errorf("the vote's pub key '%s' does not belong to any validator", v.PubKey)
				continue
			}
			if validVotesSize < validVotesSize*2/3 {
				reQueryFromVotePool = true
				break
			}
		}
		if reQueryFromVotePool {
			time.Sleep(time.Second * 1)
			continue
		}
		return
	}
}

// queryMoreThanTwoThirdVotesForTransaction query votes from votePool
func (p *InscriptionVoteProcessor) queryMoreThanTwoThirdVotesForTransaction(tx *model.InscriptionRelayTransaction, validatorsSize int) ([]*Vote, error) {
	for {
		var err error

		eventHash, err := p.getEventHash(tx)
		if err != nil {
			return nil, err
		}

		votes, err := p.votePoolExecutor.QueryVotes(eventHash, ToBscCrossChainEvent)
		if err != nil {
			return nil, err
		}

		if len(votes) < validatorsSize*3/2 {
			//will re-broadcast if queried votes dont include validator's vote
			validatorVoteExist := false

			//re-broadcast current validator vote if votes result from votePool not exist
			for _, v := range votes {
				if bytes.Equal(v.PubKey[:], []byte(p.config.VotePoolConfig.BlsPublicKey)) {
					validatorVoteExist = true
					break
				}
			}
			if !validatorVoteExist {
				v, err := p.constructVoteAndSign(tx)
				if err != nil {
					return nil, err
				}
				err = p.votePoolExecutor.SubmitVote(v)
				if err != nil {
					return nil, err
				}
			}
			time.Sleep(time.Second * 1)
			continue
		}

		return votes, nil
	}
}

func (p *InscriptionVoteProcessor) constructVoteAndSign(tx *model.InscriptionRelayTransaction) (*Vote, error) {
	var v *Vote
	v.EvenType = ToBscCrossChainEvent
	eventHash, err := p.getEventHash(tx)
	if err != nil {
		return nil, err
	}
	err = p.signer.SignVote(v, eventHash)
	return v, err
}

func (p *InscriptionVoteProcessor) constructVoteData(tx *model.InscriptionRelayTransaction) (*VoteData, error) {
	eventHash, err := p.getEventHash(tx)
	if err != nil {
		return nil, err
	}
	return &VoteData{
		EventHash: eventHash,
		Sequence:  tx.Sequence,
		ChannelId: tx.ChannelId,
	}, nil
}

func (p *InscriptionVoteProcessor) getEventHash(tx *model.InscriptionRelayTransaction) ([]byte, error) {
	b, err := rlp.EncodeToBytes(tx.PayLoad)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256Hash(b).Bytes(), nil
}
