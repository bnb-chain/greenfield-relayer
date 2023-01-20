package vote

import (
	"encoding/hex"
	"time"

	"github.com/tendermint/tendermint/votepool"

	"github.com/bnb-chain/inscription-relayer/db/model"
)

func DtoToEntity(v *model.Vote) (*votepool.Vote, error) {
	pubKeyBts, err := hex.DecodeString(v.PubKey)
	if err != nil {
		return nil, err
	}
	sigBts, err := hex.DecodeString(v.Signature)
	if err != nil {
		return nil, err
	}
	res := votepool.Vote{}
	res.EventType = votepool.EventType(v.EventType)
	res.PubKey = append(res.PubKey, pubKeyBts...)
	res.Signature = append(res.Signature, sigBts...)
	res.EventHash = append(res.EventHash, v.EventHash...)
	return &res, nil
}

func EntityToDto(from *votepool.Vote, channelId uint8, sequence uint64, payload []byte) *model.Vote {
	v := model.Vote{
		PubKey:       hex.EncodeToString(from.PubKey[:]),
		Signature:    hex.EncodeToString(from.Signature[:]),
		EventType:    uint32(from.EventType),
		ClaimPayload: payload,
		EventHash:    from.EventHash,
		ChannelId:    channelId,
		Sequence:     sequence,
		CreatedTime:  time.Now().Unix(),
	}
	return &v
}
