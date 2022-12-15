package dao

import (
	"github.com/jinzhu/gorm"
	relayercommon "inscription-relayer/common"
	"inscription-relayer/db/model"
	"inscription-relayer/vote"
	"time"
)

type BSCDao struct {
	DB *gorm.DB
}

func NewBSCDao(db *gorm.DB) *BSCDao {
	return &BSCDao{
		DB: db,
	}
}

func (d *BSCDao) GetLatestBlock() (*model.BscBlock, error) {
	block := model.BscBlock{}
	err := d.DB.Model(model.BscBlock{}).Order("created_at desc").First(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *BSCDao) GetUnSignedPackages() ([]*model.BscRelayPackage, error) {
	unSignedTxs := make([]*model.BscRelayPackage, 0)

	err := d.DB.Where("status = ?", vote.SAVED).Find(&unSignedTxs).Order("created_at asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return unSignedTxs, nil
}

func (d *BSCDao) GetUnVotedPackagesAtHeight(height uint64) ([]*model.BscRelayPackage, error) {
	unVotedTxs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("status = ? and height = ", vote.SAVED, height).Find(&unVotedTxs).Order("created_at asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return unVotedTxs, nil
}

func (d *BSCDao) GetLatestVotedPackagesHeight() (uint64, error) {
	var result uint64
	row := d.DB.Table("bsc_relay_package").Select("MAX(height)").Where("status = ?", vote.VOTED).Row()
	err := row.Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (d *BSCDao) UpdateBatchPackagesStatus(txIds []int64, status vote.InternalStatus) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	err := d.DB.Model(model.BscRelayPackage{}).Where("id IN ?", txIds).Updates(
		map[string]interface{}{
			"status":     status,
			"updated_at": time.Now().Unix(),
		}).Error
	if err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

func (d *BSCDao) GetVotedPackages() ([]*model.BscRelayPackage, error) {
	votedTxs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("status = ? ", vote.VOTED).Find(&votedTxs).Order("created_at desc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return votedTxs, nil
}

func (d *BSCDao) GetAllVotedPackages(channelId relayercommon.ChannelId, sequence uint64) ([]*model.BscRelayPackage, error) {
	pkgs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("channel_id = ? and oracle_sequence = ? and status = ? ", channelId, sequence, vote.VOTED_ALL).Find(&pkgs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return pkgs, nil
}

func (d *BSCDao) SaveBlockAndBatchPackages(b *model.BscBlock, txs []*model.BscRelayPackage) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	if err := dbTx.Create(txs).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	if err := dbTx.Create(b).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

func (d *BSCDao) SaveBatchPackages(txs []*model.BscRelayPackage) error {
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

func (d *BSCDao) DeleteBlockAtHeight(height uint64) error {
	return d.DB.Where("height = ?", height).Delete(model.BscBlock{}).Error
}

func (d *BSCDao) DeletePackagesAtHeight(height uint64) error {
	return d.DB.Where("height = ?", height).Delete(model.BscRelayPackage{}).Error
}
