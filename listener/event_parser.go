package listener

import (
	"github.com/bnb-chain/inscription-relayer/db/model"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	//TODO to be confirmed
	CrossChainPackageEventName = "CrossChainPackage"

	//deploy contract and get address and eventhash
	CrossChainPackageEventHash = common.HexToHash("")
)

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
	p.TxTime = int64(timestamp)
	p.UpdatedTime = int64(timestamp)
	p.Height = log.BlockNumber
	p.Status = model.SAVED
	return &p, nil
}
