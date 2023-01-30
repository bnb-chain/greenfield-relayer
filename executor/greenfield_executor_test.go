package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func InitGnfdExecutor() *GreenfieldExecutor {
	cfg := InitTestConfig()
	return NewGreenfieldExecutor(cfg)
}

func TestGetLatestBlockHeightWithRetry(t *testing.T) {
	e := InitGnfdExecutor()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextReceiveOracleSequence(t *testing.T) {
	e := InitGnfdExecutor()
	oracleSeq, err := e.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetNextReceiveSequenceForChannel(t *testing.T) {
	e := InitGnfdExecutor()
	oracleSeq, err := e.GetNextReceiveSequenceForChannel(2)
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetValidators(t *testing.T) {
	e := InitGnfdExecutor()
	validators, err := e.queryLatestValidators()
	require.NoError(t, err)
	t.Log(validators)
}
