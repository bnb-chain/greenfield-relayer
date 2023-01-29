package app

import (
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bnb-chain/inscription-relayer/assembler"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/listener"
	"github.com/bnb-chain/inscription-relayer/relayer"
	"github.com/bnb-chain/inscription-relayer/vote"
)

type App struct {
	BSCRelayer         *relayer.BSCRelayer
	InscriptionRelayer *relayer.InscriptionRelayer
}

func NewApp(cfg *config.Config) *App {
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.DBPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	model.InitBSCTables(db)
	model.InitInscriptionTables(db)
	model.InitVoteTables(db)

	inscriptionDao := dao.NewInscriptionDao(db)
	bscDao := dao.NewBSCDao(db)
	voteDao := dao.NewVoteDao(db)
	daoManager := dao.NewDaoManager(inscriptionDao, bscDao, voteDao)

	inscriptionExecutor := executor.NewInscriptionExecutor(cfg)
	bscExecutor := executor.NewBSCExecutor(cfg)

	inscriptionExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetInscriptionExecutor(inscriptionExecutor)

	votePoolExecutor := vote.NewVotePoolExecutor(cfg)

	// listeners
	inscriptionListener := listener.NewInscriptionListener(cfg, inscriptionExecutor, bscExecutor, daoManager)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, inscriptionExecutor, daoManager)

	// vote signer
	signer := vote.NewVoteSigner(ethcommon.Hex2Bytes(cfg.VotePoolConfig.BlsPrivateKey))

	// voteProcessors
	inscriptionVoteProcessor := vote.NewInscriptionVoteProcessor(cfg, daoManager, signer, inscriptionExecutor, votePoolExecutor)
	bscVoteProcessor := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor, votePoolExecutor)

	// assemblers
	inscriptionAssembler := assembler.NewInscriptionAssembler(cfg, inscriptionExecutor, daoManager, bscExecutor, votePoolExecutor)
	bscAssembler := assembler.NewBSCAssembler(cfg, bscExecutor, daoManager, votePoolExecutor, inscriptionExecutor)

	// relayers
	insRelayer := relayer.NewInscriptionRelayer(inscriptionListener, inscriptionExecutor, bscExecutor, votePoolExecutor, inscriptionVoteProcessor, inscriptionAssembler)
	bscRelayer := relayer.NewBSCRelayer(bscListener, inscriptionExecutor, bscExecutor, votePoolExecutor, bscVoteProcessor, bscAssembler)

	return &App{
		BSCRelayer:         bscRelayer,
		InscriptionRelayer: insRelayer,
	}
}

func (a *App) Start() {
	a.InscriptionRelayer.Start()
	a.BSCRelayer.Start()
}
