package vote

import (
	"encoding/hex"
	"reflect"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cometbft/cometbft/votepool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
	"github.com/willf/bitset"

	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/types"
)

// VerifySignature verifies vote signature
func VerifySignature(vote *votepool.Vote, eventHash []byte) error {
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

// AggregateSignatureAndValidatorBitSet aggregates signature from multiple votes, and marks the bitset of validators who contribute votes
func AggregateSignatureAndValidatorBitSet(votes []*model.Vote, validators interface{}) ([]byte, *bitset.BitSet, error) {
	signatures := make([][]byte, 0, len(votes))
	voteAddrSet := make(map[string]struct{}, len(votes))
	valBitSet := bitset.New(ValidatorsCapacity)
	for _, v := range votes {
		voteAddrSet[v.PubKey] = struct{}{}
		signatures = append(signatures, common.Hex2Bytes(v.Signature))
	}
	if reflect.TypeOf(validators).Elem() == reflect.TypeOf(types.Validator{}) {
		for idx, valInfo := range validators.([]types.Validator) {
			if _, ok := voteAddrSet[hex.EncodeToString(valInfo.BlsPublicKey[:])]; ok {
				valBitSet.Set(uint(idx))
			}
		}
	} else {
		for idx, valInfo := range validators.([]*tmtypes.Validator) {
			if _, ok := voteAddrSet[hex.EncodeToString(valInfo.BlsKey[:])]; ok {
				valBitSet.Set(uint(idx))
			}
		}
	}
	sigs, err := bls.MultipleSignaturesFromBytes(signatures)
	if err != nil {
		return nil, valBitSet, err
	}
	return bls.AggregateSignatures(sigs).Marshal(), valBitSet, nil
}
