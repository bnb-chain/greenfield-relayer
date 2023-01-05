package main

import (
	"flag"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"inscription-relayer/assembler"
	"inscription-relayer/common"
	"inscription-relayer/config"
	"inscription-relayer/db/dao"
	"inscription-relayer/db/model"
	"inscription-relayer/executor"
	"inscription-relayer/listener"
	"inscription-relayer/relayer"
	"inscription-relayer/vote"
)

const (
	flagConfigPath         = "config-path"
	flagConfigType         = "config-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config file path")
	flag.String(flagConfigType, "local_private_key", "config type, local_private_key or aws_private_key")
	flag.String(flagConfigAwsRegion, "", "aws region")
	flag.String(flagConfigAwsSecretKey, "", "aws secret key")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Print("usage: ./inscription-relayer --config-type local --config-path configFile\n")
	fmt.Print("usage: ./inscription-relayer --config-type aws --aws-region awsRegin --aws-secret-key awsSecretKey\n")
}

func main() {
	initFlags()

	configType := viper.GetString(flagConfigType)
	if configType != config.AWSConfig && configType != config.LocalConfig {
		printUsage()
		return
	}

	var cfg *config.Config
	if configType == config.AWSConfig {
		awsSecretKey := viper.GetString(flagConfigAwsSecretKey)
		if awsSecretKey == "" {
			printUsage()
			return
		}

		awsRegion := viper.GetString(flagConfigAwsRegion)
		if awsRegion == "" {
			printUsage()
			return
		}

		configContent, err := config.GetSecret(awsSecretKey, awsRegion)
		if err != nil {
			fmt.Printf("get aws config error, err=%s", err.Error())
			return
		}
		cfg = config.ParseConfigFromJson(configContent)
	} else {
		configFilePath := viper.GetString(flagConfigPath)
		configFilePath = "config/config.json"
		if configFilePath == "" {
			printUsage()
			return
		}

		cfg = config.ParseConfigFromFile(configFilePath)
	}

	if cfg == nil {
		fmt.Println("failed to get configuration")
		return
	}

	common.InitLogger(&cfg.LogConfig)

	var db *gorm.DB
	if cfg.DBConfig.DBPath != "" {
		var err error
		db, err = gorm.Open(cfg.DBConfig.Dialect, cfg.DBConfig.DBPath)
		if err != nil {
			panic(fmt.Sprintf("open db error, err=%s", err.Error()))
		}
		defer db.Close()
		model.InitBSCTables(db)
		model.InitInscriptionTables(db)
		model.InitVoteTables(db)
	}

	inscriptionDao := dao.NewInscriptionDao(db)
	bscDao := dao.NewBSCDao(db)
	voteDao := dao.NewVoteDao(db)
	daoManager := dao.NewDaoManager(inscriptionDao, bscDao, voteDao)

	inscriptionExecutor, err := executor.NewInscriptionExecutor(cfg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	bscExecutor, err := executor.NewBSCExecutor(cfg, daoManager)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	inscriptionExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetInscriptionExecutor(inscriptionExecutor)

	votePoolExecutor := vote.NewVotePoolExecutor(cfg, inscriptionExecutor)

	//listener
	inscriptionListener := listener.NewInscriptionListener(cfg, inscriptionExecutor, daoManager)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, daoManager)

	//vote signer
	signer, err := vote.NewVoteSigner(ethcommon.Hex2Bytes(cfg.VotePoolConfig.BlsPrivateKey))
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	//voteProcessor
	inscriptionVoteProcessor, err := vote.NewInscriptionVoteProcessor(cfg, daoManager, signer, inscriptionExecutor, votePoolExecutor)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	bscVoteProcessor, err := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor, votePoolExecutor)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	//assembler
	inscriptionAssembler := assembler.NewInscriptionAssembler(cfg, inscriptionExecutor, daoManager, bscExecutor, votePoolExecutor)
	bscAssembler := assembler.NewBSCAssembler(cfg, bscExecutor, daoManager, votePoolExecutor, inscriptionExecutor)

	//Relayer
	insRelayer := relayer.NewInscriptionRelayer(inscriptionListener, inscriptionExecutor, bscExecutor, votePoolExecutor, inscriptionVoteProcessor, inscriptionAssembler)
	bscRelayer := relayer.NewBSCRelayer(bscListener, inscriptionExecutor, bscExecutor, votePoolExecutor, bscVoteProcessor, bscAssembler)

	go insRelayer.Start()
	go bscRelayer.Start()

	select {}
}
