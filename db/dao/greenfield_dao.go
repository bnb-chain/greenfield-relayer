package dao

import (
	"database/sql"
	"time"

	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/types"
)

type GreenfieldDao struct {
	DB *gorm.DB
}

func NewGreenfieldDao(db *gorm.DB) *GreenfieldDao {
	return &GreenfieldDao{
		DB: db,
	}
}

func (d *GreenfieldDao) GetLatestBlock() (*model.GreenfieldBlock, error) {
	block := model.GreenfieldBlock{}
	err := d.DB.Model(model.GreenfieldBlock{}).Order("height desc").Take(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *GreenfieldDao) GetTransactionsByStatus(s db.TxStatus) ([]*model.GreenfieldRelayTransaction, error) {
	txs := make([]*model.GreenfieldRelayTransaction, 0)
	err := d.DB.Where("status = ? ", s).Find(&txs).Order("tx_time desc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *GreenfieldDao) GetTransactionsByStatusAndHeight(status db.TxStatus, height uint64) ([]*model.GreenfieldRelayTransaction, error) {
	txs := make([]*model.GreenfieldRelayTransaction, 0)
	err := d.DB.Where("status = ? and height = ?", status, height).Find(&txs).Order("tx_time asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *GreenfieldDao) GetLeastSavedTransactionHeight() (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("greenfield_relay_transaction").Select("MIN(height)").Where("status = ?", db.Saved)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *GreenfieldDao) GetTransactionByChannelIdAndSequence(channelId types.ChannelId, sequence uint64) (*model.GreenfieldRelayTransaction, error) {
	tx := model.GreenfieldRelayTransaction{}
	err := d.DB.Where("channel_id = ? and sequence = ?", channelId, sequence).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *GreenfieldDao) GetTransactionByChannelIdAndSequenceAndStatus(channelId types.ChannelId, sequence uint64, status db.TxStatus) (*model.GreenfieldRelayTransaction, error) {
	tx := model.GreenfieldRelayTransaction{}
	err := d.DB.Where("channel_id = ? and sequence = ? and status = ?", channelId, sequence, status).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *GreenfieldDao) GetTransactionByChannelIdAndStatus(channelId types.ChannelId, status db.TxStatus) (*model.GreenfieldRelayTransaction, error) {
	tx := model.GreenfieldRelayTransaction{}
	err := d.DB.Where("channel_id = ? and status = ?", channelId, status).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *GreenfieldDao) GetLatestSequenceByChannelIdAndStatus(channelId types.ChannelId, status db.TxStatus) (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("greenfield_relay_transaction").Select("MAX(sequence)").Where("channel_id = ? and status = ?", channelId, status)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *GreenfieldDao) UpdateTransactionStatus(id int64, status db.TxStatus) error {
	err := d.DB.Model(model.GreenfieldRelayTransaction{}).Where("id = ?", id).Updates(
		model.GreenfieldRelayTransaction{Status: status, UpdatedTime: time.Now().Unix()}).Error
	return err
}

func (d *GreenfieldDao) UpdateTransactionClaimedTxHash(id int64, claimedTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.GreenfieldRelayTransaction{}).Where("id = ?", id).Updates(
			model.GreenfieldRelayTransaction{UpdatedTime: time.Now().Unix(), ClaimedTxHash: claimedTxHash}).Error
	})
}

func (d *GreenfieldDao) UpdateTransactionStatusAndClaimedTxHash(id int64, status db.TxStatus, claimedTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.GreenfieldRelayTransaction{}).Where("id = ?", id).Updates(
			model.GreenfieldRelayTransaction{Status: status, UpdatedTime: time.Now().Unix(), ClaimedTxHash: claimedTxHash}).Error
	})
}

func (d *GreenfieldDao) SaveBlockAndBatchTransactions(b *model.GreenfieldBlock, txs []*model.GreenfieldRelayTransaction) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Create(b).Error
		if err != nil {
			return err
		}

		if len(txs) != 0 {
			err := dbTx.Create(txs).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *GreenfieldDao) SaveSyncLightBlockTransaction(t *model.SyncLightBlockTransaction) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Create(t).Error
	})
}
