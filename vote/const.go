package vote

import (
	"time"
)

const (
	RetryInterval      = 1 * time.Second
	ValidatorsCapacity = 256

	QueryVotepoolMaxRetryTimes = 5

	VotePoolBroadcastMethodName   = "broadcast_vote"
	VotePoolBroadcastParameterKey = "vote"

	VotePoolQueryMethodName         = "query_vote"
	VotePoolQueryParameterEventType = "event_type"
	VotePoolQueryParameterEventHash = "event_hash"
	VotePoolQueryRetryInterval      = 5 * time.Second // retry every 5 second, total queries 6 times. Approximately equal to a vote expiration time
)
