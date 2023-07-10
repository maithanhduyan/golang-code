package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	Username string
	Password string
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ...

	// Đăng nhập người dùng (ví dụ)
	user := User{
		ID:       1,
		Username: "john",
		Password: "secret",
	}

	// Tạo JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Ký JWT với một secret key (hãy thay thế bằng secret key thực tế của bạn)
	secretKey := []byte("your-secret-key")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	// Lưu JWT vào cơ sở dữ liệu
	insertTokenQuery := "INSERT INTO tokens (user_id, token) VALUES (?, ?)"
	_, err = db.Exec(insertTokenQuery, user.ID, tokenString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JWT created and saved successfully.")
}
