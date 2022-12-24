package main

import (
	"github.com/buabaj/post-api/initializers"
	"github.com/buabaj/post-api/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
