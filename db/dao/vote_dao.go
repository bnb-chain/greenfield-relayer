package dao

import (
	"github.com/jinzhu/gorm"
	"inscription-relayer/db/model"
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

func (d *VoteDao) GetVoteDataByChannelAndSequence(channelId uint8, sequence uint64) (*model.VoteData, error) {
	voteData := model.VoteData{}
	err := d.DB.Model(model.VoteData{}).Where("channel_id = ? and sequence = ?", channelId, sequence).Find(&voteData).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &voteData, nil
}

func (d *VoteDao) SaveVote(vote *model.Vote) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	if err := dbTx.Create(vote).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

func (d *VoteDao) SaveVoteData(voteData *model.VoteData) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	if err := dbTx.Create(voteData).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}

func (d *VoteDao) SaveBatchVotes(votes []*model.Vote) error {
	dbTx := d.DB.Begin()
	if err := dbTx.Error; err != nil {
		return err
	}
	if err := dbTx.Create(votes).Error; err != nil {
		dbTx.Rollback()
		return err
	}
	return dbTx.Commit().Error
}
