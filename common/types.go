package common

import tmtypes "github.com/tendermint/tendermint/types"

type ChannelId uint8
type ChainId uint16

type Header struct {
	tmtypes.SignedHeader
	ValidatorSet     *tmtypes.ValidatorSet `json:"validator_set"`
	NextValidatorSet *tmtypes.ValidatorSet `json:"next_validator_set"`
}
