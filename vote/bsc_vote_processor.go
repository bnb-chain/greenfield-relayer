package vote

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/avast/retry-go/v4"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tendermint/tendermint/votepool"
	"gorm.io/gorm"
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
	votePoolExecutor *VotePoolExecutor,
) *BSCVoteProcessor {
	return &BSCVoteProcessor{
		config:           cfg,
		daoManager:       dao,
		signer:           signer,
		bscExecutor:      bscExecutor,
		votePoolExecutor: votePoolExecutor,
		blsPublicKey:     util.GetBlsPubKeyFromPrivKeyStr(cfg.VotePoolConfig.BlsPrivateKey),
	}
}

func (p *BSCVoteProcessor) SignAndBroadcast() {
	for {
		err := p.signAndBroadcast()
		if err != nil {
			time.Sleep(RetryInterval)
		}
	}
}

// SignAndBroadcast Will sign using the bls private key, and broadcast the vote to votepool
func (p *BSCVoteProcessor) signAndBroadcast() error {
	latestHeight, err := p.bscExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
		return err
	}

	leastSavedPkgHeight, err := p.daoManager.BSCDao.GetLeastSavedPackagesHeight()
	if err != nil {
		return err
	}

	if leastSavedPkgHeight+p.config.BSCConfig.NumberOfBlocksForFinality > latestHeight {
		return nil
	}
	pkgs, err := p.daoManager.BSCDao.GetPackagesByStatusAndHeight(db.Saved, leastSavedPkgHeight)
	if err != nil {
		relayercommon.Logger.Errorf("failed to get packages at height %d from db, error: %s", leastSavedPkgHeight, err.Error())
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

		// check if oracle sequence is filled on inscription, if so, update packages status to filled and skip to next oracle sequence
		nextDeliverySeqOnInscription, err := p.bscExecutor.GetNextDeliveryOracleSequence()
		if err != nil {
			return err
		}
		if seq < nextDeliverySeqOnInscription {
			err = p.daoManager.BSCDao.UpdateBatchPackagesStatus(pkgIds, db.Filled)
			if err != nil {
				relayercommon.Logger.Errorf("failed to update packages error %s", err.Error())
				return err
			}
			relayercommon.Logger.Infof("packages' oracle sequence %d is less than nex delivery oracle sequence %d", seq, nextDeliverySeqOnInscription)
			continue
		}

		encodedPayload, err := rlp.EncodeToBytes(aggPkgs)
		blsClaim := oracletypes.BlsClaim{
			// chain ids are validated when packages persisted into DB, non-matched ones would be omitted
			SrcChainId:  uint32(p.config.BSCConfig.ChainId),
			DestChainId: uint32(p.config.InscriptionConfig.ChainId),
			Timestamp:   uint64(pkgsForSeq[0].TxTime),
			Sequence:    seq,
			Payload:     encodedPayload,
		}
		eventHash := blsClaim.GetSignBytes()
		if err != nil {
			return fmt.Errorf("encode packages error, err=%s", err.Error())
		}
		channelId := relayercommon.OracleChannelId
		v := p.constructVoteAndSign(eventHash[:])

		// broadcast v
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

func (p *BSCVoteProcessor) CollectVotes() {
	for {
		err := p.collectVotes()
		if err != nil {
			time.Sleep(RetryInterval)
		}
	}
}

func (p *BSCVoteProcessor) collectVotes() error {
	pkgs, err := p.daoManager.BSCDao.GetPackagesByStatus(db.SelfVoted)
	if err != nil {
		relayercommon.Logger.Errorf("failed to get voted packages from db, error: %s", err.Error())
		return err
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

		err = p.daoManager.BSCDao.UpdateBatchPackagesStatus(txIds, db.AllVoted)
		if err != nil {
			return err
		}
	}
	return nil
}

// prepareEnoughValidVotesForPackages will prepare fetch and validate votes result, store in votes
func (p *BSCVoteProcessor) prepareEnoughValidVotesForPackages(channelId relayercommon.ChannelId, sequence uint64) error {
	localVote, err := p.daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(uint8(channelId), sequence, hex.EncodeToString(p.blsPublicKey))
	if err != nil {
		return err
	}
	validators, err := p.bscExecutor.InscriptionExecutor.QueryCachedLatestValidators()
	if err != nil {
		return err
	}
	// Query from votePool until there are more than 2/3 votes
	err = p.queryMoreThanTwoThirdValidVotes(localVote, validators)
	if err != nil {
		return err
	}
	return nil
}

// queryMoreThanTwoThirdValidVotes queries votes from votePool
func (p *BSCVoteProcessor) queryMoreThanTwoThirdValidVotes(localVote *model.Vote, validators []stakingtypes.Validator) error {
	triedTimes := 0
	validVotesTotalCnt := 1 // assume local vote is valid
	channelId := localVote.ChannelId
	seq := localVote.Sequence
	ticker := time.NewTicker(2 * time.Second)
	for {
		<-ticker.C
		triedTimes++
		if triedTimes > QueryVotepoolMaxRetryTimes {
			return nil
		}
		queriedVotes, err := p.votePoolExecutor.QueryVotes(localVote.EventHash, votepool.FromBscCrossChainEvent)
		if err != nil {
			relayercommon.Logger.Errorf("encounter error when query votes. will retry.")
			return err
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

			if err := VerifySignature(v, localVote.EventHash[:]); err != nil {
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
			err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, channelId, seq, localVote.ClaimPayload))
			if err != nil {
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
			err = p.votePoolExecutor.BroadcastVote(v)
			if err != nil {
				return err
			}
		}
		continue
	}
}

func (p *BSCVoteProcessor) constructVoteAndSign(eventHash []byte) *votepool.Vote {
	var v votepool.Vote
	v.EventType = votepool.FromBscCrossChainEvent
	v.EventHash = eventHash
	p.signer.SignVote(&v)
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
