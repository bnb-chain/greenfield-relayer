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
	LocalRelayerHexBlsPrivKey = "275ebf409f5caa121bdc37841660b4548bfbf81eb8593442640ffe0b66dfc86f"
)

// PrivKeys Bls Private keys from other mock relayers
var PrivKeys = []string{
	"60e5839445580b001576ce8fb0b08cf2b37f8289faaf49a2a3d1e36dbbe588a1",
	"4b4d06d9c4af19c175962190596ed7e01e1b818821ef4cbf593f6ec84345a0f0",
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
