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

type BSCVoteProcessor struct {
	votePoolExecutor *VotePoolExecutor
	daoManager       *dao.DaoManager
	config           *config.Config
	signer           *VoteSigner
	bscExecutor      *executor.BSCExecutor
	blsPublicKey     []byte
}

func NewBSCVoteProcessor(cfg *config.Config, dao *dao.DaoManager, signer *VoteSigner, bscExecutor *executor.BSCExecutor,
	votePoolExecutor *VotePoolExecutor) (*BSCVoteProcessor, error) {

	pubKey, err := util.GetBlsPubKeyFromPrivKeyStr(cfg.VotePoolConfig.BlsPrivateKey)
	if err != nil {
		return nil, err
	}
	return &BSCVoteProcessor{
		config:           cfg,
		daoManager:       dao,
		signer:           signer,
		bscExecutor:      bscExecutor,
		votePoolExecutor: votePoolExecutor,
		blsPublicKey:     pubKey,
	}, nil
}

// SignAndBroadcast Will sign using the bls private key, and broadcast the vote to votepool
func (p *BSCVoteProcessor) SignAndBroadcast() error {

	for {
		latestHeight, err := p.bscExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
			time.Sleep(listener.GetBlockHeightRetryInterval)
			continue
		}

		leastSavedPkgHeight, err := p.daoManager.BSCDao.GetLeastSavedPackagesHeight()
		if err != nil {
			return err
		}

		if leastSavedPkgHeight+p.config.BSCConfig.NumberOfBlocksForFinality > latestHeight {
			time.Sleep(2 * time.Second)
			continue
		}
		pkgs, err := p.daoManager.BSCDao.GetPackagesByStatusAndHeight(model.SAVED, leastSavedPkgHeight)

		if err != nil {
			relayercommon.Logger.Errorf("Failed to get packages at height %d from db, error: %s", leastSavedPkgHeight, err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		if len(pkgs) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}

		//For packages with same oracle sequence, aggregate their payload and make single vote to votepool
		pkgsGroupByOracleSeq := make(map[uint64][]*model.BscRelayPackage)
		for _, tx := range pkgs {
			pkgsGroupByOracleSeq[tx.OracleSequence] = append(pkgsGroupByOracleSeq[tx.OracleSequence], tx)
		}

		for seq, pkgsForSeq := range pkgsGroupByOracleSeq {
			aggPkgs := make(Packages, 0)
			var txIds []int64

			for _, pkg := range pkgsForSeq {
				// aggregate pkgs with same oracle seq
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
			channelId := relayercommon.OracleChannelId

			v := p.constructVoteAndSign(encodedPackages)

			//broadcast v
			if err = retry.Do(func() error {
				err = p.votePoolExecutor.BroadcastVote(v)
				if err != nil {
					return fmt.Errorf("failed to submit vote for events with channel id %d and sequence %d", channelId, seq)
				}
				return nil
			}, retry.Context(context.Background()), relayercommon.RtyAttem, relayercommon.RtyDelay, relayercommon.RtyErr); err != nil {
				return err
			}

			err = p.daoManager.BSCDao.DB.Transaction(func(dbTx *gorm.DB) error {
				err := p.daoManager.BSCDao.UpdateBatchPackagesStatus(txIds, model.VOTED)
				if err != nil {
					return err
				}
				err = p.daoManager.VoteDao.SaveVote(FromEntityToDto(v, uint8(channelId), seq))
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

func (p *BSCVoteProcessor) CollectVotes() error {

	for {
		pkgs, err := p.daoManager.BSCDao.GetPackagesByStatus(model.VOTED)
		if err != nil {
			relayercommon.Logger.Errorf("failed to get voted packages from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		pkgsGroupByOracleSeq := make(map[uint64][]*model.BscRelayPackage)
		for _, pkg := range pkgs {
			pkgsGroupByOracleSeq[pkg.OracleSequence] = append(pkgsGroupByOracleSeq[pkg.OracleSequence], pkg)
		}

		for seq, pkgsForSeq := range pkgsGroupByOracleSeq {
			var txIds []int64
			oracleChannelId := relayercommon.OracleChannelId

			for _, tx := range pkgsForSeq {
				txIds = append(txIds, tx.Id)
			}

			err := p.prepareEnoughValidVotesForPackages(oracleChannelId, seq)
			if err != nil {
				return err
			}

			err = p.daoManager.BSCDao.UpdateBatchPackagesStatus(txIds, model.VOTED_ALL)
			if err != nil {
				return err
			}
		}
	}
}

// prepareEnoughValidVotesForPackages will prepare fetch and validate votes result, store in votes
func (p *BSCVoteProcessor) prepareEnoughValidVotesForPackages(channelId relayercommon.ChannelId, sequence uint64) error {
	localVote, err := p.daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(uint8(channelId), sequence, hex.EncodeToString(p.blsPublicKey))
	if err != nil {
		return err
	}

	validators, err := p.bscExecutor.InscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return err
	}
	//Query from votePool until there are more than 2/3 votes
	err = p.queryMoreThanTwoThirdValidVotes(localVote, validators)
	if err != nil {
		return err
	}
	return nil
}

// queryMoreThanTwoThirdValidVotes query votes from votePool
func (p *BSCVoteProcessor) queryMoreThanTwoThirdValidVotes(localVote *model.Vote, validators []stakingtypes.Validator) error {

	validVotesTotalCnt := 1 // assume local vote is valid
	channelId := localVote.ChannelId
	seq := localVote.Sequence
	for {
		queriedVotes, err := p.votePoolExecutor.QueryVotes(localVote.EventHash, votepool.FromBscCrossChainEvent)
		if err != nil {
			continue
		}

		validVotesCntPerReq := len(queriedVotes)

		if validVotesCntPerReq == 0 {
			continue
		}
		isLocalVoteIncluded := false

		for _, v := range queriedVotes {
			if !p.isVotePubKeyValid(v, validators) {
				relayercommon.Logger.Errorf("vote's pub-key %s does not belong to any validator", hex.EncodeToString(v.PubKey[:]))
				validVotesCntPerReq--
				continue
			}

			if err := Verify(v, localVote.EventHash); err != nil {
				relayercommon.Logger.Errorf("verify vote's signature failed,  err=%s", err)
				validVotesCntPerReq--
				continue
			}

			if bytes.Equal(v.PubKey[:], p.blsPublicKey) {
				isLocalVoteIncluded = true
				validVotesCntPerReq--
				continue
			}

			exist, err := p.daoManager.VoteDao.IsVoteExist(channelId, seq, hex.EncodeToString(v.PubKey[:]))
			if err != nil {
				return err
			}
			if exist {
				validVotesCntPerReq--
				continue
			}
			err = p.daoManager.VoteDao.SaveVote(FromEntityToDto(v, channelId, seq))
			if err != nil {
				return err
			}
		}

		validVotesTotalCnt += validVotesCntPerReq
		if validVotesTotalCnt < len(validators)*2/3 {
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
		return nil
	}
}

func (p *BSCVoteProcessor) constructVoteAndSign(eventHash []byte) *votepool.Vote {
	var v votepool.Vote
	v.EventType = votepool.FromBscCrossChainEvent
	p.signer.SignVote(&v, eventHash)
	return &v
}

func (p *BSCVoteProcessor) isVotePubKeyValid(v *votepool.Vote, validators []stakingtypes.Validator) bool {
	for _, validator := range validators {
		if bytes.Equal(v.PubKey[:], validator.RelayerBlsKey[:]) {
			return true
		}
	}
	return false
}
