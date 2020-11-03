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

// Decode ...
func Decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("error decoding JSON:", err)
		return
	}

	return
}

// Unmarshal ...
func Unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
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

	json.Unmarshal(jsonData, &post)
	return
}

func main() {
	_, err := Decode("post.json")
	if err != nil {
		fmt.Println("error:", err)
	}
}
