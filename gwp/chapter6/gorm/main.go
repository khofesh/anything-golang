package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Post ...
type Post struct {
	ID        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

// Comment ...
type Comment struct {
	ID        int
	Content   string
	Author    string `sql:"not null"`
	PostID    int    `sql:"index"`
	CreatedAt time.Time
}

// Db ...
var Db gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "hello world", Author: "sau sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "good post", Author: "joe"}
	Db.Modal(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "sau sheong").First(&readPost)
	var comments []Comment
	Db.Modal(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
