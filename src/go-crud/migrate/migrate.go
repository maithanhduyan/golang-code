package main

import (
	"github.com/maithanhduyan/golang-code/go-crud/initializers"
	"github.com/maithanhduyan/golang-code/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
