package executor

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestGetBlockHeight(t *testing.T) {
	e, _ := InitExecutors()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	t.Log(height)
}

func TestGetOracleSequence(t *testing.T) {
	e, _ := InitExecutors()
	seq, err := e.GetNextDeliveryOracleSequence()
	require.NoError(t, err)
	t.Log(seq)
}

func TestGetBlockHeader(t *testing.T) {
	e, _ := InitExecutors()
	height, err := e.GetLatestBlockHeightWithRetry()
	require.NoError(t, err)
	header, err := e.GetBlockHeaderAtHeight(height)
	require.NoError(t, err)
	json, err := header.MarshalJSON()
	require.NoError(t, err)
	t.Log(string(json))
}

func TestGetSequence(t *testing.T) {
	e, _ := InitExecutors()
	seq, err := e.GetNextReceiveSequenceForChannel(1)
	require.NoError(t, err)
	t.Log(seq)
}

func TestGetLightClientHeight(t *testing.T) {
	e, _ := InitExecutors()
	height, err := e.GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(height)
}

func TestQueryLatestValidators(t *testing.T) {
	e, _ := InitExecutors()
	relayers, err := e.QueryLatestValidators()
	require.NoError(t, err)
	t.Log(relayers)
}

func TestSyncTendermintHeader(t *testing.T) {
	e, _ := InitExecutors()
	curLightClientHeight, err := e.GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(curLightClientHeight)
	hash, err := e.SyncTendermintLightClientHeader(curLightClientHeight + 1)
	time.Sleep(10 * time.Second)
	require.NoError(t, err)
	t.Log(hash.String())
	nextHeight, err := e.GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(nextHeight)
	require.EqualValues(t, curLightClientHeight+1, nextHeight)
}

func TestGetLogsFromHeader(t *testing.T) {
	e, _ := InitExecutors()
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
		Addresses: []ethereumcommon.Address{CrossChainContractAddr},
	})
	require.NoError(t, err)
	t.Log(logs)
}
