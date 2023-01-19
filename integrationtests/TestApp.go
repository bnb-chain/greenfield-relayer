package integrationtests

import (
	"fmt"

	"github.com/bnb-chain/inscription-relayer/app"
	"github.com/bnb-chain/inscription-relayer/config"
	"github.com/bnb-chain/inscription-relayer/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	err = db.Migrator().DropTable(&model.InscriptionRelayTransaction{})
	if err != nil {
		panic(err)
	}
	err = db.Migrator().DropTable(&model.InscriptionBlock{})
	if err != nil {
		panic(err)
	}
	return *app.NewApp(cfg)
}

func GetTestConfig() *config.Config {
	return config.ParseConfigFromFile("config/config_test.json")
}
