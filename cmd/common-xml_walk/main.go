package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Sample XML structure
const xmlData = `
<library>
	<book>
		<title>Go Programming</title>
		<author>John Doe</author>
	</book>
	<book>
		<title>Advanced Go</title>
		<author>Jane Smith</author>
	</book>
</library>`

// Define struct for XML parsing
type Library struct {
	Books []Book `xml:"book"`
}

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}

func main() {
	var lib Library

	// Decode XML
	err := xml.Unmarshal([]byte(xmlData), &lib)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		os.Exit(1)
	}

	// Walk through leaf elements
	for _, book := range lib.Books {
		fmt.Println("Book Title:", book.Title)
		fmt.Println("Author:", book.Author)
		fmt.Println("---")
	}
}
