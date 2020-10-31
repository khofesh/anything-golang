package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error reading JSON data:", err)
		return
	}

	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
