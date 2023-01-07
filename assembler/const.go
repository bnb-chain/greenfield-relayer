package assembler

var (
	RelayWindowInSecond                  = int64(15)
	RelayIntervalBetweenRelayersInSecond = 3
)

type MsgClaim struct {
	FromAddress    string
	SrcChainId     uint32
	DestChainId    uint32
	Sequence       uint64
	TimeStamp      uint64
	Payload        []byte
	VoteAddressSet []uint64
	AggSignature   []byte
}

type MsgClaimResponse struct {
}
