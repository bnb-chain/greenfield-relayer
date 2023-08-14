package executor

import (
	"encoding/binary"
	"fmt"

	"github.com/cometbft/cometbft/crypto/ed25519"
	cbfttypes "github.com/cometbft/cometbft/types"

	"github.com/bnb-chain/greenfield-relayer/config"
)

func GetTestConfig() *config.Config {
	return config.ParseConfigFromFile("../integrationtest/config/config_test.json")
}

func InitExecutors() (*BSCExecutor, *GreenfieldExecutor) {
	cfg := GetTestConfig()
	gnfdExecutor := NewGreenfieldExecutor(cfg)
	bscExecutor := NewBSCExecutor(cfg, nil)
	gnfdExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetGreenfieldExecutor(gnfdExecutor)
	return bscExecutor, gnfdExecutor
}

const (
	chainIDLength              uint64 = 32
	heightLength               uint64 = 8
	validatorSetHashLength     uint64 = 32
	validatorPubkeyLength      uint64 = 32
	validatorVotingPowerLength uint64 = 8
	relayerAddressLength       uint64 = 20
	relayerBlsKeyLength        uint64 = 48
	maxConsensusStateLength    uint64 = 32 * (128 - 1) // FIXME：maximum validator quantity 99
)

type ConsensusState struct {
	ChainID              string
	Height               uint64
	NextValidatorSetHash []byte
	ValidatorSet         *cbfttypes.ValidatorSet
}

// output:
// | chainID   | height   | nextValidatorSetHash | [{validator pubkey, voting power, relayer address, relayer bls pubkey}] |
// | 32 bytes  | 8 bytes  | 32 bytes             | [{32 bytes, 8 bytes, 20 bytes, 48 bytes}]                               |
func (cs ConsensusState) encodeConsensusState() ([]byte, error) {
	validatorSetLength := uint64(len(cs.ValidatorSet.Validators))
	singleValidatorBytesLength := validatorPubkeyLength + validatorVotingPowerLength + relayerAddressLength + relayerBlsKeyLength
	serializeLength := chainIDLength + heightLength + validatorSetHashLength + validatorSetLength*singleValidatorBytesLength
	if serializeLength > maxConsensusStateLength {
		return nil, fmt.Errorf("too many validators %d, consensus state bytes should not exceed %d", len(cs.ValidatorSet.Validators), maxConsensusStateLength)
	}
	encodingBytes := make([]byte, serializeLength)
	pos := uint64(0)
	if uint64(len(cs.ChainID)) > chainIDLength {
		return nil, fmt.Errorf("chainID length should be no more than 32")
	}
	copy(encodingBytes[pos:pos+chainIDLength], cs.ChainID)
	pos += chainIDLength
	binary.BigEndian.PutUint64(encodingBytes[pos:pos+heightLength], cs.Height)
	pos += heightLength
	copy(encodingBytes[pos:pos+validatorSetHashLength], cs.NextValidatorSetHash)
	pos += validatorSetHashLength
	for index := uint64(0); index < validatorSetLength; index++ {
		validator := cs.ValidatorSet.Validators[index]
		pubkey, ok := validator.PubKey.(ed25519.PubKey)
		if !ok {
			return nil, fmt.Errorf("invalid pubkey type")
		}
		copy(encodingBytes[pos:pos+validatorPubkeyLength], pubkey[:])
		pos += validatorPubkeyLength
		binary.BigEndian.PutUint64(encodingBytes[pos:pos+validatorVotingPowerLength], uint64(validator.VotingPower))
		pos += validatorVotingPowerLength
		copy(encodingBytes[pos:pos+relayerAddressLength], validator.RelayerAddress)
		pos += relayerAddressLength
		copy(encodingBytes[pos:pos+relayerBlsKeyLength], validator.BlsKey)
		pos += relayerBlsKeyLength
	}
	return encodingBytes, nil
}
