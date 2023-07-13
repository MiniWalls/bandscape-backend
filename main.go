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
	c.IndentedJSON(http.StatusOK, "Post updated")
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	mydb.DeletePost(id)
	c.IndentedJSON(http.StatusOK, "Post deleted")
}

func getPost(c *gin.Context) {
	id := c.Param("id")
	post := mydb.GetPost(id)
	c.IndentedJSON(http.StatusOK, post)
}

func main() {
	mydb.DbConnection()
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPost)
	router.DELETE("/posts/:id", deletePost)
	router.PUT("/post", updatePost)
	router.Run("localhost:3001")
	fmt.Println("Server running on port 3001")
}
