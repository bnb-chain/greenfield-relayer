package dao

type DaoManager struct {
	InscriptionDao *InscriptionDao
	VoteDao        *VoteDao
	BSCDao         *BSCDao
}

func NewDaoManager(inscriptionDao *InscriptionDao, bscDao *BSCDao, voteDao *VoteDao) *DaoManager {
	return &DaoManager{
		InscriptionDao: inscriptionDao,
		VoteDao:        voteDao,
		BSCDao:         bscDao,
	}
}
