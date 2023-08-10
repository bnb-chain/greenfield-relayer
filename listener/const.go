package listener

import "time"

const (
	NumOfHistoricalBlocks = int64(10000) // NumOfHistoricalBlocks is the number of blocks will be kept in DB, all transactions and votes also kept within this range
	PurgeJobInterval      = time.Minute * 2
)
