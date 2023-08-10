package executor

import (
	"math/big"
	"time"
)

const (
	DefaultGasPrice                = 20000000000 // 20 GWei
	FallBehindThreshold            = 5
	SleepSecondForUpdateClient     = 10
	DataSeedDenyServiceThreshold   = 60
	RPCTimeout                     = 3 * time.Second
	RelayerBytesLength             = 48
	UpdateCachedValidatorsInterval = 1 * time.Minute
)

var (
	BSCBalanceThreshold = big.NewInt(1000000000000000000) // when relayer is lower than 1BNB, it should try to claim rewards
	BSCRewardThreshold  = big.NewInt(100000000000000000)  // if reward is lower than 0.1 BNB, it will not be claimed.
)
