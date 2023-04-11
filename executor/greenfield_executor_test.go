package executor

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func InitGnfdExecutor() *GreenfieldExecutor {
	cfg := InitTestConfig()
	return NewGreenfieldExecutor(cfg)
}

func TestGetLatestBlockHeightWithRetry(t *testing.T) {
	e := InitGnfdExecutor()
	height, err := e.GetLatestBlockHeight()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextReceiveOracleSequence(t *testing.T) {
	e := InitGnfdExecutor()
	oracleSeq, err := e.GetNextReceiveOracleSequence()
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetNextSendSequenceForChannel(t *testing.T) {
	e := InitGnfdExecutor()
	sendSeq, err := e.GetNextSendSequenceForChannelWithRetry(1)
	require.NoError(t, err)
	t.Log(sendSeq)
}

func TestGetInturnRelayer(t *testing.T) {
	e := InitGnfdExecutor()
	relayer, err := e.GetInturnRelayer()
	require.NoError(t, err)
	t.Log(relayer)
}

func TestGetConsensusStatus(t *testing.T) {
	e := InitGnfdExecutor()
	rpcClient := e.getRpcClient()
	h := int64(1)
	validators, err := rpcClient.Validators(context.Background(), &h, nil, nil)
	assert.NoError(t, err)

	b, _, err := e.GetBlockAndBlockResultAtHeight(1)
	assert.NoError(t, err)

	t.Log("nexValidator hash: ", hex.EncodeToString(b.NextValidatorsHash))
	for i, validator := range validators.Validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hex.EncodeToString(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hex.EncodeToString(validator.BlsKey))
	}
}

func TestGetLatestValidators(t *testing.T) {
	e := InitGnfdExecutor()
	rpcClient := e.getRpcClient()
	validators, err := rpcClient.Validators(context.Background(), nil, nil, nil)
	assert.NoError(t, err)
	for i, validator := range validators.Validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hex.EncodeToString(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hex.EncodeToString(validator.BlsKey))
	}
}
