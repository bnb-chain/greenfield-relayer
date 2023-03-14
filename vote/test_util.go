package vote

import (
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	"github.com/tendermint/tendermint/votepool"

	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/util"
)

const (
	LocalRelayerHexBlsPrivKey = "your_private_key"
)

// PrivKeys Bls Private keys from other mock relayers
var PrivKeys = []string{
	"your_private_key",
	"your_private_key",
}

// BroadcastVotesFromOtherRelayers for mimic multi relayers when bsc -> gnfd
func BroadcastVotesFromOtherRelayers(
	privateKeysList []string,
	daoManager *dao.DaoManager,
	gnfdExecutor *executor.GreenfieldExecutor, channelId uint8, seq uint64, eventType uint8) {
	var vote *model.Vote
	// retry to query vote from local DB
	for {
		localVote, err := daoManager.VoteDao.GetVoteByChannelIdAndSequenceAndPubKey(channelId,
			seq, hex.EncodeToString(util.BlsPubKeyFromPrivKeyStr(LocalRelayerHexBlsPrivKey)))
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		vote = localVote
		break
	}

	for _, pk := range privateKeysList {
		secretKey, err := blst.SecretKeyFromBytes(common.Hex2Bytes(pk))
		if err != nil {
			panic(err)
		}
		pubKey := secretKey.PublicKey()
		sign := secretKey.Sign(vote.EventHash[:]).Marshal()

		mockVoteFromRelayer := &votepool.Vote{
			PubKey:    pubKey.Marshal(),
			Signature: sign,
			EventType: votepool.EventType(eventType),
			EventHash: vote.EventHash[:],
		}
		err = gnfdExecutor.BroadcastVote(mockVoteFromRelayer)
		if err != nil {
			panic(err)
		}
	}
}
