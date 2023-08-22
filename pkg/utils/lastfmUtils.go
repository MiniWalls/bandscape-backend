package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
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

	//Append secret key
	encodedParams += os.Getenv("LASTFM_API_SECRET")

	//Hash with md5
	hash := md5.New()
	hash.Write([]byte(encodedParams))
	signature := hex.EncodeToString(hash.Sum(nil))

	return signature
}

// Creates query string used in lastfm api requests
// Currently errors from not handling spaces correctly
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
		concated += k + "=" + url.QueryEscape(params[k]) + "&"
	}

	encodedParams := strings.TrimSuffix(concated, "&")

	return encodedParams
}
