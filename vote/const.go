package vote

import "time"

const (
	RetryInterval = 1 * time.Second

	VotePoolBroadcastMethodName   = "broadcast_vote"
	VotePoolBroadcastParameterKey = "vote"

	VotePoolQueryMethodName         = "query_vote"
	VotePoolQueryParameterEventType = "event_type"
	VotePoolQueryParameterEventHash = "event_hash"
)

type Packages []Package

type Package struct {
	ChannelId uint8
	Sequence  uint64
	Payload   []byte
}
