package handlers

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var db *sql.DB

func InitDB() {
	viper.SetConfigFile("../config/db.yaml")
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
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error")
	}
}
