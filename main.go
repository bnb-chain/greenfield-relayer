package main

import (
	"flag"
	"fmt"

	"github.com/bnb-chain/inscription-relayer/assembler"
	"github.com/bnb-chain/inscription-relayer/common"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db/dao"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"github.com/bnb-chain/inscription-relayer/executor"
	"github.com/bnb-chain/inscription-relayer/listener"
	"github.com/bnb-chain/inscription-relayer/relayer"
	"github.com/bnb-chain/inscription-relayer/vote"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

	configType = "local"

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
		db, err = gorm.Open(mysql.Open(cfg.DBConfig.DBPath), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("open db error, err=%s", err.Error()))
		}
		model.InitBSCTables(db)
		model.InitInscriptionTables(db)
		model.InitVoteTables(db)
	}

	inscriptionDao := dao.NewInscriptionDao(db)
	bscDao := dao.NewBSCDao(db)
	voteDao := dao.NewVoteDao(db)
	daoManager := dao.NewDaoManager(inscriptionDao, bscDao, voteDao)

	inscriptionExecutor := executor.NewInscriptionExecutor(cfg)
	bscExecutor := executor.NewBSCExecutor(cfg, daoManager)

	inscriptionExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetInscriptionExecutor(inscriptionExecutor)

	votePoolExecutor := vote.NewVotePoolExecutor(cfg, inscriptionExecutor)

	// listener
	inscriptionListener := listener.NewInscriptionListener(cfg, inscriptionExecutor, daoManager)
	bscListener := listener.NewBSCListener(cfg, bscExecutor, daoManager)

	// vote signer
	signer := vote.NewVoteSigner(ethcommon.Hex2Bytes(cfg.VotePoolConfig.BlsPrivateKey))

	// voteProcessor
	inscriptionVoteProcessor := vote.NewInscriptionVoteProcessor(cfg, daoManager, signer, inscriptionExecutor, votePoolExecutor)
	bscVoteProcessor := vote.NewBSCVoteProcessor(cfg, daoManager, signer, bscExecutor, votePoolExecutor)

	// assembler
	inscriptionAssembler := assembler.NewInscriptionAssembler(cfg, inscriptionExecutor, daoManager, bscExecutor, votePoolExecutor)
	bscAssembler := assembler.NewBSCAssembler(cfg, bscExecutor, daoManager, votePoolExecutor, inscriptionExecutor)

	// Relayer
	insRelayer := relayer.NewInscriptionRelayer(inscriptionListener, inscriptionExecutor, bscExecutor, votePoolExecutor, inscriptionVoteProcessor, inscriptionAssembler)
	bscRelayer := relayer.NewBSCRelayer(bscListener, inscriptionExecutor, bscExecutor, votePoolExecutor, bscVoteProcessor, bscAssembler)

	go insRelayer.Start()
	go bscRelayer.Start()

	select {}
}
