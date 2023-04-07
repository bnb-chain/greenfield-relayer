package dao

import (
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/db/model"
)

type VoteDao struct {
	DB *gorm.DB
}

func NewVoteDao(db *gorm.DB) *VoteDao {
	return &VoteDao{
		DB: db,
	}
}

func (d *VoteDao) GetVotesByChannelIdAndSequence(channelId uint8, sequence uint64) ([]*model.Vote, error) {
	votes := make([]*model.Vote, 0)
	err := d.DB.Where("channel_id = ? and sequence = ?", channelId, sequence).Find(&votes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return votes, nil
}

func (d *VoteDao) GetVotesCountByChannelIdAndSequence(channelId uint8, sequence uint64) (int64, error) {
	var count int64
	err := d.DB.Model(model.Vote{}).Where("channel_id = ? and sequence = ?", channelId, sequence).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *VoteDao) GetVoteByChannelIdAndSequenceAndPubKey(channelId uint8, sequence uint64, pubKey string) (*model.Vote, error) {
	vote := model.Vote{}
	err := d.DB.Model(model.Vote{}).Where("channel_id = ? and sequence = ? and pub_key = ?", channelId, sequence, pubKey).Take(&vote).Error
	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func (d *VoteDao) IsVoteExist(channelId uint8, sequence uint64, pubKey string) (bool, error) {
	exists := false
	if err := d.DB.Raw(
		"SELECT EXISTS(SELECT id FROM vote WHERE channel_id = ? and sequence = ? and pub_key = ?)",
		channelId, sequence, pubKey).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func IsVoteExist(dbTx *gorm.DB, channelId uint8, sequence uint64, pubKey string) (bool, error) {
	exists := false
	if err := dbTx.Raw(
		"SELECT EXISTS(SELECT id FROM vote WHERE channel_id = ? and sequence = ? and pub_key = ?)",
		channelId, sequence, pubKey).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (d *VoteDao) SaveVote(vote *model.Vote) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Create(vote).Error
	})
}

func SaveVote(dbTx *gorm.DB, vote *model.Vote) error {
	return dbTx.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Create(vote).Error
	})
}

func (d *VoteDao) SaveBatchVotes(votes []*model.Vote) error {
	return d.DB.Transaction(func(dbTx *gorm.DB) error {
		return dbTx.Create(votes).Error
	})
}
