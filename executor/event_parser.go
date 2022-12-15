package executor

import (
	"encoding/hex"
	"fmt"
	"inscription-relayer/db/model"
	"inscription-relayer/vote"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	relayercommon "inscription-relayer/common"
)

const (
	MirrorChannelId uint8 = 4
	BCDecimal       uint8 = 8
)

var (
	//TODO to be confirmed
	CrossChainPackageEventName = "crossChainPackage"
	CrossChainPackageEventHash = common.HexToHash("0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe")
)

type MirrorSynPackage struct {
	MirrorSender     common.Address
	ContractAddr     common.Address
	BEP20Name        [32]byte
	BEP20Symbol      [32]byte
	BEP20TotalSupply *big.Int
	BEP20Decimals    uint8
	MirrorFee        *big.Int
	ExpireTime       uint64
}

func DeserializeMirrorSynPackage(serializedPackage []byte) (*MirrorSynPackage, error) {
	var ms MirrorSynPackage
	err := rlp.DecodeBytes(serializedPackage, &ms)
	if err != nil {
		return nil, fmt.Errorf("deserialize mirror package failed")
	}
	return &ms, nil
}

func ParseRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64) (*model.BscRelayPackage, error) {
	var p model.BscRelayPackage

	err := abi.UnpackIntoInterface(&p, CrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}

	p.OracleSequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	p.PackageSequence = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	p.ChannelId = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())
	p.TxHash = log.TxHash.String()
	p.TxIndex = log.TxIndex
	p.BlockHash = log.BlockHash.Hex()
	p.CreatedAt = int64(timestamp)
	p.Height = log.BlockNumber
	p.Status = vote.SAVED

	if p.ChannelId == MirrorChannelId {
		mirrorPackage, err := DeserializeMirrorSynPackage(([]byte)(p.PayLoad)[33:])
		if err != nil {
			return nil, err
		}

		if mirrorPackage.BEP20Decimals > BCDecimal {
			decimals := mirrorPackage.BEP20Decimals - BCDecimal
			mirrorPackage.BEP20TotalSupply = getNewSupply(mirrorPackage.BEP20TotalSupply, decimals)
		}

		encodedBytes, err := rlp.EncodeToBytes(mirrorPackage)
		if err != nil {
			return nil, fmt.Errorf("encode mirror package error")
		}

		newPayLoad := make([]byte, 0)
		newPayLoad = append(newPayLoad, ([]byte)(p.PayLoad)[33:]...)
		newPayLoad = append(newPayLoad, encodedBytes...)
		p.PayLoad = hex.EncodeToString(newPayLoad[:])
		relayercommon.Logger.Infof("mirror payload: %s", p.PayLoad)
	}
	return &p, nil
}

func getBigIntForDecimal(decimal int) *big.Int {
	floatDecimal := big.NewFloat(math.Pow10(decimal))
	bigIntDecimal := new(big.Int)
	floatDecimal.Int(bigIntDecimal)

	return bigIntDecimal
}

func getNewSupply(supply *big.Int, decimals uint8) *big.Int {
	newSupply := big.NewInt(0)
	newSupply = newSupply.Div(supply, getBigIntForDecimal(int(decimals)))
	newSupply = newSupply.Mul(newSupply, getBigIntForDecimal(int(decimals)))

	return newSupply
}
