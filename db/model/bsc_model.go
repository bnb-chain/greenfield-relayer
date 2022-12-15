package model

import (
	"github.com/jinzhu/gorm"
	"inscription-relayer/vote"
)

type BscBlock struct {
	Id         int64
	BlockHash  string `gorm:"NOT NULL"`
	ParentHash string `gorm:"NOT NULL"`
	Height     uint64 `gorm:"NOT NULL"`
	BlockTime  int64  `gorm:"NOT NULL"`
}

func (*BscBlock) TableName() string {
	return "bsc_block"
}

type BscRelayPackage struct {
	Id              int64
	TxHash          string `gorm:"NOT NULL"`
	ChainId         uint16
	ChannelId       uint8  `gorm:"NOT NULL"`
	OracleSequence  uint64 `gorm:"NOT NULL"`
	PackageSequence uint64 `gorm:"NOT NULL"`
	PayLoad         string `gorm:"type:text"`
	TxIndex         uint
	Status          vote.InternalStatus `gorm:"NOT NULL"`
	BlockHash       string              `gorm:"NOT NULL"`
	Height          uint64              `gorm:"NOT NULL"`
	CreatedAt       int64               `gorm:"NOT NULL"`
	UpdatedAt       int64
}

func (l *BscRelayPackage) TableName() string {
	return "bsc_relay_package"
}

func InitBSCTables(db *gorm.DB) {
	if !db.HasTable(&BscBlock{}) {
		db.CreateTable(&BscBlock{})
		db.Model(&BscBlock{}).AddUniqueIndex("idx_bsc_block_height", "height")
		db.Model(&BscBlock{}).AddIndex("idx_bsc_block_created_at", "block_time")
	}

	if !db.HasTable(&BscRelayPackage{}) {
		db.CreateTable(&BscRelayPackage{})
		db.Model(&BscRelayPackage{}).AddIndex("idx_bsc_relay_package_channel_seq", "channel_id", "oracle_sequence")
		db.Model(&BscRelayPackage{}).AddIndex("idx_bsc_relay_package_height", "height")
	}
}
