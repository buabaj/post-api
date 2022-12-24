package controllers

import (
	"log"

	"github.com/buabaj/post-api/initializers"
	"github.com/buabaj/post-api/models"
	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome to the post api",
	})
}

func CreatePost(c *gin.Context) {
	var body struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Author string `json:"author"`
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body, Author: body.Author}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"message": "post created successfully",
		"post":    post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"message": "posts fetched successfully",
		"posts":   posts,
	})
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"message": "post fetched",
		"post":    post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	var body struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		Author string `json:"author"`
	}
	c.Bind(&body)

	result = initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, Author: body.Author})

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"message": "post updated",
		"post":    post,
	})
}

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	result = initializers.DB.Delete(&post, id)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"message": "post deleted",
	})
}
