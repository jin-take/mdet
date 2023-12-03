package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jin237/mdet/config"
	"github.com/jin237/mdet/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	if Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		connectDb(dsn, 5)
	}

	Db.AutoMigrate(&models.User{})

	return nil
}

func connectDb(dsn string, retryCount int) {
	fmt.Println("Retry connect DB")
	var err error
	retryCount--
	if Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		if retryCount > 0 {
			time.Sleep(10 * time.Second)
			connectDb(dsn, retryCount)
			return
		}
		log.Fatal(err)
		panic("failed to connect database")
	}
}
