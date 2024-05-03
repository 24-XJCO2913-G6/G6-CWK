package models

import (
	"github.com/go-xorm/xorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var Db *xorm.Engine

func BuildModel() {
	BuildModelUser()
	BuildModelTrack()
}
