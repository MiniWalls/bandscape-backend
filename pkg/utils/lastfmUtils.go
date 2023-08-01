package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/joho/godotenv"
)

// Creates api signature used to get user session key
func CreateApiSignature(params map[string]string) string {
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

// Creates query string used in lastfm api requests
func CreateQueryString(params map[string]string) string {
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
		concated += k + "=" + params[k] + "&"
	}

	encodedParams := strings.TrimSuffix(concated, "&")

	return encodedParams
}