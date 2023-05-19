#

- Create new module
> go mod init sample.com/restful-api-assets-postgresql

- Add GO dependencies
~~~~
go get github.com/gin-gonic/gin
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres

~~~~

# Dữ liệu mẫu

Dùng python để tạo nhiều dữ liệu mẫu
1 - Tạo môi trường python ảo
> python -m venv venv

    - Activate môi trường python ảo
> venv\Scripts\activate
