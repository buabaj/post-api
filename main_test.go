package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/buabaj/post-api/controllers"
	"github.com/buabaj/post-api/initializers"
	"github.com/buabaj/post-api/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.EmptyDB()
	initializers.DB.AutoMigrate(&models.Post{})
}

func setUpRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func TestRootHandler(t *testing.T) {
	mockResponse := `{"message":"welcome to the post api"}`
	r := setUpRouter()
	r.GET("/", controllers.Root)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, string(responseData))

}

func TestCreatePostHandler(t *testing.T) {
	r := setUpRouter()
	r.POST("/posts", controllers.CreatePost)

	post := models.Post{Title: "test title", Body: "test body"}

	jsonData, _ := json.Marshal(post)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestGetPostsHandler(t *testing.T) {
	r := setUpRouter()
	r.GET("/posts", controllers.GetPosts)

	req, _ := http.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetPostHandler(t *testing.T) {
	r := setUpRouter()
	r.GET("/posts/:id", controllers.GetPost)

	req, _ := http.NewRequest("GET", "/posts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdatePostHandler(t *testing.T) {
	r := setUpRouter()
	r.PUT("/posts/:id", controllers.UpdatePost)

	post := models.Post{Title: "test title", Body: "test body"}

	jsonData, _ := json.Marshal(post)
	req, _ := http.NewRequest("PUT", "/posts/1", bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestDeletePostHandler(t *testing.T) {
	r := setUpRouter()
	r.DELETE("/posts/:id", controllers.DeletePost)

	req, _ := http.NewRequest("DELETE", "/posts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
