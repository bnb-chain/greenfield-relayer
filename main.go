package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/bnb-chain/greenfield-relayer/app"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/logging"
)

func initFlags() {
	flag.String(config.FlagConfigPath, "", "config file path")
	flag.String(config.FlagConfigType, "", "config type, local_private_key or aws_private_key")
	flag.String(config.FlagConfigAwsRegion, "", "aws region")
	flag.String(config.FlagConfigAwsSecretKey, "", "aws secret key")
	flag.String(config.FlagConfigPrivateKey, "", "relayer private key")
	flag.String(config.FlagConfigBlsPrivateKey, "", "relayer bls private key")
	flag.String(config.FlagConfigDbPass, "", "relayer db password")
	flag.String(config.FlagConfigDbUsername, "", "relayer db username")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Print("usage: ./greenfield-relayer --config-type local --config-path configFile\n")
	fmt.Print("usage: ./greenfield-relayer --config-type aws --aws-region awsRegin --aws-secret-key awsSecretKey\n")
}

func main() {
	var (
		cfg                        *config.Config
		configType, configFilePath string
	)
	initFlags()
	configType = viper.GetString(config.FlagConfigType)
	if configType == "" {
		configType = os.Getenv(config.ConfigType)
	}
	if configType != config.AWSConfig && configType != config.LocalConfig {
		printUsage()
		return
	}
	if configType == config.AWSConfig {
		awsSecretKey := viper.GetString(config.FlagConfigAwsSecretKey)
		if awsSecretKey == "" {
			printUsage()
			return
		}
		awsRegion := viper.GetString(config.FlagConfigAwsRegion)
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
		configFilePath = viper.GetString(config.FlagConfigPath)
		if configFilePath == "" {
			configFilePath = os.Getenv(config.ConfigFilePath)
			if configFilePath == "" {
				printUsage()
				return
			}
		}
		cfg = config.ParseConfigFromFile(configFilePath)
	}
	if cfg == nil {
		panic("failed to get configuration")
	}
	logging.InitLogger(&cfg.LogConfig)
	app.NewApp(cfg).Start()
	select {}
}
