package listener

import (
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/model"
)

func ParseRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64, inscriptionChainId, bscChainId common.ChainId, config *config.RelayConfig) (*model.BscRelayPackage, error) {
	ev, err := parseCrossChainPackageEvent(abi, log, config)
	if err != nil {
		return nil, err
	}
	if common.ChainId(ev.SrcChainId) != bscChainId || common.ChainId(ev.DstChainId) != inscriptionChainId {
		return nil, errors.New("event log's chain id not expected")
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
	p.Status = db.Saved
	p.PayLoad = hex.EncodeToString(ev.Payload)
	return &p, nil
}

func parseCrossChainPackageEvent(abi *abi.ABI, log *types.Log, config *config.RelayConfig) (*CrossChainPackageEvent, error) {
	var ev CrossChainPackageEvent

	err := abi.UnpackIntoInterface(&ev, config.BSCCrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.OracleSequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	ev.PackageSequence = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	ev.ChannelId = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())
	return &ev, nil
}
