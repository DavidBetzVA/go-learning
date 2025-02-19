package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder

	sb.WriteString("Hello")
	sb.WriteString(", ")
	sb.WriteString("World!")

	fmt.Println(sb.String()) // Output: Hello, World!
}
