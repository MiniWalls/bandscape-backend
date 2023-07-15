package mydb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Datetime string `json:"datetime"`
	Userid   string `json:"userid"`
}

func DbConnection() {
	envErr := godotenv.Load() //Load .env file and error check
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize a new connection object to database
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               "bandscape",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func GetDB() *sql.DB {
	return db
}

func GetPosts() ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Datetime, &post.Userid)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(id string) (*Post, error) {
	post := new(Post)
	err := db.QueryRow("SELECT * FROM post WHERE postid = ?", id).Scan(&post.ID, &post.Title, &post.Body, &post.Datetime, &post.Userid)
	if err != nil {
		return nil, errors.New("No post with that id")
	}
	return post, nil
}

func UpdatePost(updatedPost Post) error {
	id := updatedPost.ID
	title := updatedPost.Title
	body := updatedPost.Body
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
	res, err := db.Exec("INSERT INTO post (title, body, userid) VALUES (?, ?, ?)", title, body, userid)
	if err != nil {
		return errors.New(err.Error())
	}
	if rowsAffected, err := res.RowsAffected(); err == nil && rowsAffected == 0 {
		return errors.New("No post with that id")
	} else {
		return nil
	}
}
