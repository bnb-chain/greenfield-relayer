package listener

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"time"

	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/util"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
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
	startHeight := l.config.InscriptionConfig.StartHeight
	go l.poll(startHeight)
	go l.monitorValidators(startHeight)
}

func (l *InscriptionListener) monitorValidators(height uint64) {
	for {
		err := l.monitorValidatorsAtHeight(height)
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		}
		height++
	}
}

func (l *InscriptionListener) poll(height uint64) {
	for {
		nextHeight, err := l.pollHelper(height)
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		}
		height = nextHeight
	}
}

func (l *InscriptionListener) pollHelper(height uint64) (uint64, error) {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block from db, error: %s", err.Error())
		return 0, err
	}
	latestPolledBlockHeight := latestPolledBlock.Height
	if height <= latestPolledBlockHeight {
		height = latestPolledBlockHeight + 1
	}

	latestBlockHeight, err := l.inscriptionExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block height, error: %s", err.Error())
		return 0, err
	}

	if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
		return height, nil
	}
	blockRes, block, err := l.getBlockAndBlockResult(height)
	if err != nil {
		relayercommon.Logger.Errorf("encounter error when retrieve block and block result at height %d, err=%s", height, err.Error())
		return 0, err
	}

	err = l.monitorCrossChainEvent(blockRes, block)
	if err != nil {
		relayercommon.Logger.Errorf("encounter error when monitor events at block %d, err=%s", height, err.Error())
		return 0, err
	}
	return height + 1, nil
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
	return l.daoManager.InscriptionDao.SaveBlockAndBatchTransactions(b, txs)
}

func (l *InscriptionListener) monitorValidatorsAtHeight(height uint64) error {
	if height == 1 {
		return nil
	}

	latestHeight, err := l.inscriptionExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		return err
	}
	if height >= latestHeight {
		return nil
	}

	relayercommon.Logger.Infof("monitoring validator at height %d", height)
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
		return nil
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
