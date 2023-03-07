package vote

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/avast/retry-go/v4"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/ethereum/go-ethereum/rlp"
	tmtypes "github.com/tendermint/tendermint/types"
	"github.com/tendermint/tendermint/votepool"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/types"
	"github.com/bnb-chain/greenfield-relayer/util"
)

type BSCVoteProcessor struct {
	daoManager   *dao.DaoManager
	config       *config.Config
	signer       *VoteSigner
	bscExecutor  *executor.BSCExecutor
	blsPublicKey []byte
}

func NewBSCVoteProcessor(cfg *config.Config, dao *dao.DaoManager, signer *VoteSigner, bscExecutor *executor.BSCExecutor) *BSCVoteProcessor {
	return &BSCVoteProcessor{
		config:       cfg,
		daoManager:   dao,
		signer:       signer,
		bscExecutor:  bscExecutor,
		blsPublicKey: util.BlsPubKeyFromPrivKeyStr(cfg.GreenfieldConfig.BlsPrivateKey),
	}
}

func (p *BSCVoteProcessor) SignAndBroadcastVoteLoop() {
	for {
		if err := p.signAndBroadcast(); err != nil {
			logging.Logger.Errorf("encounter error, err: %s", err.Error())
			time.Sleep(RetryInterval)
		}
	}
}

// SignAndBroadcastVoteLoop Will sign using the bls private key, and broadcast the vote to votepool
func (p *BSCVoteProcessor) signAndBroadcast() error {
	latestHeight, err := p.bscExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		logging.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
		return err
	}

	leastSavedPkgHeight, err := p.daoManager.BSCDao.GetLeastSavedPackagesHeight()
	if err != nil {
		logging.Logger.Errorf("failed to get least saved packages' height, error: %s", err.Error())
		return err
	}

	if leastSavedPkgHeight+p.config.BSCConfig.NumberOfBlocksForFinality > latestHeight {
		return nil
	}
	pkgs, err := p.daoManager.BSCDao.GetPackagesByStatusAndHeight(db.Saved, leastSavedPkgHeight)
	if err != nil {
		logging.Logger.Errorf("failed to get packages at height %d from db, error: %s", leastSavedPkgHeight, err.Error())
		return err
	}

	if len(pkgs) == 0 {
		return nil
	}

	// For packages with same oracle sequence, aggregate their payload and make single vote to votepool
	pkgsGroupByOracleSeq := make(map[uint64][]*model.BscRelayPackage)
	for _, pack := range pkgs {
		pkgsGroupByOracleSeq[pack.OracleSequence] = append(pkgsGroupByOracleSeq[pack.OracleSequence], pack)
	}

	for seq, pkgsForSeq := range pkgsGroupByOracleSeq {
		aggPkgs := make(oracletypes.Packages, 0)
		var pkgIds []int64

		sort.Slice(pkgsForSeq, func(i, j int) bool {
			return pkgsForSeq[i].TxIndex < pkgsForSeq[j].TxIndex
		})
		for _, pkg := range pkgsForSeq {
			// aggregate pkgs with same oracle seq
			payload, err := hex.DecodeString(pkg.PayLoad)
			if err != nil {
				return fmt.Errorf("decode payload error, payload=%s", pkg.PayLoad)
			}

			pack := oracletypes.Package{
				ChannelId: sdk.ChannelID(pkg.ChannelId),
				Sequence:  pkg.PackageSequence,
				Payload:   payload,
			}
			aggPkgs = append(aggPkgs, pack)
			pkgIds = append(pkgIds, pkg.Id)
		}

		// check if oracle sequence is filled on greenfield, if so, update packages status to filled and skip to next oracle sequence
		nextDeliverySeqOnGreenfield, err := p.bscExecutor.GetNextDeliveryOracleSequence()
		if err != nil {
			return err
		}
		if seq < nextDeliverySeqOnGreenfield {
			err = p.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.Delivered)
			if err != nil {
				logging.Logger.Errorf("failed to update packages error %s", err.Error())
				return err
			}
			logging.Logger.Infof("packages' oracle sequence %d is less than nex delivery oracle sequence %d", seq, nextDeliverySeqOnGreenfield)
			continue
		}

		encodedPayload, err := rlp.EncodeToBytes(aggPkgs)
		blsClaim := oracletypes.BlsClaim{
			// chain ids are validated when packages persisted into DB, non-matched ones would be omitted
			SrcChainId:  uint32(p.config.BSCConfig.ChainId),
			DestChainId: uint32(p.config.GreenfieldConfig.ChainId),
			Timestamp:   uint64(pkgsForSeq[0].TxTime),
			Sequence:    seq,
			Payload:     encodedPayload,
		}
		eventHash := blsClaim.GetSignBytes()
		if err != nil {
			return fmt.Errorf("encode packages error, err=%s", err.Error())
		}
		channelId := common.OracleChannelId
		v := p.constructSignedVote(eventHash[:])

		// broadcast v
		if err = retry.Do(func() error {
			err = p.bscExecutor.GreenfieldExecutor.BroadcastVote(v)
			if err != nil {
				return fmt.Errorf("failed to submit vote for events with channel id %d and sequence %d", channelId, seq)
			}
			return nil
		}, retry.Context(context.Background()), common.RtyAttem, common.RtyDelay, common.RtyErr); err != nil {
			return err
		}

		err = p.daoManager.BSCDao.DB.Transaction(func(dbTx *gorm.DB) error {
			err := p.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.SelfVoted)
			if err != nil {
				return err
			}
			exist, err := p.daoManager.VoteDao.IsVoteExist(uint8(channelId), seq, hex.EncodeToString(v.PubKey[:]))
			if err != nil {
				return err
			}
			if !exist {
				err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, uint8(channelId), seq, encodedPayload))
				if err != nil {
					return err
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

func (p *BSCVoteProcessor) CollectVotesLoop() {
	for {
		if err := p.collectVotes(); err != nil {
			logging.Logger.Errorf("encounter error, err: %s", err.Error())
			time.Sleep(RetryInterval)
		}
	}
}

func (p *BSCVoteProcessor) collectVotes() error {
	pkgs, err := p.daoManager.BSCDao.GetPackagesByStatus(db.SelfVoted)
	if err != nil {
		logging.Logger.Errorf("failed to get voted packages from db, error: %s", err.Error())
		return err
	}

	pkgsGroupByOracleSeq := make(map[uint64][]*model.BscRelayPackage)
	for _, pkg := range pkgs {
		pkgsGroupByOracleSeq[pkg.OracleSequence] = append(pkgsGroupByOracleSeq[pkg.OracleSequence], pkg)
	}

	for seq, pkgsForSeq := range pkgsGroupByOracleSeq {
		var pkgIds []int64
		oracleChannelId := common.OracleChannelId

		for _, tx := range pkgsForSeq {
			pkgIds = append(pkgIds, tx.Id)
		}

		err := p.prepareEnoughValidVotesForPackages(oracleChannelId, seq, pkgIds)
		if err != nil {
			return err
		}

		err = p.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.AllVoted)
		if err != nil {
			return err
		}
	}
	return nil
}

