package main

import (
	mydb "bandscape-backend/pkg/config"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func getToken(c *gin.Context) {
	envErr := godotenv.Load() //Load .env file and error check
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	redirectURL := "http://www.lastfm.com/api/auth/?api_key=" + os.Getenv("LASTFM_API_KEY") + "&cb=" + os.Getenv("LASTFM_CALLBACK_URL")
	c.Redirect(http.StatusFound, redirectURL)
}

func getAuth(c *gin.Context) {
	token := c.Query("token")

	envErr := godotenv.Load() //Load .env file and error check
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	api_key := os.Getenv("LASTFM_API_KEY")

	params := map[string]string{
		"api_key": api_key,
		"method":  "auth.getSession",
		"token":   token,
	}
	api_sig := createApiSignature(params)

	//Send request to last.fm api
	url := os.Getenv("LASTFM_API_URL")
	query_params := map[string]string{
		"method":  "auth.getSession",
		"token":   token,
		"api_key": api_key,
		"api_sig": api_sig,
		"format":  "json",
	}

	// Create the query string
	queryString := ""
	for k, v := range query_params {
		queryString += "&" + k + "=" + v
	}
	queryString = strings.TrimPrefix(queryString, "&")

	// Construct the complete URL for the API request
	requestURL := url + "?" + queryString

	// Make the HTTP request
	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("Error making the API request:", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", body)
}

func createApiSignature(params map[string]string) string {
	//Parameters sorted alphabetically
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//Concatenate all parameters
	var concated string
	for _, k := range keys {
		fmt.Println(k, params[k])
		concated += k + params[k]
	}

	encodedParams := url.QueryEscape(concated)

	//Append secret
	envErr := godotenv.Load() //Load .env file and error check
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	encodedParams += os.Getenv("LASTFM_API_SECRET")

	//Hash with md5
	hash := md5.New()
	hash.Write([]byte(encodedParams))
	signature := hex.EncodeToString(hash.Sum(nil))

	return signature
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
	router.GET("/login", getToken)
	router.GET("/auth", getAuth)
	router.Run("localhost:3001")
	fmt.Println("Server running on port 3001")
}
