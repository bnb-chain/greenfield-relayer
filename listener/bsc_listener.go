package listener

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	bscabi "inscription-relayer/executor/abi"
	"strings"
	"time"
)

type BSCListener struct {
	config        *config.Config
	bscExecutor   *executor.BSCExecutor
	daoManager    *dao.DaoManager
	CrossChainAbi abi.ABI
}

func NewBSCListener(cfg *config.Config, executor *executor.BSCExecutor, dao *dao.DaoManager) *BSCListener {
	crossChainAbi, err := abi.JSON(strings.NewReader(bscabi.CrossChainABI))
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
			relayercommon.Logger.Errorf("Failed to get latest block from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}
		nextBlockHeight := startHeight
		latestPolledBlockHeight := latestPolledBlock.Height
		if startHeight != 0 && startHeight <= latestPolledBlockHeight {
			relayercommon.Logger.Infof("block at height %d has been polled, current DB block height %d", startHeight, latestPolledBlockHeight)
			nextBlockHeight = latestPolledBlockHeight + 1
		}
		relayercommon.Logger.Infof("start from height %s ", nextBlockHeight)

		latestBlockHeight, err := l.bscExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get latest block height, error: %s", err.Error())
			time.Sleep(GetBlockHeightRetryInterval)
			continue
		}

		if latestPolledBlockHeight == latestBlockHeight-1 {
			relayercommon.Logger.Infof("Pause polling Block, current block height in db : %d, block height is: %d", latestPolledBlock, latestBlockHeight)
			time.Sleep(GetBlockHeightRetryInterval)
			continue
		}
		nextBlockHeight = latestPolledBlockHeight + 1
		err = l.monitorCrossChainPkgAtBlockHeight(latestPolledBlock, nextBlockHeight)
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

func (l *BSCListener) monitorCrossChainPkgAtBlockHeight(latestPolledBlock *model.BscBlock, nextBlockHeight uint64) error {
	relayercommon.Logger.Infof("Retrieve block nextHeightHeader at height=%d", nextBlockHeight)
	nextHeightHeader, err := l.bscExecutor.GetBlockHeaderAtHeight(nextBlockHeight)
	isForked, err := l.validateLatestPolledBlockIsForkedAndDelete(latestPolledBlock, nextHeightHeader)
	if err != nil {
		return err
	}
	if isForked {
		relayercommon.Logger.Infof("Deleted block at height %d from DB due to it is forked", latestPolledBlock.Height)
		return nil
	}
	logs, err := l.bscExecutor.GetLogsFromHeader(nextHeightHeader)
	if err != nil {
		return fmt.Errorf("failed to get logs at height, height=%d, err=%s", nextBlockHeight, err.Error())
	}
	if err != nil {
		return err
	}

	relayPkgs := make([]*model.BscRelayPackage, 0)
	for _, log := range logs {
		relayercommon.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())

		//TODO Convert to transaction directly
		relayPkg, err := executor.ParseRelayPackage(&l.CrossChainAbi, &log, nextHeightHeader.Time)
		if err != nil {
			relayercommon.Logger.Errorf("parse event log error, er=%s", err.Error())
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
		Height:     nextBlockHeight,
		BlockTime:  int64(nextHeightHeader.Time),
	}
	err = l.daoManager.BSCDao.SaveBlockAndBatchPackages(b, relayPkgs)
	return err
}

func (l *BSCListener) validateLatestPolledBlockIsForkedAndDelete(latestProcessedBlock *model.BscBlock, header *types.Header) (bool, error) {
	parentBlockHash := header.ParentHash

	if latestProcessedBlock.Height != 0 && parentBlockHash.String() != latestProcessedBlock.BlockHash {
		//delete latestProcessedBlock from DB
		dbTx := l.daoManager.BSCDao.DB.Begin()

		err := l.daoManager.BSCDao.DeleteBlockAtHeight(latestProcessedBlock.Height)
		if err != nil {
			dbTx.Rollback()
			return true, err
		}
		err = l.daoManager.BSCDao.DeletePackagesAtHeight(latestProcessedBlock.Height)
		if err != nil {
			dbTx.Rollback()
			return true, err
		}
		return true, dbTx.Commit().Error
	}
	return false, nil
}
