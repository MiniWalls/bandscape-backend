package mydb

import (
	"database/sql"
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

func GetPosts() []Post {
	var posts []Post
	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Datetime, &post.Userid)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	return posts
}

func GetPost(id string) Post {
	var post Post
	err := db.QueryRow("SELECT * FROM post WHERE postid = ?", id).Scan(&post.ID, &post.Title, &post.Body, &post.Datetime, &post.Userid)
	if err != nil {
		log.Fatal(err)
	}
	return post
}

func UpdatePost(updatedPost Post) string {
	id := updatedPost.ID
	title := updatedPost.Title
	body := updatedPost.Body
	res, err := db.Exec("UPDATE post SET title = ?, body = ? WHERE postid = ?", title, body, id)
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected, err := res.RowsAffected(); err == nil {
		if rowsAffected == 0 {
			return "No post with that id"
		}
	}
	return "Post updated"
}

func DeletePost(id string) string {
	res, err := db.Exec("DELETE FROM post WHERE postid = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected, err := res.RowsAffected(); err == nil {
		if rowsAffected == 0 {
			return "No post with that id"
		}
	}
	return "Post deleted"
}
