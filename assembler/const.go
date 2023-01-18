package assembler

import "time"

var (
	BSCRelayingDelayInSecond                 = int64(10) // TODO reset to 30 after testing
	InscriptionRelayingDelayInSecond         = int64(10)
	FirstInturnRelayerRelayingWindowInSecond = int64(40)
	InturnRelayerRelayingWindowInSecond      = int64(3)
	RetryInterval                            = 1 * time.Second
)
