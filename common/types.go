package common

import tmtypes "github.com/tendermint/tendermint/types"

type ChannelId uint8
type ChainId uint16

const (
	OracleChannelId ChannelId = 0
)

type Header struct {
	SignedHeader tmtypes.SignedHeader
	Height       uint64
	BlsPubKeys   []byte
	Relayers     []string
}
