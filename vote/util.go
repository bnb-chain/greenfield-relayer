package vote

import (
	"time"
)

const (
	ValidatorsCapacity = 256

	QueryVotepoolMaxRetryTimes = 5

	VotePoolQueryRetryInterval = 300 * time.Millisecond
)

func votesEnough(voteCount, validatorsCount int) bool {
	return voteCount >= ceilDiv(validatorsCount*2, 3)
}

func ceilDiv(x, y int) int {
	if y == 0 {
		return 0
	}
	return (x + y - 1) / y
}
