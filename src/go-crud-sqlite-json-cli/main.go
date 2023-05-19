package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Tạo bảng "items" nếu chưa tồn tại
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		code TEXT,
		price REAL,
		description TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Kiểm tra và xử lý các tham số dòng lệnh
	if len(os.Args) < 2 {
		fmt.Println("Vui lòng cung cấp tham số 'create', 'read', 'update', hoặc 'delete'")
		return
	}

	command := os.Args[1]

	switch command {
	case "create":
		if len(os.Args) < 6 {
			fmt.Println("Vui lòng cung cấp đầy đủ thông tin (name, code, price, description)")
			return
		}

		name := os.Args[2]
		code := os.Args[3]
		price := os.Args[4]
		description := os.Args[5]

		_, err := db.Exec("INSERT INTO items (name, code, price, description) VALUES (?, ?, ?, ?)", name, code, price, description)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Mục đã được tạo thành công")

	case "read":
		rows, err := db.Query("SELECT id, name, code, price, description FROM items")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var items []Item

		for rows.Next() {
			var item Item
			err := rows.Scan(&item.ID, &item.Name, &item.Code, &item.Price, &item.Description)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}

		jsonData, err := json.Marshal(items)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonData))

	case "update":
		if len(os.Args) < 7 {
			fmt.Println("Vui lòng cung cấp đầy đủ thông tin (id, name, code, price, description)")
			return
		}

		id := os.Args[2]
		name := os.Args[3]
		code := os.Args[4]
		price := os.Args[5]
		description := os.Args[6]

		_, err := db.Exec("UPDATE items SET name=?, code=?, price=?, description=? WHERE id=?", name, code, price, description, id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Mục đã được cập nhật thành công")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Vui lòng cung cấp id của mục để xóa")
			return
		}

		id := os.Args[2]

		_, err := db.Exec("DELETE FROM items WHERE id=?", id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Mục đã được xóa thành công")

	default:
		fmt.Println("Tham số không hợp lệ")
	}
}
