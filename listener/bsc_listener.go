package listener

import (
	"context"
	"fmt"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethereumcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"strings"

	"github.com/bnb-chain/inscription-relayer/executor/crosschain"
	_ "strings"
	"time"
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
	go l.poll(l.config.BSCConfig.StartHeight)
}

func (l *BSCListener) poll(startHeight uint64) {
	for {
		latestPolledBlock, err := l.getLatestPolledBlock()
		if err != nil {
			relayercommon.Logger.Errorf("failed to get latest block from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}

		latestPolledBlockHeight := latestPolledBlock.Height
		if startHeight <= latestPolledBlockHeight {
			startHeight = latestPolledBlockHeight + 1
		}

		latestBlockHeight, err := l.bscExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
			continue
		}

		if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
			continue
		}

		err = l.monitorCrossChainPkgAtBlockHeight(latestPolledBlock, startHeight)
		if err != nil {
			relayercommon.Logger.Errorf("Encounter error when monitorCrossChainPkgAtBlockHeight, err=%s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
		}
	}
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
	if nextHeightHeader == nil {
		relayercommon.Logger.Infof("BSC header at height %d not found", height)
		return nil
	}

	isForked, err := l.validateLatestPolledBlockIsForkedAndDelete(latestPolledBlock, nextHeightHeader)
	if err != nil {
		return err
	}
	if isForked {
		relayercommon.Logger.Infof("Deleted block at height %d from DB due to it is forked", latestPolledBlock.Height)
		return nil
	}
	logs, err := l.getLogsFromHeader(nextHeightHeader)
	if err != nil {
		return fmt.Errorf("failed to get logs at height, height=%d, err=%s", height, err.Error())
	}

	relayPkgs := make([]*model.BscRelayPackage, 0)
	for _, log := range logs {
		relayercommon.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())

		relayPkg, err := ParseRelayPackage(&l.CrossChainAbi, &log, nextHeightHeader.Time)
		if err != nil {
			return fmt.Errorf("failed to parse event log, txHash=%s, err=%s", log.TxHash, err.Error())
			continue
		}
		if relayPkg == nil {
			continue
		}
		relayPkgs = append(relayPkgs, relayPkg)
	}

	b := &model.BscBlock{
		BlockHash:  nextHeightHeader.Hash().String(),
		ParentHash: latestPolledBlock.BlockHash,
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
		//delete latestPolledBlock from DB
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
