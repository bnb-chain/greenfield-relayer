package model

import "gorm.io/gorm"

type Sequence struct {
	ChannelId uint8 `gorm:"primaryKey"`
	Sequence  int64 `gorm:"NOT NULL"`
}

func InitSequenceTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&Sequence{}) {
		err := db.Migrator().CreateTable(&Sequence{})
		if err != nil {
			panic(err)
		}
	}
}
