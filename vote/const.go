package vote

type Packages []Package

type Package struct {
	ChannelId uint8
	Sequence  uint64
	Payload   []byte
}
