package main

import (
	"encoding/json"
	"fmt"
)

// Struct for JSON example
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"` // omitempty removes empty fields
}

func main() {
	// Encoding to JSON
	p := Person{Name: "Alice", Age: 30}
	jsonData, _ := json.Marshal(p)
	fmt.Println("JSON Output:", string(jsonData))

	// Decoding from JSON
	jsonInput := `{"name": "Bob", "age": 25}`
	var decoded Person
	err := json.Unmarshal([]byte(jsonInput), &decoded)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	fmt.Println("Decoded Struct:", decoded)
}
