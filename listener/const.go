package listener

import "time"

const (
	GetBlockHeightRetryInterval = 2 * time.Second

	EventTypeCrossChainPackage         = "EventCrossChain"
	EventAttributeKeyCrossChainPackage = "destChainID,channelID,sequence"
)
