package vote

import (
	"time"
)

const (
	RetryInterval      = 1 * time.Second
	ValidatorsCapacity = 256

	QueryVotepoolMaxRetryTimes = 5

	VotePoolQueryRetryInterval = 2 * time.Second
)
