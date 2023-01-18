package listener

import (
	"context"
	"fmt"
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
	config              *config.Config
	bscExecutor         *executor.BSCExecutor
	inscriptionExecutor *executor.InscriptionExecutor
	DaoManager          *dao.DaoManager
	crossChainAbi       abi.ABI
}

func NewBSCListener(cfg *config.Config, bscExecutor *executor.BSCExecutor, insExecutor *executor.InscriptionExecutor, dao *dao.DaoManager) *BSCListener {
	crossChainAbi, err := abi.JSON(strings.NewReader(crosschain.CrosschainMetaData.ABI))
	if err != nil {
		panic("marshal abi error")
	}
	return &BSCListener{
		config:              cfg,
		bscExecutor:         bscExecutor,
		inscriptionExecutor: insExecutor,
		DaoManager:          dao,
		crossChainAbi:       crossChainAbi,
	}
}

func (l *BSCListener) Start() {
	curHeight := l.config.BSCConfig.StartHeight
	for {
		nextHeight, err := l.poll(curHeight)
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		}
		curHeight = nextHeight
	}
}

func (l *BSCListener) poll(blockHeight uint64) (uint64, error) {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block from db, error: %s", err.Error())
		return 0, err
	}
	if (*latestPolledBlock != model.BscBlock{}) {
		latestPolledBlockHeight := latestPolledBlock.Height
		if blockHeight <= latestPolledBlockHeight {
			blockHeight = latestPolledBlockHeight + 1
		}

		latestBlockHeight, err := l.bscExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("failed to get latest block blockHeight, error: %s", err.Error())
			return 0, err
		}

		if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
			return blockHeight, nil
		}
	}

	err = l.monitorCrossChainPkgAtBlockHeight(latestPolledBlock, blockHeight)
	if err != nil {
		relayercommon.Logger.Errorf("encounter error when monitor cross-chain packages at blockHeight %d, err=%s", blockHeight, err.Error())
		return 0, err
	}
	return blockHeight + 1, nil
}

func (l *BSCListener) getLatestPolledBlock() (*model.BscBlock, error) {
	block, err := l.DaoManager.BSCDao.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (l *BSCListener) monitorCrossChainPkgAtBlockHeight(latestPolledBlock *model.BscBlock, curHeight uint64) error {
	relayercommon.Logger.Infof("retrieve BSC block header at height=%d", curHeight)
	nextHeightHeader, err := l.bscExecutor.GetBlockHeaderAtHeight(curHeight)
	if err != nil {
		return err
	}
	if nextHeightHeader == nil {
		relayercommon.Logger.Infof("BSC header at height %d not found", curHeight)
		return nil
	}

	isForked, err := l.validateLatestPolledBlockIsForkedAndDelete(latestPolledBlock, nextHeightHeader)
	if err != nil {
		return err
	}
	if isForked {
		return fmt.Errorf("There is fork at current  block height, height=%d", curHeight)
	}
	logs, err := l.getLogsFromHeader(nextHeightHeader)
	if err != nil {
		return fmt.Errorf("failed to get logs at height, height=%d, err=%s", curHeight, err.Error())
	}
	relayPkgs := make([]*model.BscRelayPackage, 0)
	for _, log := range logs {
		relayercommon.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())
		relayPkg, err := ParseRelayPackage(&l.crossChainAbi, &log, nextHeightHeader.Time, relayercommon.ChainId(l.config.InscriptionConfig.ChainId), relayercommon.ChainId(l.config.BSCConfig.ChainId))
		if err != nil {
			relayercommon.Logger.Errorf("failed to parse event log, txHash=%s, err=%s", log.TxHash, err.Error())
			continue
		}
		//TODO  remove after testing
		nextDeliverySeqOnInscription, err := l.bscExecutor.GetNextDeliveryOracleSequence()
		relayPkg.OracleSequence = nextDeliverySeqOnInscription
		seq, err := l.inscriptionExecutor.GetNextReceiveSequenceForChannel(relayercommon.ChannelId(relayPkg.ChannelId))
		relayPkg.OracleSequence = nextDeliverySeqOnInscription
		relayPkg.PackageSequence = seq

		if relayPkg == nil {
			continue
		}
		relayPkgs = append(relayPkgs, relayPkg)
	}

	b := &model.BscBlock{
		BlockHash:  nextHeightHeader.Hash().String(),
		ParentHash: nextHeightHeader.ParentHash.String(),
		Height:     curHeight,
		BlockTime:  int64(nextHeightHeader.Time),
	}
	return l.DaoManager.BSCDao.SaveBlockAndBatchPackages(b, relayPkgs)
}

func (l *BSCListener) getLogsFromHeader(header *types.Header) ([]types.Log, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client := l.bscExecutor.GetRpcClient()
	topics := [][]ethereumcommon.Hash{{CrossChainPackageEventHash}}
	blockHash := header.Hash()
	logs, err := client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethereumcommon.Address{executor.CrossChainContractAddr},
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
		err := l.DaoManager.BSCDao.DB.Transaction(func(tx *gorm.DB) error {
			err := l.DaoManager.BSCDao.DeleteBlockAtHeight(latestPolledBlock.Height)
			if err != nil {
				return err
			}
			err = l.DaoManager.BSCDao.DeletePackagesAtHeight(latestPolledBlock.Height)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return true, err
		}
		relayercommon.Logger.Infof("deleted block at height %d from DB due to it is forked", latestPolledBlock.Height)
		return true, nil
	}
	return false, nil
}
