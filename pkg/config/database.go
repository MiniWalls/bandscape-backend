package mydb

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// StringInterfaceMap is used to represent a JSON object
type Post struct {
	ID               int              `json:"id"`
	Title            string           `json:"title"`
	Body             string           `json:"body"`
	Datetime         string           `json:"datetime"`
	Userid           string           `json:"userid"`
	LastfmAttachment *json.RawMessage `json:"lastfmattachment"`
}

type User struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
}

func DbConnection() {
	envErr := godotenv.Load() //Load .env file and error check
	if envErr != nil {
		log.Fatal("Error loading .env file", envErr)
	}

	// Initialize a new connection object to database
	cfg := mysql.Config{
		/* User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               "bandscape",
		AllowNativePasswords: true, */
	}
	fmt.Println(cfg) //VSCode removing mysql driver if I dont specifically use something from mysql

	// Initialize a new connection object to google MYSQL database
	var (
		dbUser    = os.Getenv("DB_USER")
		dbPwd     = os.Getenv("DB_PASS")
		dbName    = os.Getenv("DB_NAME")
		dbPort    = os.Getenv("DB_PORT")
		dbTCPHost = os.Getenv("DB_HOST")
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	// Get a database handle.
	var err error
	/* db, err = sql.Open("mysql", cfg.FormatDSN()) */ //For local MySQL
	db, err = sql.Open("mysql", dbURI)                 //For google MySQL
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// Return the database object
func GetDB() *sql.DB {
	return db
}

// Every function returns an error and GET functions return objects
func GetPosts() ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Datetime, &post.Userid, &post.LastfmAttachment)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Pointer to post object, so we can return nil if no post is found
func GetPost(id string) (*Post, error) {
	post := new(Post)
	err := db.QueryRow("SELECT * FROM post WHERE postid = ?", id).Scan(&post.ID, &post.Title, &post.Body, &post.Datetime, &post.Userid, &post.LastfmAttachment)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return post, nil
}

func UpdatePost(updatedPost Post) error {
	//Get the values from the updated post
	id := updatedPost.ID
	title := updatedPost.Title
	body := updatedPost.Body

	//Update the post with the new values
	res, err := db.Exec("UPDATE post SET title = ?, body = ? WHERE postid = ?", title, body, id)
	if err != nil {
		return errors.New(err.Error())
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return errors.New("No post with that id")
	} else {
		return nil
	}
}

func DeletePost(id string) error {
	res, err := db.Exec("DELETE FROM post WHERE postid = ?", id)
	if err != nil {
		return errors.New(err.Error())
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return errors.New("No post with that id")
	} else {
		return nil
	}
}

func CreatePost(post Post) error {
	title := post.Title
	body := post.Body
	userid := post.Userid
	lastfm_attachment := post.LastfmAttachment
	res, err := db.Exec("INSERT INTO post (title, body, userid, lastfm_attachment) VALUES (?, ?, ?, CAST(? AS JSON))", title, body, userid, lastfm_attachment)
	if err != nil {
		return errors.New(err.Error())
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return errors.New("No post was created")
	} else {
		return nil
	}
}

func GetUser(id string) bool {
	var user User
	err := db.QueryRow("SELECT * FROM user WHERE userid = ?", id).Scan(&user.Userid, &user.Username)
	if err != nil {
		return false
	}
	log.Println(user)
	if user.Userid == id {
		return true
	} else {
		return false
	}
}

func CreateUser(user User) error {
	res, err := db.Exec("INSERT INTO user (userid, username) VALUES (?, ?)", user.Userid, user.Username)
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return errors.New("No user was created")
	} else {
		return nil
	}
}

func UpdateUser(updatedUser User) error {
	res, err := db.Exec("UPDATE user SET username = ? WHERE userid = ?", updatedUser.Username, updatedUser.Userid)
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return errors.New("No user was created")
	} else {
		return nil
	}
}
