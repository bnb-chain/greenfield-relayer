package executor

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
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

func TestGetConsensusStatus(t *testing.T) {
	e := InitGnfdExecutor()
	validators, err := e.getRpcClient().Validators(context.Background(), nil, nil, nil)
	assert.NoError(t, err)

	b, err := e.GetBlockAtHeight(1)
	t.Log("nexValidator hash: ", hex.EncodeToString(b.NextValidatorsHash))
	for i, validator := range validators.Validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hex.EncodeToString(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hex.EncodeToString(validator.RelayerBlsKey))
	}
}
