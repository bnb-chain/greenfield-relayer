package executor

import (
	"time"

	ethereumcommon "github.com/ethereum/go-ethereum/common"
)

const (
	prefixLength                   = 1
	destChainIDLength              = 2
	channelIDLength                = 1
	DefaultGasPrice                = 20000000000 // 20 GWei
	FallBehindThreshold            = 5
	SleepSecondForUpdateClient     = 10
	DataSeedDenyServiceThreshold   = 60
	SequenceStoreName              = "crosschain"
	RPCTimeout                     = 3 * time.Second
	RelayerBytesLength             = 48
	UpdateCachedValidatorsInterval = 1 * time.Minute
)

var (
	PrefixForReceiveSequenceKey = []byte{0xf1}

	InscriptionLightClientContractAddr = ethereumcommon.HexToAddress("0x71b750F84B4d1d72C17EcEba811fA6E4C8c9CfdC")
	CrossChainContractAddr             = ethereumcommon.HexToAddress("0x2078fEEF78BD06AcDdb46619a681327aaEeeAE20")
)
