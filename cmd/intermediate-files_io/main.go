package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("example.txt")
	defer func() {
		_ = file.Close()
	}()
	_, _ = file.WriteString("Hello, Go!")
	fmt.Println("File written successfully.")
}
