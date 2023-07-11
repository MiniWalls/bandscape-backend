package main

import (
	mydb "bandscape-backend/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getPosts(c *gin.Context) {
	posts := mydb.GetPosts()
	c.IndentedJSON(http.StatusOK, posts)
}

func updatePost(c *gin.Context) {
	var updatedPost mydb.Post

	if err := c.BindJSON(&updatedPost); err != nil {
		return
	}
	mydb.UpdatePost(updatedPost)
	c.IndentedJSON(http.StatusOK, updatedPost)
}

func main() {
	mydb.DbConnection()
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/posts", getPosts)
	router.PUT("/post", updatePost)
	router.Run("localhost:3001")
	fmt.Println("Server running on port 3001")
}
