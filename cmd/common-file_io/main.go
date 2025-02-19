package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the file name
	fileName := "example.txt"

	// Write to a file
	content := []byte("Hello, Go developers!")
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully")

	// Read from a file
	readContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("File content:", string(readContent))

	// Clean up by removing the file
	err = os.Remove(fileName)
	if err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
	fmt.Println("File removed successfully")
}
