package dao

import (
	"database/sql"
	"time"

	"github.com/cometbft/cometbft/votepool"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/db"
	"github.com/bnb-chain/greenfield-relayer/db/model"
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

func (d *BSCDao) GetPackagesByStatus(status db.TxStatus) ([]*model.BscRelayPackage, error) {
	votedTxs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("status = ? ", status).Find(&votedTxs).Order("tx_time asc").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return votedTxs, nil
}

func (d *BSCDao) GetPackagesByHeightAndStatus(status db.TxStatus, height uint64) ([]*model.BscRelayPackage, error) {
	unVotedTxs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("status = ? and height = ?", status, height).Order("height asc").Find(&unVotedTxs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return unVotedTxs, nil
}

func (d *BSCDao) GetLeastSavedPackagesHeight() (uint64, error) {
	var result sql.NullInt64
	res := d.DB.Table("bsc_relay_package").Select("MIN(height)").Where("status = ?", db.Saved)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	return uint64(result.Int64), nil
}

func (d *BSCDao) GetLatestOracleSequenceByStatus(status db.TxStatus) (int64, error) {
	var result sql.NullInt64
	res := d.DB.Table("bsc_relay_package").Select("MAX(oracle_sequence)").Where("status = ?", status)
	err := res.Row().Scan(&result)
	if err != nil {
		return 0, err
	}
	if !result.Valid {
		return -1, nil
	}
	return result.Int64, nil
}

func (d *BSCDao) GetPackagesByOracleSequence(sequence uint64) ([]*model.BscRelayPackage, error) {
	pkgs := make([]*model.BscRelayPackage, 0)
	err := d.DB.Where("oracle_sequence = ?", sequence).Find(&pkgs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return pkgs, nil
}

func (d *BSCDao) UpdateBatchPackagesStatus(txIds []int64, status db.TxStatus) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("id IN (?)", txIds).Updates(
			model.BscRelayPackage{Status: status, UpdatedTime: time.Now().Unix()}).Error
	})
}

func UpdateBatchPackagesStatus(dbTx *gorm.DB, txIds []int64, status db.TxStatus) error {
	return dbTx.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("id IN (?)", txIds).Updates(
			model.BscRelayPackage{Status: status, UpdatedTime: time.Now().Unix()}).Error
	})
}

func (d *BSCDao) UpdateBatchPackagesStatusToDelivered(seq uint64) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("oracle_sequence < ? and status = 2", seq).Updates(
			model.BscRelayPackage{Status: db.Delivered, UpdatedTime: time.Now().Unix()}).Error
	})
}

func (d *BSCDao) UpdateBatchPackagesClaimedTxHash(txIds []int64, claimTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("id IN (?)", txIds).Updates(
			model.BscRelayPackage{UpdatedTime: time.Now().Unix(), ClaimTxHash: claimTxHash}).Error
	})
}

func (d *BSCDao) UpdateBatchPackagesStatusAndClaimedTxHash(txIds []int64, status db.TxStatus, claimTxHash string) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Model(model.BscRelayPackage{}).Where("id IN (?)", txIds).Updates(
			model.BscRelayPackage{Status: status, UpdatedTime: time.Now().Unix(), ClaimTxHash: claimTxHash}).Error
	})
}

func (d *BSCDao) SaveBlockAndBatchPackages(b *model.BscBlock, pkgs []*model.BscRelayPackage) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Create(b).Error
		if err != nil {
			return err
		}
		if len(pkgs) != 0 {
			err := dbTx.Create(pkgs).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *BSCDao) SaveBatchPackages(pkgs []*model.BscRelayPackage) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(pkgs) != 0 {
			err := dbTx.Create(pkgs).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *BSCDao) DeleteBlockAndPackagesAndVotesAtHeight(height uint64) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Where("height = ?", height).Delete(model.BscBlock{}).Error
		if err != nil {
			return err
		}
		err = dbTx.Where("height = ?", height).Delete(model.BscRelayPackage{}).Error
		if err != nil {
			return err
		}
		err = dbTx.Where("height = ? and event_type = ?", height, votepool.FromOpCrossChainEvent).Delete(model.Vote{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *BSCDao) DeleteBlocksBelowHeight(threshHold int64) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Where("height < ?", threshHold).Delete(model.BscBlock{}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *BSCDao) DeletePackagesBelowHeightWithLimit(threshHold int64, limit int) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		err := dbTx.Where("height < ?", threshHold).Delete(model.BscRelayPackage{}).Limit(limit).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *BSCDao) ExistsUnprocessedPackage(threshHold int64) (bool, error) {
	tx := model.BscRelayPackage{}
	err := d.DB.Model(model.BscRelayPackage{}).Where("status = ? or status = ? and height < ?", db.Saved, db.SelfVoted, threshHold).Take(&tx).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
