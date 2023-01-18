package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInsBlockHeight(t *testing.T) {
	_, e := InitExecutors()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextOracleSequence(t *testing.T) {
	_, e := InitExecutors()
	oracleSeq, err := e.GetNextOracleSequence()
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetValidators(t *testing.T) {
	_, e := InitExecutors()
	validators, err := e.queryLatestValidators()
	require.NoError(t, err)
	t.Log(validators)
}
