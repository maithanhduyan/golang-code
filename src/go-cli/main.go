package main

import (
	"flag"
	"fmt"
)

func main() {
	// Định nghĩa các cờ và tham số dòng lệnh
	greeting := flag.String("greeting", "Hello", "Greeting message")
	name := flag.String("name", "World", "Name to greet")
	flag.Parse()

	// In ra lời chào
	message := fmt.Sprintf("%s, %s!", *greeting, *name)
	fmt.Println(message)
}
