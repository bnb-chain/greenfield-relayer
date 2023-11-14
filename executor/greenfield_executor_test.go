package executor

import (
	"context"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"testing"

	cbfttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func GnfdExecutor() *GreenfieldExecutor {
	return NewGreenfieldExecutor(GetTestConfig())
}

func TestGetLatestBlockHeightWithRetry(t *testing.T) {
	e := GnfdExecutor()
	height, err := e.GetLatestBlockHeight()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextReceiveOracleSequence(t *testing.T) {
	e := GnfdExecutor()
	oracleSeq, err := e.GetNextReceiveOracleSequence(sdk.ChainID(e.config.BSCConfig.ChainId))
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetNextSendSequenceForChannel(t *testing.T) {
	e := GnfdExecutor()
	sendSeq, err := e.GetNextSendSequenceForChannelWithRetry(sdk.ChainID(e.config.BSCConfig.ChainId), 1)
	require.NoError(t, err)
	t.Log(sendSeq)
}

func TestGetInturnRelayer(t *testing.T) {
	e := GnfdExecutor()
	relayer, err := e.GetInturnRelayer(oracletypes.CLAIM_SRC_CHAIN_BSC)
	require.NoError(t, err)
	t.Log(relayer)

	relayer, err = e.GetInturnRelayer(oracletypes.CLAIM_SRC_CHAIN_OP_BNB)
	require.NoError(t, err)
	t.Log(relayer)
}

func TestGetLatestValidators(t *testing.T) {
	e := GnfdExecutor()
	_, validators, err := e.GetGnfdClient().GetValidatorSet(context.Background())
	assert.NoError(t, err)
	for i, validator := range validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hex.EncodeToString(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hex.EncodeToString(validator.BlsKey))
	}
}

func TestGetConsensusStatus(t *testing.T) {
	e := GnfdExecutor()
	height := int64(1)
	validators, err := e.GetGnfdClient().GetValidatorsByHeight(context.Background(), height)
	assert.NoError(t, err)
	b, _, err := e.GetBlockAndBlockResultAtHeight(height)
	assert.NoError(t, err)
	t.Logf("NexValidator Hash: %s", hexutil.Encode(b.NextValidatorsHash))
	for i, validator := range validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hexutil.Encode(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hexutil.Encode(validator.BlsKey))
	}
	cs, err := getCysString(e, height)
	assert.NoError(t, err)
	t.Logf("consensus: %s", cs)
}

func getCysString(e *GreenfieldExecutor, height int64) (string, error) {
	validators, err := e.GetGnfdClient().GetValidatorsByHeight(context.Background(), height)
	if err != nil {
		return "", err
	}
	block, err := e.GetGnfdClient().GetBlockByHeight(context.Background(), height)
	if err != nil {
		return "", err
	}
	cs := ConsensusState{
		ChainID:              block.ChainID,
		Height:               uint64(block.Height),
		NextValidatorSetHash: block.NextValidatorsHash,
		ValidatorSet: &cbfttypes.ValidatorSet{
			Validators: validators,
		},
	}
	csBytes, err := cs.encodeConsensusState()
	if err != nil {
		return "", err
	}
	return hexutil.Encode(csBytes), nil
}
