package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Record struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Text string `json:"text"`
}

func main() {
	dsn := "root:12345678@tcp(mysql:3306)/mydbapp?charset=utf8&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ { // Попытка подключиться 10 раз с интервалом в 2 секунды
		db, err = gorm.Open("mysql", dsn)
		if err == nil {
			break
		}
		log.Printf("failed to connect database: %v", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect database after retries: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&Record{})

	r := gin.Default()

	// Добавление CORS middleware
	r.Use(cors.Default())

	r.POST("/records", func(c *gin.Context) {
		var record Record
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&record)
		c.JSON(200, record)
	})

	r.GET("/records", func(c *gin.Context) {
		var records []Record
		db.Find(&records)
		c.JSON(200, records)
	})

	r.Run(":8080")
}
