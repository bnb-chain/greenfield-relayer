package vote

import (
	"encoding/hex"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	"github.com/tendermint/tendermint/votepool"
	"time"
)

const (
	LocalRelayerHexBlsPrivKey = "4f1f561f5835d0f310d3cf072821155834653ab5f417ff0905eaab3023d56b1e"
)

// PrivKeys Bls Private keys from other mock relayers
var PrivKeys = []string{
	"3a1055a667eddef7405a554f2994aedea43c8258712a013b3a61532e8cd0f032",
	"139ace9a52fa78b9f4dc2f151231225f9503d60f3aefeef89a1ee82d6d48ef9a",
}

// BroadcastVotesFromOtherRelayers for mimic multi relayers when bsc -> gnfd
func BroadcastVotesFromOtherRelayers(
	privateKeysList []string,
	daoManager *dao.DaoManager,
	votePoolExecutor *VotePoolExecutor, channelId uint8, seq uint64, eventType uint8) {
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
		err = votePoolExecutor.BroadcastVote(mockVoteFromRelayer)
		if err != nil {
			panic(err)
		}
	}
}
