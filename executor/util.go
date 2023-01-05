package executor

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	"github.com/evmos/ethermint/crypto/ethsecp256k1"
	"inscription-relayer/common"
)

func BuildChannelSequenceKey(destChainId common.ChainId, chanelId common.ChannelId) []byte {
	key := make([]byte, prefixLength+destChainIDLength+channelIDLength)
	copy(key[:prefixLength], prefixForSequenceKey)
	binary.BigEndian.PutUint16(key[prefixLength:prefixLength+destChainIDLength], uint16(destChainId))
	copy(key[prefixLength+destChainIDLength:], []byte{byte(chanelId)})
	return key
}

func isHeaderNonExistingErr(err error) bool {
	if err != nil && err.Error() == "RPC error -32603 - Internal error: Height must be less than or equal to the current blockchain height" {
		return true
	}
	return false
}

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
