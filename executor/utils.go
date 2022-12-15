package executor

import (
	"encoding/binary"
	"inscription-relayer/common"
)

func BuildChannelSequenceKey(destChainId common.ChainId, chanelId common.ChannelId) []byte {
	key := make([]byte, prefixLength+destChainIDLength+channelIDLength)
	copy(key[:prefixLength], prefixForSequenceKey)
	binary.BigEndian.PutUint16(key[prefixLength:prefixLength+destChainIDLength], uint16(destChainId))
	copy(key[prefixLength+destChainIDLength:], []byte{byte(chanelId)})
	return key
}
