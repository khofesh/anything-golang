package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Post ...
type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// Author ...
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Comment ...
type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		ID:      1,
		Content: "hello world",
		Author: Author{
			ID:   2,
			Name: "sau sheong",
		},
		Comments: []Comment{
			{
				ID:      3,
				Content: "have a great day",
				Author:  "Adam",
			},
			{
				ID:      4,
				Content: "how are you today?",
				Author:  "betty",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("post.json", output, 0644)
	if err != nil {
		fmt.Println("error writing JSON to file:", err)
		return
	}
}
