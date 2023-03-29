package vote

import (
	"time"
)

const (
	ValidatorsCapacity = 256

	QueryVotepoolMaxRetryTimes = 20

	VotePoolQueryRetryInterval = 500 * time.Millisecond
)
