package main

import (
	"embed"
	"fmt"
	"log"
)

// Embedding a single file.

//go:embed example.txt
var singleFile string

// Embedding multiple files.
// The "embed.FS" type allows us to work with multiple files as a filesystem.

//go:embed files/*
var multipleFiles embed.FS

func main() {
	// Example: Reading the content of a single embedded file.
	fmt.Println("Contents of example.txt:")
	fmt.Println(singleFile)

	// Example: Reading multiple embedded files from the "files" directory.
	fmt.Println("\nReading multiple embedded files:")
	readEmbeddedFiles()
}

// readEmbeddedFiles iterates over embedded files and prints their content.
func readEmbeddedFiles() {
	// List of files we expect in the embedded filesystem.
	fileNames := []string{"files/sample1.txt", "files/sample2.txt"}

	for _, fileName := range fileNames {
		data, err := multipleFiles.ReadFile(fileName)
		if err != nil {
			log.Printf("Failed to read %s: %v\n", fileName, err)
			continue
		}
		fmt.Printf("Contents of %s:\n%s\n", fileName, string(data))
	}
}
