package listener

import (
	"bytes"
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
	DaoManager          *dao.DaoManager
}

func NewInscriptionListener(cfg *config.Config, executor *executor.InscriptionExecutor, dao *dao.DaoManager) *InscriptionListener {
	return &InscriptionListener{
		config:              cfg,
		inscriptionExecutor: executor,
		DaoManager:          dao,
	}
}

func (l *InscriptionListener) Start() {
	heightToPoll := l.config.InscriptionConfig.StartHeight
	go l.poll(heightToPoll)
	go l.monitorValidators(heightToPoll)
}

func (l *InscriptionListener) monitorValidators(heightToPoll uint64) {
	for {
		nextHeight, err := l.monitorValidatorsAtHeight(heightToPoll)
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		}
		heightToPoll = nextHeight
	}
}

func (l *InscriptionListener) poll(heightToPoll uint64) {
	for {
		nextHeight, err := l.pollHelper(heightToPoll)
		if err != nil {
			time.Sleep(RetryInterval)
			continue
		}
		heightToPoll = nextHeight
	}
}

func (l *InscriptionListener) pollHelper(heightToPoll uint64) (uint64, error) {
	latestPolledBlock, err := l.getLatestPolledBlock()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block from db, error: %s", err.Error())
		return 0, err
	}
	latestPolledBlockHeight := latestPolledBlock.Height
	if heightToPoll <= latestPolledBlockHeight {
		heightToPoll = latestPolledBlockHeight + 1
	}

	latestBlockHeight, err := l.inscriptionExecutor.GetLatestBlockHeightWithRetry()
	if err != nil {
		relayercommon.Logger.Errorf("failed to get latest block heightToPoll, error: %s", err.Error())
		return 0, err
	}
	if int64(latestPolledBlockHeight) >= int64(latestBlockHeight)-1 {
		return heightToPoll, nil
	}
	blockRes, block, err := l.getBlockAndBlockResult(heightToPoll)
	if err != nil {
		relayercommon.Logger.Errorf("encounter error when retrieve block and block result at heightToPoll %d, err=%s", heightToPoll, err.Error())
		return 0, err
	}

	err = l.monitorCrossChainEvent(blockRes, block)
	if err != nil {
		relayercommon.Logger.Errorf("encounter error when monitor events at block %d, err=%s", heightToPoll, err.Error())
		return 0, err
	}
	return heightToPoll + 1, nil
}

func (l *InscriptionListener) getLatestPolledBlock() (*model.InscriptionBlock, error) {
	block, err := l.DaoManager.InscriptionDao.GetLatestBlock()
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
			// event.ma
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
						relayercommon.Logger.Errorf("unexpected attr, key is %s", attr.Key)
					}
				}

				//TODO for tesitng
				//relayTx.Sequence =
				nextDeliverySeqOnInscription, _ := l.inscriptionExecutor.GetNextDeliverySequenceForChannel(relayercommon.ChannelId(relayTx.ChannelId))
				relayTx.Sequence = nextDeliverySeqOnInscription

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

func (l *InscriptionListener) monitorValidatorsAtHeight(height uint64) (uint64, error) {
	if height == 1 {
		return height + 1, nil
	}

	lightClientLatestHeight, err := l.inscriptionExecutor.BscExecutor.GetLightClientLatestHeight()
	if err != nil {
		return height, err
	}

	if height <= lightClientLatestHeight {
		return lightClientLatestHeight + 1, nil
	}

	relayercommon.Logger.Infof("monitoring validator at height %d", height)
	curValidators, err := l.inscriptionExecutor.QueryValidatorsAtHeight(height)
	if err != nil {
		return height, err
	}

	prevValidators, err := l.inscriptionExecutor.QueryValidatorsAtHeight(height - 1)
	if err != nil {
		return height, err
	}

	if len(curValidators) != len(prevValidators) {
		txHash, err := l.inscriptionExecutor.BscExecutor.SyncTendermintLightClientHeader(height)
		if err != nil {
			return height, err
		}
		relayercommon.Logger.Infof("synced tendermint header at height %d with txHash %s", height, txHash.String())
		return height + 1, nil
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
				return height, err
			}
			relayercommon.Logger.Infof("synced tendermint header at height %d with txHash %s", height, txHash.String())
			return height + 1, nil
		}
	}
	return height + 1, nil
}
