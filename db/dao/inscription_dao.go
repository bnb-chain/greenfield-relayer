package dao

import (
	"github.com/jinzhu/gorm"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/db/model"
	"inscription-relayer/vote"
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
	err := d.DB.Model(model.InscriptionBlock{}).Order("created_at desc").First(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *InscriptionDao) SaveBlock(b *model.InscriptionBlock) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	if err := dbTx.Create(b).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

func (d *InscriptionDao) GetUnSignedTransactions() ([]*model.InscriptionRelayTransaction, error) {
	unSignedTxs := make([]*model.InscriptionRelayTransaction, 0)

	err := d.DB.Where("status = ?", vote.SAVED).Find(&unSignedTxs).Order("created_at asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return unSignedTxs, nil
}

func (d *InscriptionDao) GetUnVotedTransactionsAtHeight(height uint64) ([]*model.InscriptionRelayTransaction, error) {
	unVotedTxs := make([]*model.InscriptionRelayTransaction, 0)

	err := d.DB.Where("status = ? and height = ", vote.SAVED, height).Find(&unVotedTxs).Order("created_at asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return unVotedTxs, nil
}

func (d *InscriptionDao) GetLatestVotedTransactionHeight() (uint64, error) {
	var result uint64
	row := d.DB.Table("inscription_relay_transaction").Select("MAX(height)").Where("status = ?", vote.VOTED).Row()
	err := row.Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (d *InscriptionDao) UpdateTransactionStatus(txId int64, status vote.InternalStatus) error {
	err := d.DB.Model(model.InscriptionRelayTransaction{}).Where("id = ?", txId).Updates(
		map[string]interface{}{
			"status":     status,
			"updated_at": time.Now().Unix(),
		}).Error
	return err
}

func (d *InscriptionDao) GetVotedTransactions() ([]*model.InscriptionRelayTransaction, error) {
	votedTxs := make([]*model.InscriptionRelayTransaction, 0)

	err := d.DB.Where("status = ? ", vote.VOTED).Find(&votedTxs).Order("created_at desc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return votedTxs, nil
}

func (d *InscriptionDao) GetTransactionByChannelIdAndSequenceWithStatusAllVoted(channelId relayercommon.ChannelId, sequence uint64) (*model.InscriptionRelayTransaction, error) {
	tx := model.InscriptionRelayTransaction{}
	err := d.DB.Where("channel_id = ? and sequence = ? and ", channelId, sequence).Find(&tx).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tx, nil
}

func (d *InscriptionDao) SaveBatchTransactions(txs []*model.InscriptionRelayTransaction) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	if err := dbTx.Create(txs).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}
