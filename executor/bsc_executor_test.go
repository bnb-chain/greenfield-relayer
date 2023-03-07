package executor

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func InitBSCExecutor() *BSCExecutor {
	cfg := InitTestConfig()
	return NewBSCExecutor(cfg)
}

func TestGetBlockHeight(t *testing.T) {
	height, err := InitBSCExecutor().GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetNextReceiveSequence(t *testing.T) {
	seq, err := InitBSCExecutor().GetNextReceiveSequenceForChannel(3)
	require.NoError(t, err)
	t.Log(seq)
}

func TestGetBlockHeader(t *testing.T) {
	e := InitBSCExecutor()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	header, err := e.GetBlockHeaderAtHeight(height)
	require.NoError(t, err)
	json, err := header.MarshalJSON()
	require.NoError(t, err)
	t.Log(string(json))
}

func TestGetLightClientHeight(t *testing.T) {
	height, err := InitBSCExecutor().GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(height)
}

func TestQueryLatestValidators(t *testing.T) {
	relayers, err := InitBSCExecutor().QueryLatestValidators()
	require.NoError(t, err)
	for _, r := range relayers {
		t.Log(r.RelayerAddress.String())
		t.Log(hex.EncodeToString(r.BlsPublicKey))
	}
	t.Log(relayers)
}

func TestSyncTendermintHeader(t *testing.T) {
	e, _ := InitExecutors()
	curLightClientHeight, err := e.GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(curLightClientHeight)
	hash, err := e.SyncTendermintLightBlock(10)
	require.NoError(t, err)
	time.Sleep(10 * time.Second)
	t.Log(hash.String())
	nextHeight, err := e.GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(nextHeight)
}

func TestGetLogsFromHeader(t *testing.T) {
	e := InitBSCExecutor()
	client := e.GetRpcClient()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	header, err := e.GetBlockHeaderAtHeight(height)
	require.NoError(t, err)
	topics := [][]ethereumcommon.Hash{{ethereumcommon.HexToHash("0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a")}}
	blockHash := header.Hash()
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethereumcommon.Address{ethereumcommon.HexToAddress(e.config.RelayConfig.CrossChainContractAddr)},
	})
	require.NoError(t, err)
	t.Log(logs)
}
