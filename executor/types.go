package executor

import "github.com/ethereum/go-ethereum/common"

// Validator queried  from BSC
type Validator struct {
	RelayerAddress common.Address
	BlsPublicKey   []byte
}
