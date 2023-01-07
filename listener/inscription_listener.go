package listener

import (
	"bytes"
	"encoding/hex"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
	"strconv"
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
	l.poll(l.config.InscriptionConfig.StartHeight)
}

func (l *InscriptionListener) poll(startHeight uint64) {
	for {
		latestPolledBlock, err := l.getLatestPolledBlock()
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get latest block from db, error: %s", err.Error())
			time.Sleep(db.QueryDBRetryInterval)
			continue
		}
		latestPolledBlockHeight := latestPolledBlock.Height
		if startHeight != 0 && startHeight <= latestPolledBlockHeight {
			startHeight = latestPolledBlockHeight + 1
		}

		latestBlockHeight, err := l.inscriptionExecutor.GetLatestBlockHeightWithRetry()
		if err != nil {
			relayercommon.Logger.Errorf("Failed to get latest block height, error: %s", err.Error())
			continue
		}

		if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
			continue
		}
		blockRes, block, err := l.getBlockAndBlockResult(startHeight)
		if err != nil {
			relayercommon.Logger.Errorf("encounter error when retrieve block and block result at height %d, err=%s", startHeight, err.Error())
			continue
		}

		err = l.monitorCrossChainEvent(blockRes, block)
		if err != nil {
			relayercommon.Logger.Errorf("encounter error when monitor events at block %d, err=%s", startHeight, err.Error())
			continue
		}
		go func() {
			err := l.monitorValidators(startHeight)
			if err != nil {
				relayercommon.Logger.Errorf("encounter error when monitor validators at block %d, err=%s", startHeight, err.Error())
				return
			}
		}()
	}
}

func (l *InscriptionListener) getLatestPolledBlock() (*model.InscriptionBlock, error) {
	block, err := l.daoManager.InscriptionDao.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (l *InscriptionListener) getBlockAndBlockResult(height uint64) (*ctypes.ResultBlockResults, *tmtypes.Block, error) {
	relayercommon.Logger.Infof("retrieve inscription block at height=%d", height)
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

func (l *InscriptionListener) monitorCrossChainEvent(blockResults *ctypes.ResultBlockResults, block *tmtypes.Block) error {
	txs := make([]*model.InscriptionRelayTransaction, 0)

	for _, tx := range blockResults.TxsResults {

		for _, event := range tx.Events {
			relayTx := model.InscriptionRelayTransaction{}
			if event.Type == EventTypeCrossChain {
				for _, attr := range event.Attributes {
					switch string(attr.Key) {
					case "channel_id":
						chanelId, err := strconv.ParseInt(string(attr.Value), 10, 8)
						if err != nil {
							return err
						}
						relayTx.ChannelId = uint8(chanelId)
					case "src_chain_id":
						continue
					case "dest_chain_id":
						continue
					case "package_load":
						relayTx.PayLoad = hex.EncodeToString(attr.Value)
					case "sequence":
						seq, err := util.QuotedStrToIntWithBitSize(string(attr.Value), 64)
						if err != nil {
							return err
						}
						relayTx.Sequence = uint64(seq)
					case "package_type":
						relayTx.Type = string(attr.Value)
					case "timestamp":
						ts, err := util.QuotedStrToIntWithBitSize(string(attr.Value), 64)
						if err != nil {
							return err
						}
						relayTx.TxTime = ts
					case "relayer_fee":
						feeStr, err := strconv.Unquote(string(attr.Value))
						if err != nil {
							return err
						}
						relayTx.RelayerFee = feeStr
					default:
						relayercommon.Logger.Errorf("unexpected attr, key is %s", attr.Key)
					}
				}
				relayTx.Status = model.SAVED
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
	return l.daoManager.InscriptionDao.SaveBlockAndBatchTransactions(b, txs)
}

func (l *InscriptionListener) monitorValidators(height uint64) error {
	if height == 1 {
		return nil
	}

	curValidators, err := l.inscriptionExecutor.QueryValidatorsAtHeight(height)
	if err != nil {
		return err
	}

	prevValidators, err := l.inscriptionExecutor.QueryValidatorsAtHeight(height - 1)
	if err != nil {
		return err
	}

	if len(curValidators) != len(prevValidators) {
		txHash, err := l.inscriptionExecutor.BscExecutor.SyncTendermintLightClientHeader(height)
		if err != nil {
			return err
		}
		relayercommon.Logger.Infof("synced tendermint header at height %d with txHash %s", height, txHash.String())
	}

	for idx, curVal := range curValidators {
		prevVal := prevValidators[idx]

		// validators should be in same order if there is no change to existing validators
		if curVal.OperatorAddress != prevVal.OperatorAddress ||
			!bytes.Equal(curVal.ConsensusPubkey.Value, prevVal.ConsensusPubkey.Value) ||
			!bytes.Equal(curVal.RelayerBlsKey, prevVal.RelayerBlsKey) ||
			curVal.RelayerAddress != prevVal.RelayerAddress {
			relayercommon.Logger.Infof("syncing tendermint header at height %d", height)
			txHash, err := l.inscriptionExecutor.BscExecutor.SyncTendermintLightClientHeader(height)
			if err != nil {
				return err
			}
			relayercommon.Logger.Infof("synced tendermint header at height %d with txHash %s", height, txHash.String())
			return nil
		}
	}
	return nil
}
