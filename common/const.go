package common

import (
	"github.com/bnb-chain/greenfield-relayer/types"
	"time"

	"github.com/avast/retry-go/v4"
)

var (
	RtyAttNum     = uint(5)
	RtyAttem      = retry.Attempts(RtyAttNum)
	RtyDelay      = retry.Delay(time.Millisecond * 500)
	RtyErr        = retry.LastErrorOnly(true)
	RetryInterval = 1 * time.Second
)

const (
	OracleChannelId types.ChannelId = 0
)
