package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maithanhduyan/golang-code/go-crud/initializers"
	"github.com/maithanhduyan/golang-code/go-crud/models"
)

func PostCreate(c *gin.Context) {

	// get data of req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"post": posts,
	})

}

func PostShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	println(id)
	// Get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond result
	c.Status(200)

}
