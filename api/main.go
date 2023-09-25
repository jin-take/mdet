package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: アーキテクチャを決め次第別ファイルに定義する
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// TODO: アーキテクチャを決め次第configで別ファイルに定義する
	dsn := "mdet:password@tcp(db:3306)/mdet?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
