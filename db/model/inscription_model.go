package model

import (
	"gorm.io/gorm"

	"github.com/bnb-chain/inscription-relayer/db"
)

type InscriptionBlock struct {
	Id        int64
	Chain     string
	Height    uint64 `gorm:"NOT NULL;index:idx_inscription_block_height"`
	BlockTime int64  `gorm:"NOT NULL"`
}

func (*InscriptionBlock) TableName() string {
	return "inscription_block"
}

type InscriptionRelayTransaction struct {
	Id            int64
	SrcChainId    uint32 `gorm:"NOT NULL"`
	DestChainId   uint32 `gorm:"NOT NULL"`
	ChannelId     uint8  `gorm:"NOT NULL;index:idx_inscription_relay_transaction_channel_seq_status"`
	Sequence      uint64 `gorm:"NOT NULL;index:idx_inscription_relay_transaction_channel_seq_status"`
	PackageType   uint32 `gorm:"NOT NULL"`
	Height        uint64 `gorm:"NOT NULL;index:idx_inscription_relay_transaction_height"`
	PayLoad       string `gorm:"type:text"`
	RelayerFee    string `gorm:"NOT NULL"`
	AckRelayerFee string `gorm:"NOT NULL"`
	ClaimedTxHash string
	Status        db.TxStatus `gorm:"NOT NULL;index:idx_inscription_relay_transaction_status"`
	TxTime        int64       `gorm:"NOT NULL"`
	UpdatedTime   int64       `gorm:"NOT NULL"`
}

func (*InscriptionRelayTransaction) TableName() string {
	return "inscription_relay_transaction"
}

func InitInscriptionTables(db *gorm.DB) {
	if !db.Migrator().HasTable(&InscriptionBlock{}) {
		err := db.Migrator().CreateTable(&InscriptionBlock{})
		if err != nil {
			panic(err)
		}
	}

	if !db.Migrator().HasTable(&InscriptionRelayTransaction{}) {
		err := db.Migrator().CreateTable(&InscriptionRelayTransaction{})
		if err != nil {
			panic(err)
		}
	}
}
