package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Post ...
type Post struct {
	ID         int
	Content    string
	AuthorName string `db: author`
}

// Db ...
var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// GetPost ...
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

// Create ...
func (post *Post) Create() (err error) {
	err = Db.QueryRowx("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.ID)
	return
}

func main() {
	post := Post{Content: "hello world!", AuthorName: "sau sheong"}
	post.Create()
	fmt.Println(post)
}
