package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	ID   int
	Name string
}

func main() {
	// Kết nối tới cơ sở dữ liệu SQLite
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println("Lỗi khi kết nối tới cơ sở dữ liệu:", err)
		return
	}
	defer db.Close()

	// Tạo bảng "items" trong cơ sở dữ liệu
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	if err != nil {
		fmt.Println("Lỗi khi tạo bảng:", err)
		return
	}

	// Thêm một mục mới vào cơ sở dữ liệu
	item := Item{Name: "Mục 1"}
	_, err = db.Exec("INSERT INTO items (name) VALUES (?)", item.Name)
	if err != nil {
		fmt.Println("Lỗi khi thêm mục:", err)
		return
	}

	// Truy vấn tất cả các mục từ cơ sở dữ liệu
	rows, err := db.Query("SELECT id, name FROM items")
	if err != nil {
		fmt.Println("Lỗi khi truy vấn dữ liệu:", err)
		return
	}
	defer rows.Close()

	// Lặp qua các mục và in ra thông tin
	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.Name)
		if err != nil {
			fmt.Println("Lỗi khi đọc dữ liệu:", err)
			return
		}
		items = append(items, item)
	}

	// In thông tin các mục
	for _, item := range items {
		fmt.Println(item.ID, item.Name)
	}
}