// prepareEnoughValidVotesForPackages will prepare fetch and validate votes result, store in votes
func (p *BSCVoteProcessor) prepareEnoughValidVotesForPackages(channelId types.ChannelId, sequence uint64, pkgIds []int64) error {
	localVote, err := p.daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(uint8(channelId), sequence, hex.EncodeToString(p.blsPublicKey))
	if err != nil {
		return err
	}
	validators, err := p.bscExecutor.GreenfieldExecutor.QueryCachedLatestValidators()
	if err != nil {
		return err
	}
	// Query from votePool until there are more than 2/3 votes
	if err = p.queryMoreThanTwoThirdValidVotes(localVote, validators, pkgIds); err != nil {
		return err
	}
	return nil
}

// queryMoreThanTwoThirdValidVotes queries votes from votePool
func (p *BSCVoteProcessor) queryMoreThanTwoThirdValidVotes(localVote *model.Vote, validators []*tmtypes.Validator, pkgIds []int64) error {
	triedTimes := 0
	validVotesTotalCnt := 1
	channelId := localVote.ChannelId
	seq := localVote.Sequence
	ticker := time.NewTicker(VotePoolQueryRetryInterval)
	for range ticker.C {
		triedTimes++
		if triedTimes > QueryVotepoolMaxRetryTimes {
			if err := p.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.Saved); err != nil {
				logging.Logger.Errorf("failed to update packages status to 'Saved', packages' id=%v", pkgIds)
				return err
			}
			return errors.New("exceed max retry")
		}
		queriedVotes, err := p.bscExecutor.GreenfieldExecutor.QueryVotesByEventHashAndType(localVote.EventHash, votepool.FromBscCrossChainEvent)
		if err != nil {
			logging.Logger.Errorf("encounter error when query votes.")
			return err
		}

		validVotesCntPerReq := len(queriedVotes)

		if validVotesCntPerReq == 0 {
			continue
		}
		isLocalVoteIncluded := false

		for _, v := range queriedVotes {
			if !p.isVotePubKeyValid(v, validators) {
				validVotesCntPerReq--
				continue
			}

			if err := VerifySignature(v, localVote.EventHash[:]); err != nil {
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
			if err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, channelId, seq, localVote.ClaimPayload)); err != nil {
				return err
			}
		}

		validVotesTotalCnt += validVotesCntPerReq

		if validVotesTotalCnt > len(validators)*2/3 {
			return nil
		}
		if !isLocalVoteIncluded {
			v, err := DtoToEntity(localVote)
			if err != nil {
				return err
			}
			if err = p.bscExecutor.GreenfieldExecutor.BroadcastVote(v); err != nil {
				return err
			}
		}
		continue
	}
	return nil
}

func (p *BSCVoteProcessor) constructSignedVote(eventHash []byte) *votepool.Vote {
	var v votepool.Vote
	v.EventType = votepool.FromBscCrossChainEvent
	v.EventHash = eventHash
	p.signer.SignVote(&v)
	return &v
}

func (p *BSCVoteProcessor) isVotePubKeyValid(v *votepool.Vote, validators []*tmtypes.Validator) bool {
	for _, validator := range validators {
		if bytes.Equal(v.PubKey[:], validator.RelayerBlsKey[:]) {
			return true
		}
	}
	return false
}
