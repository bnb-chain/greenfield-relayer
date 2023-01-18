package listener

import (
	"time"
)

const (
	RetryInterval = 1 * time.Second

	EventTypeCrossChain = "cosmos.crosschain.v1.EventCrossChain"
)
