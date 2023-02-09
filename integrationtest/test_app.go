package integrationtest

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/app"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db/model"
)

func InitTestApp() app.App {
	cfg := GetTestConfig()
	db, err := gorm.Open(mysql.Open(cfg.DBConfig.DBPath), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open db error, err=%s", err.Error()))
	}
	err = db.Migrator().DropTable(&model.BscRelayPackage{})
	if err != nil {
		panic(err)
	}
	err = db.Migrator().DropTable(&model.BscBlock{})
	if err != nil {
		panic(err)
	}
	err = db.Migrator().DropTable(&model.Vote{})
	if err != nil {
		panic(err)
	}
	err = db.Migrator().DropTable(&model.GreenfieldRelayTransaction{})
	if err != nil {
		panic(err)
	}
	err = db.Migrator().DropTable(&model.GreenfieldBlock{})
	if err != nil {
		panic(err)
	}
	return *app.NewApp(cfg)
}

func GetTestConfig() *config.Config {
	return config.ParseConfigFromFile("config/config_test.json")
}
