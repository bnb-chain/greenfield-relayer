package dao

type DaoManager struct {
	GreenfieldDao *GreenfieldDao
	VoteDao       *VoteDao
	BSCDao        *BSCDao
	SequenceDao   *SequenceDao
}

func NewDaoManager(greenfieldDao *GreenfieldDao, bscDao *BSCDao, voteDao *VoteDao, seqDao *SequenceDao) *DaoManager {
	return &DaoManager{
		GreenfieldDao: greenfieldDao,
		VoteDao:       voteDao,
		BSCDao:        bscDao,
		SequenceDao:   seqDao,
	}
}
