package listener

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"sync"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cometbft/cometbft/votepool"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/bnb-chain/greenfield-relayer/metric"
	"github.com/bnb-chain/greenfield-relayer/util"
)

type GreenfieldListener struct {
	config             *config.Config
	greenfieldExecutor *executor.GreenfieldExecutor
	bscExecutor        *executor.BSCExecutor
	DaoManager         *dao.DaoManager
	metricService      *metric.MetricService
}

func NewGreenfieldListener(cfg *config.Config, gnfdExecutor *executor.GreenfieldExecutor, bscExecutor *executor.BSCExecutor,
	dao *dao.DaoManager, ms *metric.MetricService) *GreenfieldListener {
	return &GreenfieldListener{
		config:             cfg,
		greenfieldExecutor: gnfdExecutor,
		bscExecutor:        bscExecutor,
		DaoManager:         dao,
		metricService:      ms,
	}
}

func (l *GreenfieldListener) StartLoop() {
	for {
		if err := l.poll(); err != nil {
			logging.Logger.Errorf("encounter err, err=%s", err.Error())
			time.Sleep(common.ErrorRetryInterval)
			continue
		}
	}
}

func (l *GreenfieldListener) poll() error {
	nextHeight, err := l.calNextHeight()
	if err != nil {
		return fmt.Errorf("failed to cal next height, error: %s", err.Error())
	}
	blockResults, block, err := l.getBlockAndBlockResult(nextHeight)
	if err != nil {
		return fmt.Errorf("failed to get block and block result at height %d, error: %s", nextHeight, err.Error())
	}
	txs := make([]*model.GreenfieldRelayTransaction, 0)
	wg := new(sync.WaitGroup)
	wg.Add(3)
	relayTxCh := make(chan *model.GreenfieldRelayTransaction)
	errChan := make(chan error)
	waitCh := make(chan struct{})

	go func() {
		go l.monitorTxEvents(block, blockResults.TxsResults, relayTxCh, errChan, wg)
		go l.monitorEndBlockEvents(uint64(block.Height), blockResults.EndBlockEvents, relayTxCh, errChan, wg)
		go l.monitorValidators(block, errChan, wg)
		wg.Wait()
		close(waitCh)
	}()

	for {
		select {
		case err := <-errChan:
			return fmt.Errorf("encounter error when monitoring block at Height=%d, err=%s", nextHeight, err.Error())
		case tx := <-relayTxCh:
			txs = append(txs, tx)
		case <-waitCh:
			b := &model.GreenfieldBlock{
				Chain:     block.ChainID,
				Height:    uint64(block.Height),
				BlockTime: block.Time.Unix(),
			}
			if err := l.DaoManager.GreenfieldDao.SaveBlockAndBatchTransactions(b, txs); err != nil {
				return fmt.Errorf("failed to persist block and tx to DB, err=%s", err.Error())
			}
			l.metricService.SetGnfdSavedBlockHeight(uint64(block.Height))
			return nil
		}
	}
}

func (l *GreenfieldListener) getLatestPolledBlock() (*model.GreenfieldBlock, error) {
	return l.DaoManager.GreenfieldDao.GetLatestBlock()
}

func (l *GreenfieldListener) getBlockAndBlockResult(height uint64) (*ctypes.ResultBlockResults, *tmtypes.Block, error) {
	block, blockResults, err := l.greenfieldExecutor.GetBlockAndBlockResultAtHeight(int64(height))
	if err != nil {
		return nil, nil, err
	}
	logging.Logger.Infof("retrieved greenfield block at height=%d", height)
	return blockResults, block, nil
}

func (l *GreenfieldListener) monitorTxEvents(block *tmtypes.Block, txRes []*abci.ResponseDeliverTx, txChan chan *model.GreenfieldRelayTransaction, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	// Cross chain Transfer events
	for idx, tx := range txRes {
		for _, event := range tx.Events {
			if event.Type == GreenfieldEventTypeCrossChain {
				relayTx, err := constructRelayTx(event, uint64(block.Height))
				if err != nil {
					errChan <- err
					return
				}
				if relayTx.DestChainId != l.destChainId() {
					break
				}
				relayTx.TxHash = hex.EncodeToString(block.Txs[idx].Hash())
				txChan <- relayTx
			}
		}
	}
}

