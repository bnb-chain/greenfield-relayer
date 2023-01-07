package dao

import (
	"database/sql"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"gorm.io/gorm"
	"time"
)

type InscriptionDao struct {
	DB *gorm.DB
}

func NewInscriptionDao(db *gorm.DB) *InscriptionDao {
	return &InscriptionDao{
		DB: db,
	}
}

func (d *InscriptionDao) GetLatestBlock() (*model.InscriptionBlock, error) {
	block := model.InscriptionBlock{}
	err := d.DB.Model(model.InscriptionBlock{}).Order("height desc").Take(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *InscriptionDao) GetTransactionsByStatus(s model.InternalStatus) ([]*model.InscriptionRelayTransaction, error) {
	txs := make([]*model.InscriptionRelayTransaction, 0)
	err := d.DB.Where("status = ? ", s).Find(&txs).Order("tx_time desc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *InscriptionDao) GetTransactionsByStatusAndHeight(status model.InternalStatus, height uint64) ([]*model.InscriptionRelayTransaction, error) {
	txs := make([]*model.InscriptionRelayTransaction, 0)
	err := d.DB.Where("status = ? and height = ?", status, height).Find(&txs).Order("tx_time asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *InscriptionDao) GetLatestVotedTxHeight() (uint64, error) {
	var result uint64
	res := d.DB.Table("inscription_relay_transaction").Select("MAX(height)").Where("status = ?", model.VOTED)
	if res.RowsAffected == 0 {
		return 0, nil
	}
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (d *InscriptionDao) GetLeastSavedTxHeight() (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("inscription_relay_transaction").Select("MIN(height)").Where("status = ?", model.SAVED)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *InscriptionDao) UpdateTxStatus(id int64, status model.InternalStatus) error {
	err := d.DB.Model(model.InscriptionRelayTransaction{}).Where("id = ?", id).Updates(
		model.InscriptionRelayTransaction{Status: status, UpdatedTime: time.Now().Unix()}).Error
	return err
}

func (d *InscriptionDao) UpdateTxStatusAndClaimTxHash(id int64, status model.InternalStatus, claimTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.InscriptionRelayTransaction{}).Where("id = ?", id).Updates(
			model.InscriptionRelayTransaction{Status: status, UpdatedTime: time.Now().Unix(), ClaimTxHash: claimTxHash}).Error
	})
}

func (d *InscriptionDao) GetTransactionByChannelIdAndSequence(channelId relayercommon.ChannelId, sequence uint64) (*model.InscriptionRelayTransaction, error) {
	tx := model.InscriptionRelayTransaction{}
	err := d.DB.Where("channel_id = ? and sequence = ?", channelId, sequence).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *InscriptionDao) SaveBatchTransaction(tx *model.InscriptionRelayTransaction) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Create(tx).Error
	})
}

func (d *InscriptionDao) SaveBlockAndBatchTransactions(b *model.InscriptionBlock, txs []*model.InscriptionRelayTransaction) error {
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
