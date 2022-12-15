package executor

import (
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"inscription-relayer/common"
)

const (
	prefixLength      = 1
	destChainIDLength = 2
	channelIDLength   = 1

	DefaultGasPrice = 20000000000 // 20 GWei

	FallBehindThreshold          = 5
	SleepSecondForUpdateClient   = 10
	DataSeedDenyServiceThreshold = 60

	SequenceStoreName = "sc"
)

var (
	prefixForCrossChainPackageKey = []byte{0x00}
	prefixForSequenceKey          = []byte{0xf0}

	PureHeaderSyncChannelID common.ChannelId = 1

	tendermintLightClientContractAddr = ethereumcommon.HexToAddress("0x0000000000000000000000000000000000001003")
	relayerIncentivizeContractAddr    = ethereumcommon.HexToAddress("0x0000000000000000000000000000000000001005")
	relayerHubContractAddr            = ethereumcommon.HexToAddress("0x0000000000000000000000000000000000001006")
	crossChainContractAddr            = ethereumcommon.HexToAddress("0x0000000000000000000000000000000000002000")
)
