package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type (
	ChannelId uint8
	ChainId   uint16
)

// Validator queried from BSC light-client
type Validator struct {
	RelayerAddress common.Address
	BlsPublicKey   []byte
}

// InturnRelayer queired from BSC for the cur inturn relayer and interval
type InturnRelayer struct {
	BlsPublicKey string
	Start        uint64
	End          uint64
}

type CrossChainPackageEvent struct {
	SrcChainId      uint32
	DstChainId      uint32
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
}
