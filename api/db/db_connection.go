package db

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/jin237/mdet/config"
	"github.com/jin237/mdet/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	var err error
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		connectDb(dsn, 5)
	}

	db.AutoMigrate(&models.User{})

	return nil
}

func connectDb(dsn string, retryCount int) {
	slog.Info("Retry connect DB")
	var err error
	retryCount--
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		if retryCount > 0 {
			time.Sleep(10 * time.Second)
			connectDb(dsn, retryCount)
			return
		}
		slog.Error(err.Error())
		panic("failed to connect database")
	}
}
