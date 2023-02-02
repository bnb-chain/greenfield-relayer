package types

import (
	"github.com/ethereum/go-ethereum/common"
	tmtypes "github.com/tendermint/tendermint/types"
)

type (
	ChannelId uint8
	ChainId   uint16
)

type Header struct {
	SignedHeader tmtypes.SignedHeader
	Height       uint64
	BlsPubKeys   []byte
	Relayers     []common.Address
}

// Validator queried  from BSC light-client
type Validator struct {
	RelayerAddress common.Address
	BlsPublicKey   []byte
}

type CrossChainPackageEvent struct {
	SrcChainId      uint32
	DstChainId      uint32
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
}
