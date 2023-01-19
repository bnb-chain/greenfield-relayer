package main

import (
	"flag"
	"fmt"

	"github.com/bnb-chain/inscription-relayer/app"

	"github.com/bnb-chain/inscription-relayer/common"
	config "github.com/bnb-chain/inscription-relayer/config"
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

	if cfg.DBConfig.DBPath == "" {
		panic("DB config is not present in config file, please follow instruction to specify it")
	}

	app.NewApp(cfg).Start()
	select {}
}
