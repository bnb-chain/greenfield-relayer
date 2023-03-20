package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/bnb-chain/greenfield-relayer/common"
	"github.com/bnb-chain/greenfield-relayer/db/model"
)

const OFFSET = 100

type SequenceDao struct {
	DB *gorm.DB
}

func NewSequenceDao(db *gorm.DB) *SequenceDao {
	return &SequenceDao{
		DB: db,
	}
}

func (d *SequenceDao) GetByChannelId(channelId uint8) (*model.Sequence, error) {
	if channelId == uint8(common.OracleChannelId) {
		channelId = channelId + OFFSET
	}
	seq := model.Sequence{}
	err := d.DB.Where("channel_id = ?", channelId).Find(&seq).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &seq, nil
}

func (d *SequenceDao) Upsert(channelId uint8, sequence uint64) error {
	if channelId == uint8(common.OracleChannelId) {
		channelId = channelId + OFFSET
	}
	seq := model.Sequence{
		ChannelId: channelId,
		Sequence:  int64(sequence),
	}
	return d.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "channel_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sequence"}),
	}).Create(&seq).Error
}
