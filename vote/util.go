package vote

import (
	"encoding/hex"
	"reflect"

	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/crypto/bls"
	"github.com/tendermint/tendermint/votepool"
	"github.com/willf/bitset"
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

func AggregateSignatureAndValidatorBitSet(votes []*model.Vote, validators interface{}) ([]byte, *bitset.BitSet, error) {
	signatures := make([][]byte, 0, len(votes))
	voteAddrSet := make(map[string]struct{}, len(votes))
	valBitSet := bitset.New(256)
	for _, v := range votes {
		voteAddrSet[v.PubKey] = struct{}{}
		signatures = append(signatures, common.Hex2Bytes(v.Signature))
	}
	if reflect.TypeOf(validators).Elem() == reflect.TypeOf(executor.Validator{}) {
		for idx, valInfo := range validators.([]executor.Validator) {
			if _, ok := voteAddrSet[hex.EncodeToString(valInfo.BlsPublicKey[:])]; ok {
				valBitSet.Set(uint(idx))
			}
		}
	} else {
		for idx, valInfo := range validators.([]stakingtypes.Validator) {
			if _, ok := voteAddrSet[hex.EncodeToString(valInfo.RelayerBlsKey[:])]; ok {
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
