package vote

import (
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	blscmn "github.com/prysmaticlabs/prysm/crypto/bls/common"
	"github.com/tendermint/tendermint/votepool"
)

type VoteSigner struct {
	privKey blscmn.SecretKey
	pubKey  blscmn.PublicKey
}

func NewVoteSigner(privkey []byte) (*VoteSigner, error) {
	privKey, err := blst.SecretKeyFromBytes(privkey)
	if err != nil {
		return nil, err
	}
	pubKey := privKey.PublicKey()
	return &VoteSigner{
		privKey: privKey,
		pubKey:  pubKey,
	}, nil
}

// SignVote sign a vote, data is used to signed to generate the signature
func (signer *VoteSigner) SignVote(vote *votepool.Vote, data []byte) {
	signature := signer.privKey.Sign(data[:])
	vote.EventHash = append(vote.EventHash, data[:]...)
	vote.PubKey = append(vote.PubKey, signer.pubKey.Marshal()...)
	vote.Signature = append(vote.Signature, signature.Marshal()...)
}
