package executor

import (
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
	GnfdSequenceUpdateLatency      = 3 * time.Second
	BSCSequenceUpdateLatency       = 10 * time.Second

	VotePoolBroadcastMethodName   = "broadcast_vote"
	VotePoolBroadcastParameterKey = "vote"

	VotePoolQueryMethodName         = "query_vote"
	VotePoolQueryParameterEventType = "event_type"
	VotePoolQueryParameterEventHash = "event_hash"
)
