package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Post ...
type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

// Author ...
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: "hello world",
		Author: Author{
			ID:   "2",
			Name: "sau sheong",
		},
	}

	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println("error creating XML file:", err)
		return
	}

	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("error encoding XML to file:", err)
		return
	}
}
