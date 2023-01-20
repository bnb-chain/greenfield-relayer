package vote

import "time"

const (
	RetryInterval = 1 * time.Second

	QueryVotepoolMaxRetryTimes = 10

	VotePoolBroadcastMethodName   = "broadcast_vote"
	VotePoolBroadcastParameterKey = "vote"

	VotePoolQueryMethodName         = "query_vote"
	VotePoolQueryParameterEventType = "event_type"
	VotePoolQueryParameterEventHash = "event_hash"
	VotePoolQueryRetryInterval      = 2 * time.Second
)
