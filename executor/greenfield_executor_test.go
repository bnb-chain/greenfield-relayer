package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func InitInsExecutor() *GreenfieldExecutor {
	cfg := InitTestConfig()
	return NewGreenfieldExecutor(cfg)
}

func TestGetLatestBlockHeightWithRetry(t *testing.T) {
	e := InitInsExecutor()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextReceiveOracleSequence(t *testing.T) {
	e := InitInsExecutor()
	oracleSeq, err := e.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetNextReceiveSequenceForChannel(t *testing.T) {
	e := InitInsExecutor()
	oracleSeq, err := e.GetNextReceiveSequenceForChannel(2)
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetValidators(t *testing.T) {
	e := InitInsExecutor()
	validators, err := e.queryLatestValidators()
	require.NoError(t, err)
	t.Log(validators)
}
