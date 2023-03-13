package dao

import (
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SequenceDao struct {
	DB *gorm.DB
}

func NewSequenceDao(db *gorm.DB) *SequenceDao {
	return &SequenceDao{
		DB: db,
	}
}

func (d *SequenceDao) GetByChannelId(channelId uint8) (*model.Sequence, error) {
	seq := model.Sequence{}
	err := d.DB.Where("channel_id = ?", channelId).Find(&seq).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &seq, nil
}

func (d *SequenceDao) Upsert(channelId uint8, sequence uint64) error {
	seq := model.Sequence{
		ChannelId: channelId,
		Sequence:  int64(sequence),
	}
	return d.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "channel_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sequence"}),
	}).Create(&seq).Error
}
