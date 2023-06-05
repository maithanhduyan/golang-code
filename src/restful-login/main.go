package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Mở kết nối đến cơ sở dữ liệu SQLite
	var err error
	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tạo bảng users trong cơ sở dữ liệu nếu nó chưa tồn tại
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Khởi tạo router
	r := gin.Default()

	// Đăng ký các đường dẫn xử lý API
	r.POST("/login", loginHandler)

	// Chạy server
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c *gin.Context) {
	// Lấy thông tin người dùng từ yêu cầu
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tìm người dùng trong cơ sở dữ liệu
	row := db.QueryRow("SELECT id FROM users WHERE username=? AND password=?", user.Username, user.Password)

	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Đăng nhập thành công
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
