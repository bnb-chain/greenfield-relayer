package model

import (
	"gorm.io/gorm"
)

type Vote struct {
	Id           int64
	Signature    string `gorm:"NOT NULL"`
	EventType    uint32 `gorm:"NOT NULL"`
	ClaimPayload []byte `gorm:"NOT NULL"`
	EventHash    []byte `gorm:"NOT NULL"`
	Sequence     uint64 `gorm:"NOT NULL;index:idx_vote_channel_id_sequence_pub_key"`
	ChannelId    uint8  `gorm:"NOT NULL;index:idx_vote_channel_id_sequence_pub_key"`
	PubKey       string `gorm:"NOT NULL;index:idx_vote_channel_id_sequence_pub_key"`
	CreatedTime  int64  `gorm:"NOT NULL"`
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
