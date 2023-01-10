package db

type TxStatus int

const (
	SAVED     TxStatus = 0
	VOTED     TxStatus = 1
	VOTED_All TxStatus = 2
	FILLED    TxStatus = 3
)
