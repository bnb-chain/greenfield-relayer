package executor

import (
	"context"
	"encoding/hex"
	sdkclient "github.com/bnb-chain/greenfield-go-sdk/client/chain"
	"github.com/bnb-chain/greenfield-relayer/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

func TestGetNextReceiveSequenceForChannel(t *testing.T) {
	e := InitGnfdExecutor()
	oracleSeq, err := e.GetNextReceiveSequenceForChannel(2)
	require.NoError(t, err)
	t.Log(oracleSeq)
}

func TestGetConsensusStatus(t *testing.T) {
	e := InitGnfdExecutor()
	rpcClient, err := e.getRpcClient()
	assert.NoError(t, err)

	validators, err := rpcClient.Validators(context.Background(), nil, nil, nil)
	assert.NoError(t, err)

	b, _, err := e.GetBlockAndBlockResultAtHeight(1)
	assert.NoError(t, err)

	t.Log("nexValidator hash: ", hex.EncodeToString(b.NextValidatorsHash))
	for i, validator := range validators.Validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hex.EncodeToString(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hex.EncodeToString(validator.RelayerBlsKey))
	}
}

func TestGetLatestValidators(t *testing.T) {
	e := InitGnfdExecutor()
	rpcClient, err := e.getRpcClient()
	assert.NoError(t, err)
	validators, err := rpcClient.Validators(context.Background(), nil, nil, nil)
	assert.NoError(t, err)
	for i, validator := range validators.Validators {
		t.Logf("validator %d", i)
		t.Logf("validator pubkey %s", hexutil.Encode(validator.PubKey.Bytes()))
		t.Logf("validator votingpower %d", validator.VotingPower)
		t.Logf("relayeraddress %s", hex.EncodeToString(validator.RelayerAddress))
		t.Logf("relayer bls pub key %s", hex.EncodeToString(validator.RelayerBlsKey))
	}
}

func TestAds(t *testing.T) {
	cfg := config.ParseConfigFromFile("../config/local/config_local_0.json")
	clients := sdkclient.NewGnfdCompositClients(
		cfg.GreenfieldConfig.RPCAddrs,
		cfg.GreenfieldConfig.GRPCAddrs,
		cfg.GreenfieldConfig.ChainIdString,
		sdkclient.WithGrpcDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	clients.GetClient()
}
