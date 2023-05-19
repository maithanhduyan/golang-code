package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Asset struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Website     string  `json:"website"`
	Description string  `json:"description"`
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=go-assets sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tạo bảng assets nếu chưa tồn tại
	db.AutoMigrate(&Asset{})

	r := gin.Default()

	// Lấy danh sách Asset
	r.GET("/assets", func(c *gin.Context) {
		var assets []Asset
		if err := db.Find(&assets).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, assets)
	})

	// Tạo mới một Asset
	r.POST("/assets", func(c *gin.Context) {
		var asset Asset
		if err := c.BindJSON(&asset); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&asset).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, asset)
	})

	r.Run(":8080")
}
