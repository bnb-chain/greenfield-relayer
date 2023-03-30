package common

import (
	"time"

	"github.com/avast/retry-go/v4"

	"github.com/bnb-chain/greenfield-relayer/types"
)

var (
	RtyAttNum = uint(5)
	RtyAttem  = retry.Attempts(RtyAttNum)
	RtyDelay  = retry.Delay(time.Millisecond * 500)
	RtyErr    = retry.LastErrorOnly(true)
)

const (
	OracleChannelId              types.ChannelId = 0
	SleepTimeAfterSyncLightBlock                 = 15 * time.Second

	ListenerPauseTime   = 2 * time.Second
	ErrorRetryInterval  = 1 * time.Second
	BroadcastInterval   = 500 * time.Millisecond
	CollectVoteInterval = 500 * time.Millisecond
	AssembleInterval    = 500 * time.Millisecond
)
