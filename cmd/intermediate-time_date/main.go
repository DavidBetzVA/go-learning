package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)

	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted time:", formatted)

	parsedTime, err := time.Parse("2006-01-02", "2025-02-19")
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed time:", parsedTime)
	}
}
