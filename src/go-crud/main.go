package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maithanhduyan/golang-code/go-crud/controllers"
	"github.com/maithanhduyan/golang-code/go-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
