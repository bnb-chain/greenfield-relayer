package executor

import (
	"github.com/bnb-chain/inscription-relayer/config"
)

func InitTestConfig() *config.Config {
	return config.ParseConfigFromFile("../integrationtests/config/config_test.json")
}

func InitExecutors() (*BSCExecutor, *InscriptionExecutor) {
	cfg := InitTestConfig()
	insExecutor := NewInscriptionExecutor(cfg)
	bscExecutor := NewBSCExecutor(cfg)
	insExecutor.SetBSCExecutor(bscExecutor)
	bscExecutor.SetInscriptionExecutor(insExecutor)
	return bscExecutor, insExecutor
}
