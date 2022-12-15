package dao

type DaoManager struct {
	InscriptionDao *InscriptionDao
	VoteDao        *VoteDao
	BSCDao         *BSCDao
}

func NewDaoManager(inscriptionDao *InscriptionDao, voteDao *VoteDao, bscDao *BSCDao) *DaoManager {
	return &DaoManager{
		InscriptionDao: inscriptionDao,
		VoteDao:        voteDao,
		BSCDao:         bscDao,
	}
}

type Result struct {
	height uint64
}
