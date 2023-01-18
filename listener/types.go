package listener

type CrossChainPackageEvent struct {
	SrcChainId      uint32
	DstChainId      uint32
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
}
