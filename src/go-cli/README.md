#### Dưới đây là một ví dụ đơn giản về cách bạn có thể bắt đầu xây dựng một ứng dụng CLI bằng Golang:

Tạo một tệp tin mới, ví dụ: main.go.



Trong tệp tin main.go, nhập các gói cần thiết và định nghĩa hàm main():
~~~~
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
~~~~


Biên dịch và chạy ứng dụng CLI:
~~~~
$ go build
$ ./main --greeting "Hi" --name "John"

~~~~

Kết quả sẽ là `Hi, John!`

Trong ví dụ trên, chúng ta sử dụng gói flag để định nghĩa và xử lý các cờ và tham số dòng lệnh. Bạn có thể tùy chỉnh các cờ và tham số dòng lệnh theo nhu cầu của ứng dụng của mình.

Đây chỉ là một ví dụ cơ bản để giúp bạn bắt đầu. Bằng cách sử dụng Golang, bạn có thể xây dựng các ứng dụng CLI phức tạp hơn, với nhiều tính năng và chức năng tùy chỉnh.