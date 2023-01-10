package listener

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/bnb-chain/inscription-relayer/db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
	"strings"
	"time"

	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"

	"github.com/bnb-chain/inscription-relayer/executor/crosschain"
	_ "strings"
)

type BSCListener struct {
	config        *config.Config
	bscExecutor   *executor.BSCExecutor
	daoManager    *dao.DaoManager
	CrossChainAbi abi.ABI
}

func NewBSCListener(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager) *BSCListener {
	crossChainAbi, err := abi.JSON(strings.NewReader(crosschain.CrosschainMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	return &BSCListener{
		config:        cfg,
		bscExecutor:   executor,
		daoManager:    dao,
		CrossChainAbi: crossChainAbi,
	}
}

func (l *BSCListener) Start() {
	height := l.config.BSCConfig.StartHeight
	for {
		nextHeight, err := l.poll(height)
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		}
		height = nextHeight
	}
}

func (l *BSCListener) poll(height uint64) (uint64, error) {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block from db, error: %s", err.Error())
		return 0, err
	}
	if (*latestPolledBlock != model.BscBlock{}) {
		latestPolledBlockHeight := latestPolledBlock.Height
		if height <= latestPolledBlockHeight {
			height = latestPolledBlockHeight + 1
		}

		latestBlockHeight, err := l.bscExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
			return 0, err
		}

		if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
			return height, nil
		}
	}

	err = l.monitorCrossChainPkgAtBlockHeight(latestPolledBlock, height)
	if err != nil {
		relayercommon.Logger.Errorf("encounter error when monitor cross-chain packages at height %d, err=%s", height, err.Error())
		return 0, err
	}
	return height + 1, nil
}

func (l *BSCListener) getLatestPolledBlock() (*model.BscBlock, error) {
	block, err := l.daoManager.BSCDao.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (l *BSCListener) monitorCrossChainPkgAtBlockHeight(latestPolledBlock *model.BscBlock, height uint64) error {
	relayercommon.Logger.Infof("retrieve BSC block header at height=%d", height)
	nextHeightHeader, err := l.bscExecutor.GetBlockHeaderAtHeight(height)
	if err != nil {
		return err
	}
	if nextHeightHeader == nil {
		relayercommon.Logger.Infof("BSC header at height %d not found", height)
		return nil
	}

	isForked, err := l.validateLatestPolledBlockIsForkedAndDelete(latestPolledBlock, nextHeightHeader)
	if err != nil {
		return err
	}
	if isForked {
		relayercommon.Logger.Infof("deleted block at height %d from DB due to it is forked", latestPolledBlock.Height)
		return nil
	}
	_, err = l.getLogsFromHeader(nextHeightHeader)
	if err != nil {
		return fmt.Errorf("failed to get logs at height, height=%d, err=%s", height, err.Error())
	}

	//TODO remvoe testing purpose code
	relayPkgs := make([]*model.BscRelayPackage, 0)
	pkgSeq := 30
	end := pkgSeq + 5
	ts := time.Now().Unix()
	for i := pkgSeq; i < end; i++ {
		relayPkg := model.BscRelayPackage{}
		relayPkg.ChannelId = 1
		relayPkg.OracleSequence = 10
		relayPkg.PackageSequence = uint64(i)
		relayPkg.PayLoad = hex.EncodeToString(GetPayload(uint64(ts)))
		relayPkg.Height = 0
		relayPkg.TxHash = "hash"
		relayPkg.TxIndex = 1
		relayPkg.Status = db.SAVED
		relayPkg.TxTime = ts
		relayPkg.UpdatedTime = ts
		relayPkgs = append(relayPkgs, &relayPkg)
	}

	//for _, log := range logs {
	//	relayercommon.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
	//	relayPkg, err := ParseRelayPackage(&l.CrossChainAbi, &log, nextHeightHeader.Time)
	//	if err != nil {
	//		return fmt.Errorf("failed to parse event log, txHash=%s, err=%s", log.TxHash, err.Error())
	//	}
	//	if relayPkg == nil {
	//		continue
	//	}
	//}

	b := &model.BscBlock{
		BlockHash:  nextHeightHeader.Hash().String(),
		ParentHash: nextHeightHeader.ParentHash.String(),
		Height:     height,
		BlockTime:  int64(nextHeightHeader.Time),
	}
	return l.daoManager.BSCDao.SaveBlockAndBatchPackages(b, relayPkgs)
}

func (l *BSCListener) getLogsFromHeader(header *types.Header) ([]types.Log, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := l.bscExecutor.GetClient()
	topics := [][]ethereumcommon.Hash{{CrossChainPackageEventHash}}
	blockHash := header.Hash()
	logs, err := client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethereumcommon.Address{l.config.BSCConfig.BSCCrossChainContractAddress},
	})
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (l *BSCListener) validateLatestPolledBlockIsForkedAndDelete(latestPolledBlock *model.BscBlock, header *types.Header) (bool, error) {
	parentBlockHash := header.ParentHash

	if latestPolledBlock.Height != 0 && parentBlockHash.String() != latestPolledBlock.BlockHash {
		// delete latestPolledBlock from DB
		err := l.daoManager.BSCDao.DB.Transaction(func(tx *gorm.DB) error {
			err := l.daoManager.BSCDao.DeleteBlockAtHeight(latestPolledBlock.Height)
			if err != nil {
				return err
			}
			err = l.daoManager.BSCDao.DeletePackagesAtHeight(latestPolledBlock.Height)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return true, err
		}
	}
	return false, nil
}

func GetPayload(ts uint64) []byte {
	payloadHeader := sdk.EncodePackageHeader(sdk.SynCrossChainPackageType, ts, *big.NewInt(1))
	payloadHeader = append(payloadHeader, []byte("test payload")...)
	return payloadHeader
}
