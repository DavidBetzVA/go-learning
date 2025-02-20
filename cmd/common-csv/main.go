package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
	}

	for _, record := range data {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	fmt.Println("CSV file written successfully")
}
