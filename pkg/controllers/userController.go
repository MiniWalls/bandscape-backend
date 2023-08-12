package controllers

import (
	mydb "bandscape-backend/pkg/config"
	"encoding/json"
	"log"
)

type AuthResponse struct {
	Key      string `json:"key"`
	Username string `json:"username"`
}

// This is test function which has its own route, not used in production
/* func RegisterUser(c *gin.Context) {
	userid := c.Query("userid")
	username := c.Query("username")
	user := mydb.User{
		Userid:   userid,
		Username: username,
	}
	log.Println("userid", userid)
	log.Println("username", username)
	if doesUserExist(userid) {
		UpdateUser(user)
		c.IndentedJSON(200, "Updated user")
	} else {
		CreateUser(user)
		c.IndentedJSON(200, "Created user")
	}
} */

// This will be real function used
// RegisterUser function is used to either create or update a user in the database
func RegisterUser(body []byte) {
	//We parse our body data, and then set it to fields used by our MySQL model
	var responseMap map[string]interface{}
	err := json.Unmarshal(body, &responseMap)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}

	session := responseMap["session"].(map[string]interface{})

	key := session["key"].(string)
	name := session["name"].(string)

	log.Println("Key:", key)
	log.Println("Name:", name)

	user := mydb.User{
		Userid:   name,
		Username: key,
	}

	if doesUserExist(user.Userid) {
		UpdateUser(user)
	} else {
		CreateUser(user)
	}
}

func doesUserExist(userid string) bool {
	res := mydb.GetUser(userid)
	return res
}

func UpdateUser(user mydb.User) {
	res := mydb.UpdateUser(user)
	if res != nil {
		log.Println("Error updating user")
	} else {
		log.Println("User updated")
	}
}

func CreateUser(user mydb.User) {
	res := mydb.CreateUser(user)
	if res != nil {
		log.Println("Error creating user")
	} else {
		log.Println("User created")
	}
}
