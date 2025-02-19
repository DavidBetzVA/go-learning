package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("APP_MODE", "development")
	mode := os.Getenv("APP_MODE")

	fmt.Println("Application Mode:", mode)
}
