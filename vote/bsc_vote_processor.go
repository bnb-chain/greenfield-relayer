package vote

import (
	"bytes"
	"context"
	"encoding/hex"
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

type BSCVoteProcessor struct {
	votePoolExecutor *executor.VotePoolExecutor
	daoManager       *dao.DaoManager
	config           *config.Config
	signer           *VoteSigner
	bscExecutor      *executor.BSCExecutor
}

func NewBSCVoteProcessor(cfg *config.Config, dao *dao.DaoManager, signer *VoteSigner, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *executor.VotePoolExecutor) *BSCVoteProcessor {
	return &BSCVoteProcessor{
		config:           cfg,
		daoManager:       dao,
		signer:           signer,
		bscExecutor:      bscExecutor,
		votePoolExecutor: votePoolExecutor,
	}
}

// SignedAndBroadcast Will sign using the bls private key, broadcast the vote to votepool
func (p *BSCVoteProcessor) SignedAndBroadcast() error {
	for {
		latestHeight, err := p.bscExecutor.GetLatestBlockHeightWithRetry()

		if err != nil {
			relayercommon.Logger.Errorf("Failed to get latest block height, error: %s", err.Error())
			time.Sleep(listener.GetBlockHeightRetryInterval)
			continue
		}

		latestVotedTxHeight, err := p.daoManager.BSCDao.GetLatestVotedPackagesHeight()
		if err != nil {
			return err
		}
		if latestVotedTxHeight+p.config.BSCConfig.NumberOfBlocksForFinality > latestHeight {
			time.Sleep(2 * time.Second)
			continue
		}
		txs, err := p.daoManager.BSCDao.GetUnVotedPackagesAtHeight(latestVotedTxHeight + 1)
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get packages from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		if len(txs) == 0 {
			relayercommon.Logger.Errorf("Current block txs have been voted, block height: %d", latestVotedTxHeight+1)
			time.Sleep(2 * time.Second)
			continue
		}

		//For packages with same oracle sequence, aggregate their payload and make single vote to votepool
		pkgsGroupByOracleSeq := make(map[uint64][]*model.BscRelayPackage)
		for _, tx := range txs {
			pkgsGroupByOracleSeq[tx.OracleSequence] = append(pkgsGroupByOracleSeq[tx.OracleSequence], tx)
		}

		for seq, pkgs := range pkgsGroupByOracleSeq {
			if len(pkgs) == 0 {
				continue
			}
			aggPkgs := make(Packages, 0)
			var txIds []int64

			for _, pkg := range pkgs {
				// aggregate txs with same oracle seq
				payload, err := hex.DecodeString(pkg.PayLoad)
				if err != nil {
					return fmt.Errorf("decode payload error, payload=%s", pkg.PayLoad)
				}
				newP := Package{
					ChannelId: pkg.ChannelId,
					Sequence:  pkg.OracleSequence,
					Payload:   payload, // aggregate payload to be signed
				}
				//voteDataPayloadBytes = append(voteDataPayloadBytes, payload...)
				aggPkgs = append(aggPkgs, newP)
				txIds = append(txIds, pkg.Id)
			}

			encBts, err := rlp.EncodeToBytes(aggPkgs)
			encodedPackages := crypto.Keccak256Hash(encBts).Bytes()

			if err != nil {
				return fmt.Errorf("encode packages error, err=%s", err.Error())
			}
			channelId := aggPkgs[0].ChannelId

			voteData := p.constructVoteData(seq, channelId, encodedPackages)

			var v *Vote
			err = p.constructVoteAndSign(v, encodedPackages)

			if err != nil {
				return err
			}
			//broadcast v
			if err = retry.Do(func() error {
				err = p.votePoolExecutor.SubmitVote(v)
				if err != nil {
					return fmt.Errorf("failed to submit vote for events with channel id %d and sequence %d", channelId, seq)
				}
				return nil
			}, retry.Context(context.Background()), relayercommon.RtyAttem, relayercommon.RtyDelay, relayercommon.RtyErr); err != nil {
				return err
			}

			//Update packages with same seq status to voted in DB
			dbTx := p.daoManager.BSCDao.DB.Begin()

			err = p.daoManager.BSCDao.UpdateBatchPackagesStatus(txIds, VOTED)
			if err != nil {
				dbTx.Rollback()
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

func (p *BSCVoteProcessor) CollectVotes() error {

	for {
		var err error
		txs, err := p.daoManager.BSCDao.GetVotedPackages()
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get voted packages from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		pkgsGroupByOracleSeq := make(map[uint64][]*model.BscRelayPackage)
		for _, tx := range txs {
			pkgsGroupByOracleSeq[tx.OracleSequence] = append(pkgsGroupByOracleSeq[tx.OracleSequence], tx)
		}
		for seq, pkgs := range pkgsGroupByOracleSeq {
			var txIds []int64
			channelId := pkgs[0].ChannelId

			for _, tx := range pkgs {
				txIds = append(txIds, tx.Id)
			}

			voteData, err := p.daoManager.VoteDao.GetVoteDataByChannelAndSequence(channelId, seq)
			if err != nil {
				return err
			}

			//this will keep query more 2/3 votes, and verify sig, if verifycation fails, will re-broadcast votes and go over again.
			//if less than 2/3, will
			var votes []*Vote
			p.prepareEnoughValidVotesForPackages(votes, voteData)

			err = p.saveVotes(votes, voteData, txIds)
			if err != nil {
				return err
			}

		}
	}
}

func (p *BSCVoteProcessor) saveVotes(votes []*Vote, voteData *model.VoteData, txIds []int64) error {
	//Votes to be persisted into DB
	vs := make([]*model.Vote, 0, len(votes))

	for _, v := range votes {
		modelVote := v.ToDbModel(voteData.ChannelId, voteData.Sequence)
		vs = append(vs, modelVote)
	}
	dbTx := p.daoManager.VoteDao.DB.Begin()

	//Update tx status in DB to ALL_VOTED
	err := p.daoManager.BSCDao.UpdateBatchPackagesStatus(txIds, VOTED_ALL)
	if err != nil {
		dbTx.Rollback()
		return err
	}
	//Persist votes result for txs to DB
	err = p.daoManager.VoteDao.SaveBatchVotes(vs)
	if err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

// prepareEnoughValidVotesForPackages will prepare fetch and validate votes result, store in votes
func (p *BSCVoteProcessor) prepareEnoughValidVotesForPackages(votes []*Vote, voteData *model.VoteData) {
	for {
		validators, _ := p.votePoolExecutor.QueryValidators()
		validatorsSize := len(validators)
		//Query from votePool until there are more than 2/3 votes
		votes, _ = p.queryMoreThanTwoThirdVotesForPackages(voteData.ChannelId, voteData.Sequence, validatorsSize)

		reQueryFromVotePool := false
		eventHash := voteData.EventHash
		voteSize := len(votes)
		relayercommon.Logger.Infof("Query %d votes from votepool with eventHash %s.", voteSize, eventHash)
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

// queryMoreThanTwoThirdVotesForPackages query votes from votePool
func (p *BSCVoteProcessor) queryMoreThanTwoThirdVotesForPackages(channelId uint8, sequence uint64, validatorsSize int) ([]*Vote, error) {
	for {
		var err error
		//Query votes by eventHash and event type
		voteData, err := p.daoManager.VoteDao.GetVoteDataByChannelAndSequence(channelId, sequence)
		if err != nil {
			return nil, err
		}
		votes, err := p.votePoolExecutor.QueryVotes(voteData.EventHash, FromBscCrossChainEvent)
		if err != nil {
			return nil, err
		}

		if len(votes) < validatorsSize*2/3 {
			//will re-broadcast if dont include valdator itself
			validatorVoteExist := false

			//re-broadcast current validator vote if votes result from votePool not exist
			for _, v := range votes {
				if bytes.Equal(v.PubKey[:], []byte(p.config.VotePoolConfig.BlsPublicKey)[:]) {
					validatorVoteExist = true
					break
				}
			}
			if !validatorVoteExist {
				var v *Vote
				voteData, err = p.daoManager.VoteDao.GetVoteDataByChannelAndSequence(channelId, sequence)
				if err != nil {
					return nil, fmt.Errorf("failed to get vote Data from db, error: %s", err.Error())
				}
				err = p.constructVoteAndSign(v, voteData.EventHash)
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

func (p *BSCVoteProcessor) constructVoteAndSign(v *Vote, eventHash []byte) error {
	v.EvenType = FromBscCrossChainEvent
	err := p.signer.SignVote(v, eventHash)
	return err
}

func (p *BSCVoteProcessor) constructVoteData(oracleSequence uint64, channelId uint8, eventHash []byte) *VoteData {
	return &VoteData{
		EventHash: eventHash,
		Sequence:  oracleSequence,
		ChannelId: channelId,
	}
}
