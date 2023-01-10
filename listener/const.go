package listener

import (
	"time"
)

const (
	RetryInterval = 1 * time.Second

	EventTypeCrossChain = "cosmos.crosschain.v1.EventCrossChain"
)

type CrossChainPackageEvent struct {
	ChainId         uint16
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
}
