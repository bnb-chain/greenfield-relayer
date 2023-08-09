package dao

import (
	"github.com/bnb-chain/greenfield-relayer/logging"
	"github.com/cometbft/cometbft/votepool"
	"time"
)

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

func (d *DaoManager) PurgeLoop() {
	numOfHistroicalBlocks := int64(100000)
	ticker := time.NewTicker(time.Minute * 10)
	for range ticker.C {
		latestBscBlock, err := d.BSCDao.GetLatestBlock()
		if err != nil {
			logging.Logger.Errorf("failed to get latest DB BSC block, err=%s", err.Error())
			continue
		}
		threshHold := int64(latestBscBlock.Height) - numOfHistroicalBlocks
		if threshHold > 0 {
			err = d.BSCDao.DeleteBlocks(threshHold)
			if err != nil {
				logging.Logger.Errorf("failed to delete Bsc blocks, err=%s", err.Error())
				continue
			}
			err = d.BSCDao.DeletePackages(threshHold)
			if err != nil {
				logging.Logger.Errorf("failed to delete bsc packages, err=%s", err.Error())
			}
			err = d.VoteDao.DeleteVotes(threshHold, uint32(votepool.FromBscCrossChainEvent))
			if err != nil {
				logging.Logger.Errorf("failed to delete votes, err=%s", err.Error())
			}
		}
		latestGnfdBlock, err := d.GreenfieldDao.GetLatestBlock()
		if err != nil {
			logging.Logger.Errorf("failed to get latest DB BSC block, err=%s", err.Error())
			continue
		}
		threshHold = int64(latestGnfdBlock.Height) - numOfHistroicalBlocks
		if threshHold > 0 {
			err = d.GreenfieldDao.DeleteBlocks(threshHold)
			if err != nil {
				logging.Logger.Errorf("failed to delete gnfd blocks, err=%s", err.Error())
				continue
			}
			err = d.GreenfieldDao.DeleteTransactions(threshHold)
			if err != nil {
				logging.Logger.Errorf("failed to delete gnfd transactions, err=%s", err.Error())
				continue
			}
			err = d.VoteDao.DeleteVotes(threshHold, uint32(votepool.ToBscCrossChainEvent))
			if err != nil {
				logging.Logger.Errorf("failed to delete votes, err=%s", err.Error())
			}
		}
	}
}
