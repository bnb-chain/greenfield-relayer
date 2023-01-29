package executor

import (
	"github.com/bnb-chain/greenfield-relayer/config"
)

func InitTestConfig() *config.Config {
	return config.ParseConfigFromFile("../integrationtests/config/config_test.json")
}

func InitExecutors() (*BSCExecutor, *GreenfieldExecutor) {
	cfg := InitTestConfig()
	insExecutor := NewGreenfieldExecutor(cfg)
	bscExecutor := NewBSCExecutor(cfg)
	insExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetGreenfieldExecutor(insExecutor)
	return bscExecutor, insExecutor
}
