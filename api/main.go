package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jin237/mdet/go/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: アーキテクチャを決め次第別ファイルに定義する
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var Db *gorm.DB

func main() {
	config.LoadConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	var err error
	if Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		connectDb(dsn, 5)
	}

	Db.AutoMigrate(&Product{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000/",
		},
	}))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
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
