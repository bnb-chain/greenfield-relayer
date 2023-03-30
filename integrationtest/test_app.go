package integrationtest

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/bnb-chain/greenfield-relayer/app"
	"github.com/bnb-chain/greenfield-relayer/config"
	"github.com/bnb-chain/greenfield-relayer/db/model"
)

func InitTestApp() app.App {
	cfg := GetTestConfig()
	username := cfg.DBConfig.Username
	password := viper.GetString(config.FlagConfigDbPass)
	url := cfg.DBConfig.Url
	dbPath := fmt.Sprintf("%s:%s@%s", username, password, url)
	db, err := gorm.Open(mysql.Open(dbPath), &gorm.Config{})
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
