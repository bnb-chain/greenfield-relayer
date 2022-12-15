package vote

import (
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/crypto/bls"
	"inscription-relayer/db/model"
)

type VoteData struct {
	EventHash []byte
	Sequence  uint64
	ChannelId uint8
}

func (d *VoteData) ToDbModel() *model.VoteData {
	return &model.VoteData{
		EventHash: d.EventHash,
		Sequence:  d.Sequence,
		ChannelId: d.ChannelId,
	}
}

type Vote struct {
	PubKey    BLSPublicKey
	Signature BLSSignature
	EvenType  EventType
	EventHash []byte
}

func (vote *Vote) ToDbModel(channelId uint8, sequence uint64) *model.Vote {
	return &model.Vote{
		PubKey:    string(vote.PubKey[:]),
		Signature: string(vote.Signature[:]),
		EvenType:  uint32(vote.EvenType),
		EventHash: vote.EventHash,
		ChannelId: channelId,
		Sequence:  sequence,
	}
}

// Verify vote
func (vote *Vote) Verify(eventHash []byte) error {
	blsPubKey, err := bls.PublicKeyFromBytes(vote.PubKey[:])
	if err != nil {
		return errors.Wrap(err, "convert public key from bytes to bls failed")
	}
	sig, err := bls.SignatureFromBytes(vote.Signature[:])
	if err != nil {
		return errors.Wrap(err, "invalid signature")
	}
	if !sig.Verify(blsPubKey, eventHash[:]) {
		return errors.New("verify bls signature failed.")
	}
	return nil
}

func AggregatedSignature(votes []*model.Vote) ([]byte, error) {
	// Prepare aggregated vote signature
	voteAddrSet := make(map[string]struct{}, len(votes))
	signatures := make([][]byte, 0, len(votes))
	for _, v := range votes {
		voteAddrSet[v.PubKey] = struct{}{}
		signatures = append(signatures, []byte(v.Signature)[:])
	}
	sigs, err := bls.MultipleSignaturesFromBytes(signatures)
	if err != nil {
		return nil, err
	}
	return bls.AggregateSignatures(sigs).Marshal(), nil
}
