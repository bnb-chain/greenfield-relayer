package vote

import (
	"encoding/hex"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/crypto/bls"
	"github.com/tendermint/tendermint/votepool"
	"inscription-relayer/db/model"
)

// Verify vote
func Verify(vote *votepool.Vote, eventHash []byte) error {
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

func AggregatedSignatureAndValidatorBitSet(votes []*model.Vote, validators []stakingtypes.Validator) ([]byte, uint64, error) {
	signatures := make([][]byte, 0, len(votes))
	voteAddrSet := make(map[string]struct{}, len(votes))
	var votedAddressSet uint64
	for _, v := range votes {
		voteAddrSet[v.PubKey] = struct{}{}
		signatures = append(signatures, common.Hex2Bytes(v.Signature))
	}

	for idx, valInfo := range validators {
		if _, ok := voteAddrSet[hex.EncodeToString(valInfo.RelayerBlsKey)]; ok {
			votedAddressSet |= 1 << idx
		}
	}

	sigs, err := bls.MultipleSignaturesFromBytes(signatures)
	if err != nil {
		return nil, 0, err
	}
	return bls.AggregateSignatures(sigs).Marshal(), votedAddressSet, nil
}
