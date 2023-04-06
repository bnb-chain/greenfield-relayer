package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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
	username := cfg.DBConfig.Username
	password := viper.GetString(config.FlagConfigDbPass)
	if password == "" {
		password = getDBPass(&cfg.DBConfig)
	}

	url := cfg.DBConfig.Url
	dbPath := fmt.Sprintf("%s:%s@%s", username, password, url)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	dbConfig, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbConfig.SetMaxIdleConns(cfg.DBConfig.MaxIdleConns)
	dbConfig.SetMaxOpenConns(cfg.DBConfig.MaxOpenConns)

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

	metricService := metric.NewMetricService(cfg)

	// vote signer
	signer := vote.NewVoteSigner(greenfieldExecutor.BlsPrivateKey)

	// voteProcessors
	greenfieldVoteProcessor := vote.NewGreenfieldVoteProcessor(cfg, daoManager, signer, greenfieldExecutor)
	bscVoteProcessor := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor)

	// listeners
	greenfieldListener := listener.NewGreenfieldListener(cfg, greenfieldExecutor, bscExecutor, daoManager, metricService)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, greenfieldExecutor, daoManager, metricService)

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

func getDBPass(cfg *config.DBConfig) string {
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			panic(err)
		}
		type DBPass struct {
			DbPass string `json:"db_pass"`
		}
		var dbPassword DBPass
		err = json.Unmarshal([]byte(result), &dbPassword)
		if err != nil {
			panic(err)
		}
		return dbPassword.DbPass
	}
	return cfg.Password
}
