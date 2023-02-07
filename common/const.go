package common

import (
	"time"

	"github.com/avast/retry-go/v4"

	"github.com/bnb-chain/greenfield-relayer/types"
)

var (
	RtyAttNum     = uint(5)
	RtyAttem      = retry.Attempts(RtyAttNum)
	RtyDelay      = retry.Delay(time.Millisecond * 500)
	RtyErr        = retry.LastErrorOnly(true)
	RetryInterval = 1 * time.Second
)

const (
	GreenfieldStartHeight                 = 1
	OracleChannelId       types.ChannelId = 0
)
