package model

import (
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/db"
)

type GreenfieldBlock struct {
	Id        int64
	Chain     string
	Height    uint64 `gorm:"NOT NULL;index:idx_greenfield_block_height"`
	BlockTime int64  `gorm:"NOT NULL"`
}

func (*GreenfieldBlock) TableName() string {
	return "greenfield_block"
}

type GreenfieldRelayTransaction struct {
	Id            int64
	SrcChainId    uint32 `gorm:"NOT NULL"`
	DestChainId   uint32 `gorm:"NOT NULL"`
	ChannelId     uint8  `gorm:"NOT NULL;index:idx_greenfield_relay_transaction_channel_seq_status"`
	Sequence      uint64 `gorm:"NOT NULL;index:idx_greenfield_relay_transaction_channel_seq_status"`
	PackageType   uint32 `gorm:"NOT NULL"`
	Height        uint64 `gorm:"NOT NULL;index:idx_greenfield_relay_transaction_height_status"`
	PayLoad       string `gorm:"type:text"`
	RelayerFee    string `gorm:"NOT NULL"`
	AckRelayerFee string `gorm:"NOT NULL"`
	ClaimedTxHash string
	Status        db.TxStatus `gorm:"NOT NULL;index:idx_greenfield_relay_transaction_channel_seq_status;idx_greenfield_relay_transaction_height_status"`
	TxTime        int64       `gorm:"NOT NULL"`
	UpdatedTime   int64       `gorm:"NOT NULL"`
}

func (*GreenfieldRelayTransaction) TableName() string {
	return "greenfield_relay_transaction"
}

func InitGreenfieldTables(db *gorm.DB) {
	if !db.Migrator().HasTable(&GreenfieldBlock{}) {
		err := db.Migrator().CreateTable(&GreenfieldBlock{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&GreenfieldRelayTransaction{}) {
		err := db.Migrator().CreateTable(&GreenfieldRelayTransaction{})
		if err != nil {
			panic(err)
		}
	}
}
