package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age,omitempty"`
}

func main() {
	u := User{Name: "Alice", Email: "alice@example.com", Age: 0}
	jsonData, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println(string(jsonData))
}
