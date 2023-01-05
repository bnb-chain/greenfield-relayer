package common

import (
	"github.com/ethereum/go-ethereum/common"
	tmtypes "github.com/tendermint/tendermint/types"
)

type ChannelId uint8
type ChainId uint16

const (
	OracleChannelId ChannelId = 0
)

type Header struct {
	SignedHeader tmtypes.SignedHeader
	Height       uint64
	BlsPubKeys   []byte
	Relayers     []common.Address
}
