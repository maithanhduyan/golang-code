package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// Đăng ký một công việc theo lịch trình
	c.AddFunc("0 0 * * * *", func() {
		fmt.Println("Công việc chạy vào mỗi giờ đầu tiên của mỗi ngày")
	})

	// Đăng ký một công việc chạy mỗi phút
	c.AddFunc("* * * * *", func() {
		fmt.Println("Công việc chạy mỗi phút")
	})

	// Khởi động lịch trình
	c.Start()

	// Chạy vòng lặp chính của ứng dụng
	// Đảm bảo ứng dụng không kết thúc ngay lập tức sau khi tạo lịch trình
	select {}
}
