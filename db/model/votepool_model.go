package model

import (
	"github.com/jinzhu/gorm"
)

type Vote struct {
	Id        int64
	PubKey    string
	Signature string
	EvenType  uint32
	EventHash []byte
	Sequence  uint64
	ChannelId uint8
	CreatedAt int64
}

func (*Vote) TableName() string {
	return "vote"
}

type VoteData struct {
	Id        int64
	EventHash []byte // rlp-hash of (payload, sequence, channelId)
	Sequence  uint64
	ChannelId uint8
	CreatedAt int64
}

func (*VoteData) TableName() string {
	return "vote_data"
}

func InitVoteTables(db *gorm.DB) {
	if !db.HasTable(&Vote{}) {
		db.CreateTable(&Vote{})
		db.Model(&Vote{}).AddIndex("idx_vote_channel_id_sequence", "channel_id", "sequence")
	}

	if !db.HasTable(&VoteData{}) {
		db.CreateTable(&VoteData{})
		db.Model(&VoteData{}).AddUniqueIndex("idx_vote_data_channel_id_sequence", "channel_id", "sequence")
	}
}
