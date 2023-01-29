package listener

import (
	"bytes"
	"github.com/bnb-chain/inscription-relayer/logging"
	"strconv"
	"time"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
)

type InscriptionListener struct {
	config              *config.Config
	inscriptionExecutor *executor.InscriptionExecutor
	bscExecutor         *executor.BSCExecutor
	DaoManager          *dao.DaoManager
}

func NewInscriptionListener(cfg *config.Config, insExecutor *executor.InscriptionExecutor, bscExecutor *executor.BSCExecutor, dao *dao.DaoManager) *InscriptionListener {
	return &InscriptionListener{
		config:              cfg,
		inscriptionExecutor: insExecutor,
		bscExecutor:         bscExecutor,
		DaoManager:          dao,
	}
}

func (l *InscriptionListener) Start() {
	go func() {
		for {
			err := l.poll()
			if err != nil {
				time.Sleep(common.RetryInterval)
				continue
			}
		}
	}()
	go func() {
		for {
			err := l.monitorValidators()
			if err != nil {
				time.Sleep(common.RetryInterval)
				continue
			}
		}
	}()
}

func (l *InscriptionListener) poll() error {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		logging.Logger.Errorf("failed to get latest block from db, error: %s", err.Error())
		return err
	}
	nextHeight := l.config.InscriptionConfig.StartHeight
	latestPolledBlockHeight := latestPolledBlock.Height

	if nextHeight <= latestPolledBlockHeight {
		nextHeight = latestPolledBlockHeight + 1
	}

	latestBlockHeight, err := l.inscriptionExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		logging.Logger.Errorf("failed to get latest polled block height, error: %s", err.Error())
		return err
	}
	if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
		time.Sleep(common.RetryInterval)
		return nil
	}
	blockRes, block, err := l.getBlockAndBlockResult(nextHeight)
	if err != nil {
		logging.Logger.Errorf("encounter error when retrieve block and block result at block height=%d, err=%s", nextHeight, err.Error())
		return err
	}

	err = l.monitorCrossChainEvents(blockRes, block)
	if err != nil {
		logging.Logger.Errorf("encounter error when monitor cross-chain events at blockHeight=%d, err=%s", nextHeight, err.Error())
		return err
	}
	return nil
}

func (l *InscriptionListener) getLatestPolledBlock() (*model.InscriptionBlock, error) {
	block, err := l.DaoManager.InscriptionDao.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (l *InscriptionListener) getBlockAndBlockResult(height uint64) (*ctypes.ResultBlockResults, *tmtypes.Block, error) {
	logging.Logger.Infof("retrieve inscription block at height=%d", height)
	blockResults, err := l.inscriptionExecutor.GetBlockResultAtHeight(int64(height))
	if err != nil {
		return nil, nil, err
	}
	block, err := l.inscriptionExecutor.GetBlockAtHeight(int64(height))
	if err != nil {
		return nil, nil, err
	}
	return blockResults, block, nil
}

func (l *InscriptionListener) monitorCrossChainEvents(blockResults *ctypes.ResultBlockResults, block *tmtypes.Block) error {
	txs := make([]*model.InscriptionRelayTransaction, 0)
	for _, tx := range blockResults.TxsResults {
		for _, event := range tx.Events {
			relayTx := model.InscriptionRelayTransaction{}
			if event.Type == l.config.RelayConfig.InscriptionEventTypeCrossChain {
				for _, attr := range event.Attributes {
					switch string(attr.Key) {
					case "channel_id":
						chanelId, err := strconv.ParseInt(string(attr.Value), 10, 8)
						if err != nil {
							return err
						}
						relayTx.ChannelId = uint8(chanelId)
					case "src_chain_id":
						srcChainId, err := strconv.ParseInt(string(attr.Value), 10, 32)
						if err != nil {
							return err
						}
						relayTx.SrcChainId = uint32(srcChainId)
					case "dest_chain_id":
						destChainId, err := strconv.ParseInt(string(attr.Value), 10, 32)
						if err != nil {
							return err
						}
						relayTx.DestChainId = uint32(destChainId)
					case "package_load":
						payloadStr, err := strconv.Unquote(string(attr.Value))
						if err != nil {
							return err
						}
						relayTx.PayLoad = payloadStr
					case "sequence":
						seq, err := util.QuotedStrToIntWithBitSize(string(attr.Value), 64)
						if err != nil {
							return err
						}
						relayTx.Sequence = seq
					case "package_type":
						packType, err := strconv.ParseInt(string(attr.Value), 10, 32)
						if err != nil {
							return err
						}
						relayTx.PackageType = uint32(packType)
					case "timestamp":
						ts, err := util.QuotedStrToIntWithBitSize(string(attr.Value), 64)
						if err != nil {
							return err
						}
						relayTx.TxTime = int64(ts)
					case "relayer_fee":
						feeStr, err := strconv.Unquote(string(attr.Value))
						if err != nil {
							return err
						}
						relayTx.RelayerFee = feeStr
					case "ack_relayer_fee":
						feeStr, err := strconv.Unquote(string(attr.Value))
						if err != nil {
							return err
						}
						relayTx.AckRelayerFee = feeStr
					default:
						logging.Logger.Errorf("unexpected attr, key is %s", attr.Key)
					}
				}
				relayTx.Status = db.Saved
				relayTx.Height = uint64(block.Height)
				relayTx.UpdatedTime = time.Now().Unix()
				txs = append(txs, &relayTx)
			}
		}
	}

	b := &model.InscriptionBlock{
		Chain:     block.ChainID,
		Height:    uint64(block.Height),
		BlockTime: block.Time.Unix(),
	}
	return l.DaoManager.InscriptionDao.SaveBlockAndBatchTransactions(b, txs)
}

func (l *InscriptionListener) monitorValidators() error {
	nextHeight := l.config.InscriptionConfig.StartHeight
	lightClientLatestHeight, err := l.bscExecutor.GetLightClientLatestHeight()
	if err != nil {
		return err
	}
	if nextHeight <= lightClientLatestHeight {
		nextHeight = lightClientLatestHeight + 1
	}
	logging.Logger.Infof("monitoring validator at height %d", nextHeight)
	nextValidators, err := l.inscriptionExecutor.QueryValidatorsAtHeight(nextHeight)
	if err != nil {
		return err
	}
	curValidators, err := l.inscriptionExecutor.QueryValidatorsAtHeight(nextHeight - 1)
	if err != nil {
		return err
	}

	if len(nextValidators) != len(curValidators) {
		logging.Logger.Infof("syncing tendermint header at height %d", nextHeight)
		txHash, err := l.bscExecutor.SyncTendermintLightClientHeader(nextHeight)
		if err != nil {
			return err
		}
		logging.Logger.Infof("synced tendermint header at height %d with txHash %s", nextHeight, txHash.String())
		return nil
	}

	for idx, nextVal := range nextValidators {
		curVal := curValidators[idx]
		if nextVal.OperatorAddress != curVal.OperatorAddress ||
			!bytes.Equal(nextVal.RelayerBlsKey, curVal.RelayerBlsKey) ||
			nextVal.RelayerAddress != curVal.RelayerAddress {
			logging.Logger.Infof("syncing tendermint header at height %d", nextHeight)
			txHash, err := l.bscExecutor.SyncTendermintLightClientHeader(nextHeight)
			if err != nil {
				return err
			}
			logging.Logger.Infof("synced tendermint header at height %d with txHash %s", nextHeight, txHash.String())
			return nil
		}
	}
	return nil
}
