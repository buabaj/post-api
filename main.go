package main

import (
	"github.com/buabaj/post-api/controllers"
	"github.com/buabaj/post-api/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.Root)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Use(cors.Default())

	r.Run()
}
