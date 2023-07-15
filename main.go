package main

import (
	mydb "bandscape-backend/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getPosts(c *gin.Context) {
	posts, err := mydb.GetPosts()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, posts)
	}
}

func updatePost(c *gin.Context) {
	var updatedPost mydb.Post

	if err := c.BindJSON(&updatedPost); err != nil {
		return
	}

	if err := mydb.UpdatePost(updatedPost); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, "Post updated")
	}
}

func deletePost(c *gin.Context) {
	id := c.Param("id")

	if err := mydb.DeletePost(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, "Post deleted")
	}
}

func createPost(c *gin.Context) {
	var newPost mydb.Post

	if err := c.BindJSON(&newPost); err != nil {
		return
	}
	if err := mydb.CreatePost(newPost); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, "Post created")
	}
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	post, err := mydb.GetPost(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, post)
	}
}

func main() {
	mydb.DbConnection()
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPost)
	router.POST("/posts", createPost)
	router.DELETE("/posts/:id", deletePost)
	router.PUT("/posts", updatePost)
	router.Run("localhost:3001")
	fmt.Println("Server running on port 3001")
}
