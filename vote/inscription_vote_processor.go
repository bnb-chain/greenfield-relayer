package vote

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/avast/retry-go/v4"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	"github.com/tendermint/tendermint/votepool"
	"gorm.io/gorm"
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
	votePoolExecutor *VotePoolExecutor) *InscriptionVoteProcessor {
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
	txs, err := p.daoManager.InscriptionDao.GetTransactionsByStatusAndHeight(db.SAVED, leastSavedTxHeight)
	if err != nil {
		relayercommon.Logger.Errorf("failed to get transactions at height %d from db, error: %s", leastSavedTxHeight, err.Error())
		return err
	}

	if len(txs) == 0 {
		time.Sleep(RetryInterval)
		return nil
	}

	// for every tx, we are going to sign it and broadcast vote of it.
	for _, tx := range txs {
		v, err := p.constructVoteAndSign(tx)
		if err != nil {
			return err
		}

		//TODO remove testing purpose code
		bs2 := common.Hex2Bytes("16f6742aee55411cc79c06af0e265b9df5ba3b54de85fbfc96c7d6a67469e4d0")
		secretKey2, err := blst.SecretKeyFromBytes(bs2)
		if err != nil {
			panic(err)
		}
		eh, _ := p.getEventHash(tx)
		pubKey2 := secretKey2.PublicKey()
		sign2 := secretKey2.Sign(eh).Marshal()

		mockVoteFromRelayer2 := &votepool.Vote{
			PubKey:    pubKey2.Marshal(),
			Signature: sign2,
			EventType: 1,
			EventHash: eh,
		}

		//broadcast v
		if err = retry.Do(func() error {
			err = p.votePoolExecutor.BroadcastVote(mockVoteFromRelayer2)
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
			err = p.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, db.VOTED)
			if err != nil {
				return err
			}
			err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, tx.ChannelId, tx.Sequence, common.Hex2Bytes(tx.PayLoad)))
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
	txs, err := p.daoManager.InscriptionDao.GetTransactionsByStatus(db.VOTED)
	if err != nil {
		relayercommon.Logger.Errorf("failed to get voted transactions from db, error: %s", err.Error())
		return err
	}
	for _, tx := range txs {
		err := p.prepareEnoughValidVotesForTx(tx)
		if err != nil {
			return err
		}

		err = p.daoManager.InscriptionDao.UpdateTransactionStatus(tx.Id, db.VOTED_All)
		if err != nil {
			return err
		}
	}
	return nil
}

// prepareEnoughValidVotesForTx will prepare fetch and validate votes result, store in votes
func (p *InscriptionVoteProcessor) prepareEnoughValidVotesForTx(tx *model.InscriptionRelayTransaction) error {

	validators, err := p.inscriptionExecutor.QueryLatestValidators()
	if err != nil {
		return err
	}
	//Query from votePool until there are more than 2/3 valid votes
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

			if err := Verify(v, localVote.EventHash); err != nil {
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
			err = p.daoManager.VoteDao.SaveVote(EntityToDto(v, channelId, seq, common.Hex2Bytes(tx.PayLoad)))
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
