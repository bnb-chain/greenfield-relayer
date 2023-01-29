package app

import (
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/assembler"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db/dao"
	"github.com/bnb-chain/greenfield-relayer/db/model"
	"github.com/bnb-chain/greenfield-relayer/executor"
	"github.com/bnb-chain/greenfield-relayer/listener"
	"github.com/bnb-chain/greenfield-relayer/relayer"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type App struct {
	BSCRelayer        *relayer.BSCRelayer
	GreenfieldRelayer *relayer.GreenfieldRelayer
}

func NewApp(cfg *config.Config) *App {
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.DBPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	model.InitBSCTables(db)
	model.InitGreenfieldTables(db)
	model.InitVoteTables(db)

	greenfieldDao := dao.NewGreenfieldDao(db)
	bscDao := dao.NewBSCDao(db)
	voteDao := dao.NewVoteDao(db)
	daoManager := dao.NewDaoManager(greenfieldDao, bscDao, voteDao)

	greenfieldExecutor := executor.NewGreenfieldExecutor(cfg)
	bscExecutor := executor.NewBSCExecutor(cfg)

	greenfieldExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetGreenfieldExecutor(greenfieldExecutor)

	votePoolExecutor := vote.NewVotePoolExecutor(cfg)

	// listeners
	greenfieldListener := listener.NewGreenfieldListener(cfg, greenfieldExecutor, bscExecutor, daoManager)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, greenfieldExecutor, daoManager)

	// vote signer
	signer := vote.NewVoteSigner(ethcommon.Hex2Bytes(cfg.VotePoolConfig.BlsPrivateKey))

	// voteProcessors
	greenfieldVoteProcessor := vote.NewGreenfieldVoteProcessor(cfg, daoManager, signer, greenfieldExecutor, votePoolExecutor)
	bscVoteProcessor := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor, votePoolExecutor)

	// assemblers
	greenfieldAssembler := assembler.NewGreenfieldAssembler(cfg, greenfieldExecutor, daoManager, bscExecutor, votePoolExecutor)
	bscAssembler := assembler.NewBSCAssembler(cfg, bscExecutor, daoManager, votePoolExecutor, greenfieldExecutor)

	// relayers
	insRelayer := relayer.NewGreenfieldRelayer(greenfieldListener, greenfieldExecutor, bscExecutor, votePoolExecutor, greenfieldVoteProcessor, greenfieldAssembler)
	bscRelayer := relayer.NewBSCRelayer(bscListener, greenfieldExecutor, bscExecutor, votePoolExecutor, bscVoteProcessor, bscAssembler)

	return &App{
		BSCRelayer:        bscRelayer,
		GreenfieldRelayer: insRelayer,
	}
}

func (a *App) Start() {
	a.GreenfieldRelayer.Start()
	a.BSCRelayer.Start()
}
