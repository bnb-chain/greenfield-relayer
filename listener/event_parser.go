package listener

import (
	"encoding/hex"
	"math/big"

	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/model"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	CrossChainPackageEventName = "CrossChainPackage"
	// TODO
	CrossChainPackageEventHash = common.HexToHash("")
)

func ParseRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64) (*model.BscRelayPackage, error) {
	ev, err := parseCrossChainPackageEvent(abi, log)
	if err != nil {
		return nil, err
	}
	var p model.BscRelayPackage
	p.OracleSequence = ev.OracleSequence
	p.PackageSequence = ev.PackageSequence
	p.ChannelId = ev.ChannelId
	p.TxHash = log.TxHash.String()
	p.TxIndex = log.TxIndex
	p.TxTime = int64(timestamp)
	p.UpdatedTime = int64(timestamp)
	p.Height = log.BlockNumber
	p.Status = db.SAVED
	p.PayLoad = hex.EncodeToString(ev.Payload)
	return &p, nil
}

func parseCrossChainPackageEvent(abi *abi.ABI, log *types.Log) (*CrossChainPackageEvent, error) {
	var ev CrossChainPackageEvent

	err := abi.UnpackIntoInterface(&ev, CrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.OracleSequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	ev.PackageSequence = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	ev.ChannelId = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())
	return &ev, nil
}
