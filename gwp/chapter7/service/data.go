package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// DB ...
var DB *sql.DB

// Post ...
type Post struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func init() {
	var err error
	DB, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Retrieve ...
func Retrieve(id int) (post Post, err error) {
	post = Post{}
	err = DB.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.ID, &post.Content, &post.Author)
	return
}

// Create ...
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.ID)
	return
}

// Update ...
func (post *Post) Update() (err error) {
	_, err = DB.Exec("update posts set content = $2, author = $3 where id = $1", post.ID, post.Content, post.Author)
	return
}

// Delete ...
func (post *Post) Delete() (err error) {
	_, err = DB.Exec("delete from posts where id = $1", post.ID)
	return
}
