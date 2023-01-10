package util

import (
	"encoding/binary"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/crypto/bls/blst"
	"github.com/willf/bitset"
)

func IndexOf(element string, data []string) int {
	for i, v := range data {
		if element == v {
			return i
		}
	}
	return -1
}

func GetBlsPubKeyFromPrivKeyStr(privKeyStr string) []byte {
	privKey, err := blst.SecretKeyFromBytes(common.Hex2Bytes(privKeyStr))
	if err != nil {
		panic(err)
	}
	return privKey.PublicKey().Marshal()
}

// QuotedStrToIntWithBitSize convert a QuoteStr ""6""  to int 6
func QuotedStrToIntWithBitSize(str string, bitSize int) (int64, error) {
	s, err := strconv.Unquote(str)
	if err != nil {
		return 0, err
	}
	num, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func BitSetToBigInt(set *bitset.BitSet) *big.Int {
	bts := make([]byte, 0)
	for _, i := range set.Bytes() {
		bt := make([]byte, 8)
		binary.LittleEndian.PutUint64(bt, i)
		bts = append(bts, bt...)
	}
	res := new(big.Int)
	res.SetBytes(bts)
	return res
}
