package controllers

import (
	mydb "bandscape-backend/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
	posts, err := mydb.GetPosts()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, posts)
	}
}

func UpdatePost(c *gin.Context) {
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

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	if err := mydb.DeletePost(id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, "Post deleted")
	}
}

func CreatePost(c *gin.Context) {
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

func GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := mydb.GetPost(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
	} else {
		c.IndentedJSON(http.StatusOK, post)
	}
}
