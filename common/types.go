package common

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
