package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Post ...
type Post struct {
	XMLName  xml.Name  `xml:"post"`
	ID       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	XML      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

// Comment ...
type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

// Author ...
type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se)
			}
		}
	}

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
