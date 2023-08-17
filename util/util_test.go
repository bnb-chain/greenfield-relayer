package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willf/bitset"
)

func TestQuotedStrToInt(t *testing.T) {
	num, err := QuotedStrToIntWithBitSize("\"666666\"", 64)
	require.NoError(t, err)
	require.Equal(t, uint64(666666), num)
}

func TestBitSetToBigInt(t *testing.T) {
	valBitSet := bitset.New(256)
	valBitSet.Set(0)
	valBitSet.Set(1)
	valBitSet.Set(2)
	valBitSet.Set(9)
	valBitSet.Set(64)
	valBitSet.Set(128)
	valBitSet.Set(255)
	bigint := BitSetToBigInt(valBitSet)

	require.EqualValues(t, 1, bigint.Bit(0))
	require.EqualValues(t, 1, bigint.Bit(1))
	require.EqualValues(t, 1, bigint.Bit(2))
	require.EqualValues(t, 1, bigint.Bit(9))
	require.EqualValues(t, 1, bigint.Bit(64))
	require.EqualValues(t, 1, bigint.Bit(128))
	require.EqualValues(t, 1, bigint.Bit(255))
	require.NotEqual(t, 1, bigint.Bit(3))
}
