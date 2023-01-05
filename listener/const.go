package listener

import "time"

const (
	GetBlockHeightRetryInterval = 2 * time.Second

	EventTypeCrossChain                = "cosmos.crosschain.v1.EventCrossChain"
	EventAttributeKeyCrossChainPackage = "destChainID,channelID,sequence"
)
