package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Tạo một đối tượng Person
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Mã hóa đối tượng thành JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Lỗi khi mã hóa dữ liệu thành JSON:", err)
		return
	}

	// In ra kết quả JSON
	fmt.Println(string(jsonData))
}
