package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
)

func IndexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1
}

func GetBlsPubKeyFromPrivKeyStr(privKeyStr string) ([]byte, error) {
	privKey, err := blst.SecretKeyFromBytes(common.Hex2Bytes(privKeyStr))
	if err != nil {
		return nil, err
	}
	return privKey.PublicKey().Marshal(), nil
}
