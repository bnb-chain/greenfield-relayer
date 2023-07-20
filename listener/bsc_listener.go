package listener

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/executor/crosschain"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/metric"
	rtypes "github.com/bnb-chain/greenfield-relayer/types"
)

type BSCListener struct {
	config             *config.Config
	bscExecutor        *executor.BSCExecutor
	greenfieldExecutor *executor.GreenfieldExecutor
	DaoManager         *dao.DaoManager
	crossChainAbi      abi.ABI
	monitorService     *metric.MetricService
}

func NewBSCListener(cfg *config.Config, bscExecutor *executor.BSCExecutor, gnfdExecutor *executor.GreenfieldExecutor, dao *dao.DaoManager, ms *metric.MetricService) *BSCListener {
	crossChainAbi, err := abi.JSON(strings.NewReader(crosschain.CrosschainMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	return &BSCListener{
		config:             cfg,
		bscExecutor:        bscExecutor,
		greenfieldExecutor: gnfdExecutor,
		DaoManager:         dao,
		crossChainAbi:      crossChainAbi,
		monitorService:     ms,
	}
}

func (l *BSCListener) StartLoop() {
	for {
		err := l.poll()
		if err != nil {
			time.Sleep(common.ErrorRetryInterval)
			continue
		}
	}
}

func (l *BSCListener) poll() error {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		logging.Logger.Errorf("failed to get latest polled block from db, error: %s", err.Error())
		return err
	}
	nextHeight := l.config.BSCConfig.StartHeight
	if (*latestPolledBlock != model.BscBlock{}) {
		latestPolledBlockHeight := latestPolledBlock.Height
		if nextHeight <= latestPolledBlockHeight {
			nextHeight = latestPolledBlockHeight + 1
		}

		latestBlockHeight, err := l.bscExecutor.GetLatestFinalizedBlockHeightWithRetry()
		if err != nil {
			logging.Logger.Errorf("failed to get latest blockHeight, error: %s", err.Error())
			return err
		}
		if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
			time.Sleep(common.ListenerPauseTime)
			return nil
		}
	}
	if err = l.monitorCrossChainPkgAt(nextHeight, latestPolledBlock); err != nil {
		logging.Logger.Errorf("encounter error when monitor cross-chain packages at blockHeight=%d, err=%s", nextHeight, err.Error())
		return err
	}
	return nil
}

func (l *BSCListener) getLatestPolledBlock() (*model.BscBlock, error) {
	return l.DaoManager.BSCDao.GetLatestBlock()
}

func (l *BSCListener) monitorCrossChainPkgAt(nextHeight uint64, latestPolledBlock *model.BscBlock) error {
	nextHeightBlockHeader, err := l.bscExecutor.GetBlockHeaderAtHeight(nextHeight)
	if err != nil {
		return err
	}
	if nextHeightBlockHeader == nil {
		logging.Logger.Infof("BSC Block header at height %d not found", nextHeight)
		return nil
	}
	logging.Logger.Infof("retrieved BSC block header at height=%d", nextHeight)
	// check if the latest polled block in DB is forked, if so, delete it.
	isForked, err := l.isForkedBlockAndDelete(latestPolledBlock, nextHeight, nextHeightBlockHeader.ParentHash)
	if err != nil {
		return err
	}
	if isForked {
		return fmt.Errorf("there is fork at block height=%d", latestPolledBlock.Height)
	}
	logs, err := l.queryCrossChainLogs(nextHeightBlockHeader.Hash())
	if err != nil {
		return fmt.Errorf("failed to get logs from block at height=%d, err=%s", nextHeight, err.Error())
	}
	relayPkgs := make([]*model.BscRelayPackage, 0)
	for _, log := range logs {
		logging.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
		relayPkg, err := ParseRelayPackage(&l.crossChainAbi,
			&log, nextHeightBlockHeader.Time,
			rtypes.ChainId(l.config.GreenfieldConfig.ChainId),
			rtypes.ChainId(l.config.BSCConfig.ChainId),
			&l.config.RelayConfig,
		)
		if err != nil {
			logging.Logger.Errorf("failed to parse event log, txHash=%s, err=%s", log.TxHash, err.Error())
			continue
		}

		if relayPkg == nil {
			continue
		}
		relayPkgs = append(relayPkgs, relayPkg)
	}

	if err := l.DaoManager.BSCDao.SaveBlockAndBatchPackages(
		&model.BscBlock{
			BlockHash:  nextHeightBlockHeader.Hash().String(),
			ParentHash: nextHeightBlockHeader.ParentHash.String(),
			Height:     nextHeight,
			BlockTime:  int64(nextHeightBlockHeader.Time),
		}, relayPkgs); err != nil {
		return err
	}
	l.monitorService.SetBSCSavedBlockHeight(nextHeight)
	return nil
}

func (l *BSCListener) queryCrossChainLogs(blockHash ethcommon.Hash) ([]types.Log, error) {
	client := l.bscExecutor.GetRpcClient()
	topics := [][]ethcommon.Hash{{l.getCrossChainPackageEventHash()}}
	logs, err := client.FilterLogs(context.Background(), ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethcommon.Address{l.getCrossChainContractAddress()},
	})
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (l *BSCListener) isForkedBlockAndDelete(latestPolledBlock *model.BscBlock, nextHeight uint64, parentHash ethcommon.Hash) (bool, error) {
	if latestPolledBlock.Height != 0 &&
		latestPolledBlock.Height+1 == nextHeight &&
		parentHash.String() != latestPolledBlock.BlockHash {
		// delete latestPolledBlock and its cross-chain packages from DB
		if err := l.DaoManager.BSCDao.DeleteBlockAndPackagesAtHeight(latestPolledBlock.Height); err != nil {
			return true, err
		}
		logging.Logger.Infof("deleted block at height=%d from DB due to it is forked", latestPolledBlock.Height)
		return true, nil
	}
	return false, nil
}

func (l *BSCListener) getCrossChainPackageEventHash() ethcommon.Hash {
	return ethcommon.HexToHash(l.config.RelayConfig.CrossChainPackageEventHex)
}

func (l *BSCListener) getCrossChainContractAddress() ethcommon.Address {
	return ethcommon.HexToAddress(l.config.RelayConfig.CrossChainContractAddr)
}
