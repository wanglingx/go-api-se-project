package app

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

var BetadineProject Store

func InitDB(c * Configs) {
	var err error

	db, err := gorm.Open(mysql.Open(cv.ConnectionMasterDB), &gorm.Config{})
	if err != nil {
		log.Errorln("Failed to connect database: ", c.ConnectionMasterDB)
		log.Errorln("Error: ", err)
		panic("Failed to connect database MySql")
	}

	log.Info("Successfully connected to database")

	BetadineProject.DB = db
}