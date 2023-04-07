package vote

import (
	"time"
)

const (
	ValidatorsCapacity = 256

	QueryVotepoolMaxRetryTimes = 5

	VotePoolQueryRetryInterval = 500 * time.Millisecond
)
