package model

import (
	"github.com/jinzhu/gorm"
)

type InscriptionBlock struct {
	Id        int64
	Chain     string
	BlockHash string `gorm:"NOT NULL"`
	Height    uint64 `gorm:"NOT NULL"`
	BlockTime int64  `gorm:"NOT NULL"`
}

func (*InscriptionBlock) TableName() string {
	return "inscription_block"
}

type InscriptionRelayTransaction struct {
	Id        int64
	TxHash    string `gorm:"NOT NULL"`
	Type      string `gorm:"NOT NULL"`
	TxStatus  string `gorm:"NOT NULL"`
	ChannelId uint8  `gorm:"NOT NULL"`
	Sequence  uint64 `gorm:"NOT NULL"`
	Height    uint64 `gorm:"NOT NULL"`

	PayLoad string `gorm:"type:text"`

	TxGasPrice uint64
	TxGasLimit uint64
	TxUsedGas  uint64
	TxFee      uint64

	Status    int   `gorm:"NOT NULL"`
	CreatedAt int64 `gorm:"NOT NULL"`
	UpdatedAt int64
}

func (*InscriptionRelayTransaction) TableName() string {
	return "inscription_relay_transaction"
}

func InitInscriptionTables(db *gorm.DB) {
	if !db.HasTable(&InscriptionBlock{}) {
		db.CreateTable(&InscriptionBlock{})
		db.Model(&InscriptionBlock{}).AddUniqueIndex("idx_inscription_block_height", "height")
		db.Model(&InscriptionBlock{}).AddIndex("idx_inscription_block_block_time", "block_time")
	}

	if !db.HasTable(&InscriptionRelayTransaction{}) {
		db.CreateTable(&InscriptionRelayTransaction{})
		db.Model(&InscriptionRelayTransaction{}).AddUniqueIndex("idx_inscription_relay_transaction_channel_seq_status", "channel_id", "sequence", "status")
		db.Model(&InscriptionRelayTransaction{}).AddIndex("idx_inscription_relay_transaction_height", "height")
	}
}
