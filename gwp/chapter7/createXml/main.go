package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("error marshalling to XML:", err)
		return
	}

	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("error writing XML to file:", err)
		return
	}
}
