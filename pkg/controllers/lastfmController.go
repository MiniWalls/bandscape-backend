package controllers

import (
	utils "bandscape-backend/pkg/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetTrack(c *gin.Context) {
	trackName := c.Query("track")
	artistName := c.Query("artist")
	username := c.Query("username")

	api_key := os.Getenv("LASTFM_API_KEY")

	queryParams := map[string]string{
		"method":      "track.getInfo",
		"track":       trackName,
		"artist":      artistName,
		"username":    username,
		"autocorrect": "1",
		"api_key":     api_key,
		"format":      "json",
	}

	// Create the query string
	queryString := utils.CreateQueryString(queryParams)

	requestURL := os.Getenv("LASTFM_API_URL") + "?" + queryString

	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", body)
}

func GetAlbum(c *gin.Context) {
	albumName := c.Query("album")
	artistName := c.Query("artist")
	username := c.Query("username")

	api_key := os.Getenv("LASTFM_API_KEY")

	queryParams := map[string]string{
		"method":      "album.getInfo",
		"album":       albumName,
		"artist":      artistName,
		"username":    username,
		"autocorrect": "1",
		"api_key":     api_key,
		"format":      "json",
	}

	// Create the query string
	queryString := utils.CreateQueryString(queryParams)

	requestURL := os.Getenv("LASTFM_API_URL") + "?" + queryString

	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", body)
}
