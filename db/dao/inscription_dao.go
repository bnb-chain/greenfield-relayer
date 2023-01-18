package dao

import (
	"database/sql"
	"time"

	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/db"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"gorm.io/gorm"
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

func (d *InscriptionDao) GetTransactionsByStatus(s db.TxStatus) ([]*model.InscriptionRelayTransaction, error) {
	txs := make([]*model.InscriptionRelayTransaction, 0)
	err := d.DB.Where("status = ? ", s).Find(&txs).Order("tx_time desc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *InscriptionDao) GetTransactionsByStatusAndHeight(status db.TxStatus, height uint64) ([]*model.InscriptionRelayTransaction, error) {
	txs := make([]*model.InscriptionRelayTransaction, 0)
	err := d.DB.Where("status = ? and height = ?", status, height).Find(&txs).Order("tx_time asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return txs, nil
}

func (d *InscriptionDao) GetLatestVotedTransactionHeight() (uint64, error) {
	var result uint64
	res := d.DB.Table("inscription_relay_transaction").Select("MAX(height)").Where("status = ?", db.SelfVoted)
	if res.RowsAffected == 0 {
		return 0, nil
	}
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (d *InscriptionDao) GetLeastSavedTransactionHeight() (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("inscription_relay_transaction").Select("MIN(height)").Where("status = ?", db.Saved)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *InscriptionDao) GetTransactionByChannelIdAndSequenceAndStatus(channelId relayercommon.ChannelId, sequence uint64, status db.TxStatus) (*model.InscriptionRelayTransaction, error) {
	tx := model.InscriptionRelayTransaction{}
	err := d.DB.Where("channel_id = ? and sequence = ? and status = ?", channelId, sequence, status).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *InscriptionDao) UpdateTransactionStatus(id int64, status db.TxStatus) error {
	err := d.DB.Model(model.InscriptionRelayTransaction{}).Where("id = ?", id).Updates(
		model.InscriptionRelayTransaction{Status: status, UpdatedTime: time.Now().Unix()}).Error
	return err
}

func (d *InscriptionDao) UpdateTransactionStatusAndClaimedTxHash(id int64, status db.TxStatus, claimedTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.InscriptionRelayTransaction{}).Where("id = ?", id).Updates(
			model.InscriptionRelayTransaction{Status: status, UpdatedTime: time.Now().Unix(), ClaimedTxHash: claimedTxHash}).Error
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
