package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Asset là mô hình dữ liệu cho tài sản
type Asset struct {
	ID          uint    `gorm:"primary_key" json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Website     string  `json:"website"`
	Description string  `json:"description"`
}

func main() {
	// Kết nối tới cơ sở dữ liệu SQLite
	db, err := gorm.Open("sqlite3", "assets.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Migrate cơ sở dữ liệu và tạo bảng 'assets' nếu chưa tồn tại
	db.AutoMigrate(&Asset{})

	// Khởi tạo router của Gin
	router := gin.Default()

	// Định nghĩa các endpoint
	router.GET("/assets", getAssets)
	router.GET("/assets/:id", getAsset)
	router.POST("/assets", createAsset)
	router.PUT("/assets/:id", updateAsset)
	router.DELETE("/assets/:id", deleteAsset)

	// Khởi động server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}

// Handler cho GET /assets
func getAssets(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*gorm.DB)

	var assets []Asset
	conn.Find(&assets)

	c.JSON(http.StatusOK, assets)
}

// Handler cho GET /assets/:id
func getAsset(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*gorm.DB)

	id := c.Param("id")
	var asset Asset
	conn.First(&asset, id)

	if asset.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

// Handler cho POST /assets
func createAsset(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*gorm.DB)

	var asset Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn.Create(&asset)

	c.JSON(http.StatusCreated, asset)
}

// Handler cho PUT /assets/:id
func updateAsset(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*gorm.DB)

	id := c.Param("id")
	var asset Asset
	conn.First(&asset, id)

	if asset.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn.Save(&asset)

	c.JSON(http.StatusOK, asset)
}

// Handler cho DELETE /assets/:id
func deleteAsset(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(*gorm.DB)

	id := c.Param("id")
	var asset Asset
	conn.First(&asset, id)

	if asset.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	conn.Delete(&asset)

	c.JSON(http.StatusNoContent, nil)
}
