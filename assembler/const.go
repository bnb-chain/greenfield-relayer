package assembler

import "time"

var (
	BSCRelayingDelayInSecond                 = int64(30)
	InscriptionRelayingDelayInSecond         = int64(10)
	FirstInturnRelayerRelayingWindowInSecond = int64(40)
	InturnRelayerRelayingWindowInSecond      = int64(3)
	RetryInterval                            = 1 * time.Second
)
