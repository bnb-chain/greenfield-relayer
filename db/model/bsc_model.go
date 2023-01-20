package model

import (
	"gorm.io/gorm"

	"github.com/bnb-chain/inscription-relayer/db"
)

type BscBlock struct {
	Id         int64
	BlockHash  string `gorm:"NOT NULL"`
	ParentHash string `gorm:"NOT NULL"`
	Height     uint64 `gorm:"NOT NULL;index:idx_bsc_block_height"`
	BlockTime  int64  `gorm:"NOT NULL"`
}

func (*BscBlock) TableName() string {
	return "bsc_block"
}

type BscRelayPackage struct {
	Id              int64
	ChannelId       uint8  `gorm:"NOT NULL;index:idx_bsc_relay_package_channel_seq"`
	OracleSequence  uint64 `gorm:"NOT NULL;index:idx_bsc_relay_package_channel_seq"`
	PackageSequence uint64 `gorm:"NOT NULL"`
	PayLoad         string `gorm:"type:text"`
	TxIndex         uint   `gorm:"NOT NULL"`
	TxHash          string `gorm:"NOT NULL"`
	ClaimTxHash     string
	Status          db.TxStatus `gorm:"NOT NULL;index:idx_bsc_relay_package_status"`
	Height          uint64      `gorm:"NOT NULL;index:idx_bsc_relay_package_height"`
	TxTime          int64       `gorm:"NOT NULL"`
	UpdatedTime     int64       `gorm:"NOT NULL"`
}

func (l *BscRelayPackage) TableName() string {
	return "bsc_relay_package"
}

func InitBSCTables(db *gorm.DB) {
	if !db.Migrator().HasTable(&BscBlock{}) {
		err := db.Migrator().CreateTable(&BscBlock{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&BscRelayPackage{}) {
		err := db.Migrator().CreateTable(&BscRelayPackage{})
		if err != nil {
			panic(err)
		}
	}
}
