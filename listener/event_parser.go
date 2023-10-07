package listener

import (
	"encoding/hex"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	rtypes "github.com/bnb-chain/greenfield-relayer/types"
)

func ParseRelayPackage(abi *abi.ABI, log *types.Log, timestamp uint64, greenfieldChainId, bscChainId sdk.ChainID) (*model.BscRelayPackage, error) {
	ev, err := parseCrossChainPackageEvent(abi, log)
	if err != nil {
		return nil, err
	}
	if sdk.ChainID(ev.SrcChainId) != bscChainId || sdk.ChainID(ev.DstChainId) != greenfieldChainId {
		return nil, fmt.Errorf("event log's chain id(s) not expected, SrcChainId=%d, DstChainId=%d", ev.SrcChainId, ev.DstChainId)
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

func parseCrossChainPackageEvent(abi *abi.ABI, log *types.Log) (*rtypes.CrossChainPackageEvent, error) {
	var ev rtypes.CrossChainPackageEvent

	err := abi.UnpackIntoInterface(&ev, BSCCrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.OracleSequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	ev.PackageSequence = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	ev.ChannelId = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())
	return &ev, nil
}
