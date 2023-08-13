package controllers

import (
	utils "bandscape-backend/pkg/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	redirectURL := "http://www.lastfm.com/api/auth/?api_key=" + os.Getenv("LASTFM_API_KEY") + "&cb=" + os.Getenv("LASTFM_CALLBACK_URL")
	c.Redirect(http.StatusFound, redirectURL)
}

func GetAuth(c *gin.Context) {
	token := c.Query("token")

	api_key := os.Getenv("LASTFM_API_KEY")

	params := map[string]string{
		"api_key": api_key,
		"method":  "auth.getSession",
		"token":   token,
	}
	api_sig := utils.CreateApiSignature(params)

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

	RegisterUser(body)

	//Instead of IndentedJSON we use Data so we can convert the byte slice to json.
	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "application/json", body)
}
