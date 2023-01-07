package model

import (
	"gorm.io/gorm"
)

type InternalStatus int

const (
	SAVED     InternalStatus = 0
	VOTED     InternalStatus = 1
	VOTED_ALL InternalStatus = 2
	FILLED    InternalStatus = 3
)

type Vote struct {
	Id          int64
	PubKey      string `gorm:"NOT NULL;index:idx_vote_channel_id_sequence_pub_key"`
	Signature   string `gorm:"NOT NULL"`
	EventType   uint32 `gorm:"NOT NULL"`
	EventHash   []byte `gorm:"NOT NULL"`
	Sequence    uint64 `gorm:"NOT NULL;index:idx_vote_channel_id_sequence_pub_key"`
	ChannelId   uint8  `gorm:"NOT NULL;index:idx_vote_channel_id_sequence_pub_key"`
	CreatedTime int64  `gorm:"NOT NULL"`
}

func (*Vote) TableName() string {
	return "vote"
}

func InitVoteTables(db *gorm.DB) {
	if !db.Migrator().HasTable(&Vote{}) {
		err := db.Migrator().CreateTable(&Vote{})
		if err != nil {
			panic(err)
		}
	}
}
