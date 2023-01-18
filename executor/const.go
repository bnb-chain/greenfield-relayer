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

	InscriptionLightClientContractAddr = ethereumcommon.HexToAddress("0xeDC2E0dDbB7F10C3CA10c2c964C134b758044Bce")
	CrossChainContractAddr             = ethereumcommon.HexToAddress("0x58d1A4Cb3622B3201Fe8E21B65824fa624a0026C")
)
