package db

type TxStatus int

const (
	Saved     TxStatus = 0
	SelfVoted TxStatus = 1 // Tx is only voted by local relayer
	AllVoted  TxStatus = 2 // TX is already voted by all relayer, and more than (2/3) * (# of validators) valid votes collected.
	Delivered TxStatus = 3 // Tx is delivered to the dest chain
)