func (l *GreenfieldListener) monitorEndBlockEvents(height uint64, endBlockEvents []abci.Event, txChan chan *model.GreenfieldRelayTransaction, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, e := range endBlockEvents {
		if e.Type == GreenfieldEventTypeCrossChain {
			relayTx, err := constructRelayTx(e, height)
			if err != nil {
				errChan <- err
				return
			}
			if relayTx.DestChainId != l.destChainId() {
				break
			}
			txChan <- relayTx
		}
	}
}

func (l *GreenfieldListener) monitorValidators(block *tmtypes.Block, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := l.monitorValidatorsHelper(block); err != nil {
		errChan <- err
	}
}

func (l *GreenfieldListener) monitorValidatorsHelper(block *tmtypes.Block) error {
	lightClientLatestHeight, err := l.bscExecutor.GetLightClientLatestHeight()
	if err != nil {
		return err
	}
	nextHeight := uint64(block.Height)
	// happen when re-process block
	if nextHeight <= lightClientLatestHeight {
		return nil
	}

	latestSyncedLightBlockTx, err := l.DaoManager.GreenfieldDao.GetLatestSyncedTransaction()
	if err != nil {
		return err
	}
	latestValidatorsHashFromDB, err := hex.DecodeString(latestSyncedLightBlockTx.ValidatorsHash)
	if err != nil {
		return err
	}

	if bytes.Equal(block.ValidatorsHash[:], latestValidatorsHashFromDB) {
		return nil
	}
	nextValidators, err := l.greenfieldExecutor.QueryValidatorsAtHeight(nextHeight)
	if err != nil {
		return err
	}

	curValidators, err := l.greenfieldExecutor.QueryValidatorsAtHeight(nextHeight - 1)
	if err != nil {
		return err
	}

	if len(nextValidators) != len(curValidators) {
		if err = l.sync(nextHeight, block.ValidatorsHash.String()); err != nil {
			return err
		}
		return nil
	}
	for idx, nextVal := range nextValidators {
		curVal := curValidators[idx]

		if !bytes.Equal(nextVal.Address.Bytes(), curVal.Address.Bytes()) ||
			!bytes.Equal(nextVal.BlsKey, curVal.BlsKey) ||
			!bytes.Equal(nextVal.RelayerAddress, curVal.RelayerAddress) {

			if err = l.sync(nextHeight, block.ValidatorsHash.String()); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

func (l *GreenfieldListener) calNextHeight() (uint64, error) {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block from db, error: %s", err.Error())
	}
	latestPolledBlockHeight := latestPolledBlock.Height

	nextHeight := l.config.GreenfieldConfig.StartHeight
	if nextHeight <= latestPolledBlockHeight {
		nextHeight = latestPolledBlockHeight + 1
	}

	latestBlockHeight, err := l.greenfieldExecutor.GetLatestBlockHeight()
	if err != nil {
		return 0, fmt.Errorf("failed to get latest block height, error: %s", err.Error())
	}
	// pauses relayer for a bit since it already caught the newest block
	if int64(nextHeight) >= int64(latestBlockHeight) {
		time.Sleep(common.ListenerPauseTime)
		return nextHeight, nil
	}
	return nextHeight, nil
}

func (l *GreenfieldListener) sync(nextHeight uint64, validatorsHash string) error {
	logging.Logger.Infof("syncing tendermint light block at height %d", nextHeight)
	txHash, err := l.bscExecutor.SyncTendermintLightBlock(nextHeight)
	if err != nil {
		return fmt.Errorf("failed to sync light block at height=%d, err=%s", nextHeight, err.Error())
	}
	t := &model.SyncLightBlockTransaction{
		ValidatorsHash: validatorsHash,
		Height:         nextHeight,
		TxHash:         txHash.String(),
	}
	if err = l.DaoManager.GreenfieldDao.SaveSyncLightBlockTransaction(t); err != nil {
		return fmt.Errorf("failed to save sync light block transaction to DB, err=%s", err.Error())
	}
	logging.Logger.Infof("synced tendermint light block at height %d with txHash %s", nextHeight, txHash.String())
	time.Sleep(common.SleepTimeAfterSyncLightBlock)
	return nil
}

func constructRelayTx(event abci.Event, height uint64) (*model.GreenfieldRelayTransaction, error) {
	relayTx := model.GreenfieldRelayTransaction{}
	for _, attr := range event.Attributes {
		switch attr.Key {
		case "channel_id":
			chanelId, err := strconv.ParseInt(attr.Value, 10, 8)
			if err != nil {
				return nil, err
			}
			relayTx.ChannelId = uint8(chanelId)
		case "src_chain_id":
			srcChainId, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			relayTx.SrcChainId = uint32(srcChainId)
		case "dest_chain_id":
			destChainId, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			relayTx.DestChainId = uint32(destChainId)
		case "package_load":
			payloadStr, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			relayTx.PayLoad = payloadStr
		case "sequence":
			seq, err := util.QuotedStrToIntWithBitSize(attr.Value, 64)
			if err != nil {
				return nil, err
			}
			relayTx.Sequence = seq
		case "package_type":
			packType, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return nil, err
			}
			relayTx.PackageType = uint32(packType)
		case "timestamp":
			ts, err := util.QuotedStrToIntWithBitSize(attr.Value, 64)
			if err != nil {
				return nil, err
			}
			relayTx.TxTime = int64(ts)
		case "relayer_fee":
			feeStr, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			relayTx.RelayerFee = feeStr
		case "ack_relayer_fee":
			feeStr, err := strconv.Unquote(attr.Value)
			if err != nil {
				return nil, err
			}
			relayTx.AckRelayerFee = feeStr
		default:
			logging.Logger.Errorf("unexpected attr, key is %s", attr.Key)
		}
	}
	relayTx.Status = db.Saved
	relayTx.Height = height
	relayTx.UpdatedTime = time.Now().Unix()
	return &relayTx, nil
}

func (l *GreenfieldListener) PurgeLoop() {
	ticker := time.NewTicker(PurgeJobInterval)
	for range ticker.C {
		latestGnfdBlock, err := l.DaoManager.GreenfieldDao.GetLatestBlock()
		if err != nil {
			logging.Logger.Errorf("failed to get latest DB BSC block, err=%s", err.Error())
			continue
		}
		threshHold := int64(latestGnfdBlock.Height) - NumOfHistoricalBlocks
		if threshHold <= 0 {
			continue
		}
		if err = l.DaoManager.GreenfieldDao.DeleteBlocksBelowHeight(threshHold); err != nil {
			logging.Logger.Errorf("failed to delete gnfd blocks, err=%s", err.Error())
			continue
		}
		exists, err := l.DaoManager.GreenfieldDao.ExistsUnprocessedTransaction(threshHold)
		if err != nil || exists {
			continue
		}
		if err = l.DaoManager.GreenfieldDao.DeleteTransactionsBelowHeightWithLimit(threshHold, DeletionLimit); err != nil {
			logging.Logger.Errorf("failed to delete gnfd transactions, err=%s", err.Error())
			continue
		}
		var eventType votepool.EventType
		if l.config.BSCConfig.IsOpCrossChain() {
			eventType = votepool.ToOpCrossChainEvent
		} else {
			eventType = votepool.ToBscCrossChainEvent
		}
		if err = l.DaoManager.VoteDao.DeleteVotesBelowHeightWithLimit(threshHold, uint32(eventType), DeletionLimit); err != nil {
			logging.Logger.Errorf("failed to delete votes, err=%s", err.Error())
		}
	}
}

func (l *GreenfieldListener) destChainId() uint32 {
	return uint32(l.config.BSCConfig.ChainId)
}
