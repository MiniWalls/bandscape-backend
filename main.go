package main

import (
	mydb "bandscape-backend/pkg/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPosts(c *gin.Context) {
	posts := mydb.GetPosts()
	c.IndentedJSON(http.StatusOK, posts)
}

func main() {
	mydb.DbConnection()
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run("localhost:3001")
	fmt.Println("Server running on port 3001")
}
