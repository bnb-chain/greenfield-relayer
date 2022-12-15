package listener

import (
	"fmt"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	"time"
)

type InscriptionListener struct {
	config              *config.Config
	inscriptionExecutor *executor.InscriptionExecutor
	daoManager          *dao.DaoManager
}

func NewInscriptionListener(cfg *config.Config, executor *executor.InscriptionExecutor, dao *dao.DaoManager) *InscriptionListener {
	return &InscriptionListener{
		config:              cfg,
		inscriptionExecutor: executor,
		daoManager:          dao,
	}
}

func (l *InscriptionListener) Start() {
	go l.poll(l.config.InscriptionConfig.StartHeight)
}

func (l *InscriptionListener) poll(startHeight uint64) {
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

		latestBlockHeight, err := l.inscriptionExecutor.GetLatestBlockHeightWithRetry()
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
		err = l.monitorCrossChainPkgAtBlockHeight(nextBlockHeight)
		if err != nil {
			relayercommon.Logger.Errorf("Encounter error when monitorCrossChainPkgAtBlockHeight, err=%s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
		}
	}
}

func (l *InscriptionListener) getLatestPolledBlock() (*model.InscriptionBlock, error) {
	block, err := l.daoManager.InscriptionDao.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (l *InscriptionListener) monitorCrossChainPkgAtBlockHeight(height uint64) error {
	relayercommon.Logger.Infof("Retrieve block at height=%d", height)
	blockResults, err := l.inscriptionExecutor.GetBlockResultAtHeight(int64(height))
	block, err := l.inscriptionExecutor.GetBlockAtHeight(int64(height))
	if err != nil {
		return fmt.Errorf("failed to get block at height, height=%d, err=%s", height, err.Error())
	}

	//TODO clarify the event type, attr k-v.
	txs := make([]*model.InscriptionRelayTransaction, 0)
	for _, event := range blockResults.EndBlockEvents {
		//- events:
		//  - attributes:
		//    - key: amount
		//      value: '{"denom":"stake","amount":"1"}'
		//    - key: expire_time
		//      value: '"1671989649"'
		//    - key: from
		//      value: '"cosmos13nhu9fn4yj6r4aqxhqr322k0q26y2m8agvram9"'
		//    - key: relayer_fee
		//      value: '{"denom":"stake","amount":"1"}'
		//    - key: sequence
		//      value: '"1"'
		//    - key: to
		//      value: '"0x72b61c6014342d914470eC7aC2975bE345796c2b"'
		//    type: bnbchain.bfs.bridge.EventCrossTransferOut

		if event.Type == EventTypeCrossChainPackage {
			for _, attr := range event.Attributes {
				if string(attr.Key) != EventAttributeKeyCrossChainPackage {
					continue
				}
			}
			// TODO Convert a event to a transaction model
			tx := model.InscriptionRelayTransaction{}
			txs = append(txs, &tx)
		}
	}

	dbTx := l.daoManager.InscriptionDao.DB.Begin()
	err = l.daoManager.InscriptionDao.SaveBatchTransactions(txs)
	if err != nil {
		dbTx.Rollback()
		return err
	}

	b := &model.InscriptionBlock{
		Chain:     block.ChainID,
		BlockHash: block.Hash().String(),
		Height:    uint64(block.Height),
		BlockTime: block.Time.Unix(),
	}
	err = l.daoManager.InscriptionDao.SaveBlock(b)
	if err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}
