package db

import "time"

const (
	QueryDBRetryInterval = 2 * time.Second
)

type PackageStatus int
