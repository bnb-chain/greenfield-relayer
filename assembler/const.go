package assembler

import "time"

var (
	RelayWindowInSecond                  = int64(15)
	RelayIntervalBetweenRelayersInSecond = 3
	RetryInterval                        = 1 * time.Second
)
