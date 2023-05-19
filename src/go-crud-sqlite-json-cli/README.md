#### create go.mod file
~~~~
go mod init go-sample.com/crud-item-json-cli
~~~~
- Cài đặt gói github.com/mattn/go-sqlite3 bằng cách chạy lệnh sau trong terminal:

~~~~
go get github.com/mattn/go-sqlite3

~~~~

- Build chương trình
~~~~
    go build .
~~~~

- Chay chương trình trong terminal với tham số sau:
- Tạo một mục mới: go run main.go create "Tên mục" "Mã" 10.99 "Mô tả"
~~~~
crud-item-json-cli  create "Tên mục" "Mã" 10.99 "Mô tả"
~~~~

- Đọc tất cả các mục: go run main.go read
~~~~
crud-item-json-cli  read

~~~~

- Cập nhật một mục: go run main.go update 1 "Tên mới" "Mã mới" 15.99 "Mô tả mới"
~~~~
crud-item-json-cli  update 1 "Tên mới" "Mã mới" 15.99 "Mô tả mới"

~~~~

- Xóa một mục: go run main.go delete 1
~~~~
crud-item-json-cli delete 1
~~~~