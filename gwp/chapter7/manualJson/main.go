package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("error creating JSON file:", err)
		return
	}

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("error encoding JSON to file:", err)
		return
	}
}
