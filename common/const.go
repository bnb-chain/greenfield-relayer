package common

import (
	"time"

	"github.com/avast/retry-go/v4"
)

var (
	RtyAttNum = uint(5)
	RtyAttem  = retry.Attempts(RtyAttNum)
	RtyDelay  = retry.Delay(time.Millisecond * 400)
	RtyErr    = retry.LastErrorOnly(true)
)
