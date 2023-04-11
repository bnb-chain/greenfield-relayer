package vote

import (
	"time"
)

const (
	ValidatorsCapacity = 256

	QueryVotepoolMaxRetryTimes = 5

	VotePoolQueryRetryInterval = 300 * time.Millisecond
)
