package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLatestBlockHeightWithRetry(t *testing.T) {
	_, e := InitExecutors()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextReceiveOracleSequence(t *testing.T) {
	_, e := InitExecutors()
	oracleSeq, err := e.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetNextReceiveSequenceForChannel(t *testing.T) {
	_, e := InitExecutors()
	oracleSeq, err := e.GetNextReceiveSequenceForChannel(2)
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetValidators(t *testing.T) {
	_, e := InitExecutors()
	validators, err := e.queryLatestValidators()
	require.NoError(t, err)
	t.Log(validators)
}
