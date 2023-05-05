package executor

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/eth/ethsecp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
)

func Cdc() *codec.ProtoCodec {
	interfaceRegistry := types.NewInterfaceRegistry()
	interfaceRegistry.RegisterInterface("AccountI", (*authtypes.AccountI)(nil))
	interfaceRegistry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&authtypes.BaseAccount{},
	)
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &oracletypes.MsgClaim{})
	return codec.NewProtoCodec(interfaceRegistry)
}
