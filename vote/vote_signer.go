package vote

import (
	"github.com/cometbft/cometbft/votepool"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	blscmn "github.com/prysmaticlabs/prysm/crypto/bls/common"
)

type VoteSigner struct {
	privKey blscmn.SecretKey
	pubKey  blscmn.PublicKey
}

func NewVoteSigner(pk []byte) *VoteSigner {
	privKey, err := blst.SecretKeyFromBytes(pk)
	if err != nil {
		panic(err)
	}
	pubKey := privKey.PublicKey()
	return &VoteSigner{
		privKey: privKey,
		pubKey:  pubKey,
	}
}

// SignVote signs a vote by relayer's private key
func (signer *VoteSigner) SignVote(vote *votepool.Vote) {
	vote.PubKey = append(vote.PubKey, signer.pubKey.Marshal()...)
	signature := signer.privKey.Sign(vote.EventHash[:])
	vote.Signature = append(vote.Signature, signature.Marshal()...)
}
