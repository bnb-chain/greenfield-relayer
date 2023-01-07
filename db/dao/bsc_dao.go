package dao

import (
	"database/sql"
	relayercommon "github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"gorm.io/gorm"
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
	err := d.DB.Model(model.BscBlock{}).Order("height desc").Take(&block).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &block, nil
}

func (d *BSCDao) GetPackagesByStatus(status model.InternalStatus) ([]*model.BscRelayPackage, error) {
	votedTxs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("status = ? ", status).Find(&votedTxs).Order("tx_time asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return votedTxs, nil
}

func (d *BSCDao) GetPackagesByStatusAndHeight(status model.InternalStatus, height uint64) ([]*model.BscRelayPackage, error) {
	unVotedTxs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("status = ? and height = ?", status, height).Find(&unVotedTxs).Order("tx_time asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return unVotedTxs, nil
}

func (d *BSCDao) GetLeastSavedPackagesHeight() (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("bsc_relay_package").Select("MIN(height)").Where("status = ?", model.SAVED)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *BSCDao) UpdateBatchPackagesStatus(txIds []int64, status model.InternalStatus) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("id IN (?)", txIds).Updates(
			model.BscRelayPackage{Status: status, UpdatedTime: time.Now().Unix()}).Error
	})
}

func (d *BSCDao) UpdateBatchPackagesStatusAndClaimTxHash(txIds []int64, status model.InternalStatus, claimTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("id IN (?)", txIds).Updates(
			model.BscRelayPackage{Status: status, UpdatedTime: time.Now().Unix(), ClaimTxHash: claimTxHash}).Error
	})
}

func (d *BSCDao) GetAllVotedPackages(channelId relayercommon.ChannelId, sequence uint64) ([]*model.BscRelayPackage, error) {
	pkgs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("channel_id = ? and oracle_sequence = ? and status = ? ", channelId, sequence, model.VOTED_ALL).Find(&pkgs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return pkgs, nil
}

func (d *BSCDao) SaveBlockAndBatchPackages(b *model.BscBlock, txs []*model.BscRelayPackage) error {
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

func (d *BSCDao) DeleteBlockAtHeight(height uint64) error {
	return d.DB.Where("height = ?", height).Delete(model.BscBlock{}).Error
}

func (d *BSCDao) DeletePackagesAtHeight(height uint64) error {
	return d.DB.Where("height = ?", height).Delete(model.BscRelayPackage{}).Error
}
