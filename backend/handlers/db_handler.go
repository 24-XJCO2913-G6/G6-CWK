package handlers

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"log"
	. "main/backend/models"
)

func InitDB() {
	viper.SetConfigFile("./backend/configs/db.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var dbConfig DatabaseConfig
	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %s", err))
	}

	dataSourceName :=
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConfig.User, dbConfig.Password,
			dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	var err error
	Db, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error")
	}

	BuildModel()
}
