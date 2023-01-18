package executor

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
	seq, err := e.GetNextSequence(1)
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
	t.Log(relayers[0].RelayerAddress)
	t.Log(hex.EncodeToString(relayers[0].BlsPublicKey))

}

func TestSyncTendermintHeader(t *testing.T) {
	e, _ := InitExecutors()
	cliHeight, err := e.GetLightClientLatestHeight()
	require.NoError(t, err)
	t.Log(cliHeight)
	//insLaestheight, err := e.GetLatestBlockHeightWithRetry()
	hash, err := e.SyncTendermintLightClientHeader(cliHeight + 1)
	time.Sleep(10 * time.Second)

	require.NoError(t, err)
	t.Log(hash.String())
	nextHeight, err := e.GetLightClientLatestHeight()
	t.Log(nextHeight)
	//require.EqualValues(t, cliHeight+1, nextHeight)
}

func TestGetLogsFromHeader(t *testing.T) {
	e, _ := InitExecutors()
	client := e.GetRpcClient()
	height, err := e.GetLatestBlockHeightWithRetry()
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
