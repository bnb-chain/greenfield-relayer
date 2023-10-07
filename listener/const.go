package listener

import (
	"time"
)

const (
	NumOfHistoricalBlocks         = int64(50000) // NumOfHistoricalBlocks is the number of blocks will be kept in DB, all transactions and votes also kept within this range
	PurgeJobInterval              = time.Minute * 1
	DeletionLimit                 = 10000
	GreenfieldEventTypeCrossChain = "cosmos.crosschain.v1.EventCrossChain"
	BSCCrossChainPackageEventName = "CrossChainPackage"
	CrossChainPackageEventHex     = "0x64998dc5a229e7324e622192f111c691edccc3534bbea4b2bd90fbaec936845a"
)
