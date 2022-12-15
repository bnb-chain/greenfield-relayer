package vote

import "time"

const (
	BLSPublicKeyLength               = 48
	BLSSignatureLength               = 96
	VoteSignerTimeout                = 5 * time.Second
	ToBscCrossChainEvent   EventType = 1
	FromBscCrossChainEvent EventType = 2

	SAVED     InternalStatus = 0
	VOTED     InternalStatus = 1
	VOTED_ALL InternalStatus = 2
	FILLED    InternalStatus = 3

	VotepoolRetryInterval = 1 * time.Second
)

type EventType int
type BLSPublicKey [BLSPublicKeyLength]byte
type BLSSignature [BLSSignatureLength]byte
type InternalStatus int

type Packages []Package

type Package struct {
	ChannelId uint8
	Sequence  uint64
	Payload   []byte
}
