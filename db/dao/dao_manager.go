package dao

type DaoManager struct {
	GreenfieldDao *GreenfieldDao
	VoteDao       *VoteDao
	BSCDao        *BSCDao
}

func NewDaoManager(greenfieldDao *GreenfieldDao, bscDao *BSCDao, voteDao *VoteDao) *DaoManager {
	return &DaoManager{
		GreenfieldDao: greenfieldDao,
		VoteDao:       voteDao,
		BSCDao:        bscDao,
	}
}
