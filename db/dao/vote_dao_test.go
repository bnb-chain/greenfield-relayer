package dao

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestExist(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:pass@/local-greenfield-relayer0?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	assert.NoError(t, err)

	voteDao := NewVoteDao(db)
	exist, err := voteDao.IsVoteExist(0, 24, "af0de669c0ab114d6cf13e48ecd49c49afa3af01e0a82fd004feabce289dd8e7404f70398885f931d42f255f3c752b12")
	assert.NoError(t, err)
	t.Log(exist)
}
