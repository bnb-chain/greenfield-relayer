package assembler

import relayercommon "inscription-relayer/common"

var (
	RelayingWindowInSecond     = int64(15)
	InscriptionMonitorChannels = [3]relayercommon.ChannelId{1, 2, 3}
	BSCMonitorChannels         = [3]relayercommon.ChannelId{1, 2, 3}
)

type MsgClaim struct {
	FromAddress    string
	ChainId        uint16
	Sequence       uint64
	TimeStamp      uint64
	Payload        []byte
	VoteAddressSet []uint64
	AggSignature   []byte
}

type MsgClaimResponse struct {
}
