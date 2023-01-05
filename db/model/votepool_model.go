package model

import (
	"github.com/jinzhu/gorm"
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
	PubKey      string
	Signature   string
	EventType   uint32
	EventHash   []byte
	Sequence    uint64
	ChannelId   uint8
	CreatedTime int64
}

func (*Vote) TableName() string {
	return "vote"
}

func InitVoteTables(db *gorm.DB) {
	if !db.HasTable(&Vote{}) {
		db.CreateTable(&Vote{})
		db.Model(&Vote{}).AddIndex("idx_vote_channel_id_sequence_pub_key", "channel_id", "sequence", "pub_key")
	}
}
