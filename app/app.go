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
	"github.com/bnb-chain/greenfield-relayer/metric"
	"github.com/bnb-chain/greenfield-relayer/relayer"
	"github.com/bnb-chain/greenfield-relayer/vote"
)

type App struct {
	BSCRelayer    *relayer.BSCRelayer
	GnfdRelayer   *relayer.GreenfieldRelayer
	metricService *metric.MetricService
}

func NewApp(cfg *config.Config) *App {
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.DBPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	model.InitBSCTables(db)
	model.InitGreenfieldTables(db)
	model.InitVoteTables(db)
	model.InitSequenceTable(db)

	greenfieldDao := dao.NewGreenfieldDao(db)
	bscDao := dao.NewBSCDao(db)
	voteDao := dao.NewVoteDao(db)
	seqDao := dao.NewSequenceDao(db)
	daoManager := dao.NewDaoManager(greenfieldDao, bscDao, voteDao, seqDao)

	greenfieldExecutor := executor.NewGreenfieldExecutor(cfg)
	bscExecutor := executor.NewBSCExecutor(cfg)

	greenfieldExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetGreenfieldExecutor(greenfieldExecutor)

	metricService := metric.NewMonitorService(cfg)

	// listeners
	greenfieldListener := listener.NewGreenfieldListener(cfg, greenfieldExecutor, bscExecutor, daoManager, metricService)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, greenfieldExecutor, daoManager, metricService)

	// vote signer
	signer := vote.NewVoteSigner(ethcommon.Hex2Bytes(cfg.GreenfieldConfig.BlsPrivateKey))

	// voteProcessors
	greenfieldVoteProcessor := vote.NewGreenfieldVoteProcessor(cfg, daoManager, signer, greenfieldExecutor)
	bscVoteProcessor := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor)

	// assemblers
	greenfieldAssembler := assembler.NewGreenfieldAssembler(cfg, greenfieldExecutor, daoManager, bscExecutor, metricService)
	bscAssembler := assembler.NewBSCAssembler(cfg, bscExecutor, daoManager, greenfieldExecutor, metricService)

	// relayers
	gnfdRelayer := relayer.NewGreenfieldRelayer(greenfieldListener, greenfieldExecutor, bscExecutor, greenfieldVoteProcessor, greenfieldAssembler)
	bscRelayer := relayer.NewBSCRelayer(bscListener, greenfieldExecutor, bscExecutor, bscVoteProcessor, bscAssembler)

	return &App{
		BSCRelayer:    bscRelayer,
		GnfdRelayer:   gnfdRelayer,
		metricService: metricService,
	}
}

func (a *App) Start() {
	a.GnfdRelayer.Start()
	a.BSCRelayer.Start()
	a.metricService.Start()
}
