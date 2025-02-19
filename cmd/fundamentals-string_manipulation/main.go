package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Go is Awesome!  "
	fmt.Println("Original:", str)
	fmt.Println("Trimmed:", strings.TrimSpace(str))
	fmt.Println("Uppercase:", strings.ToUpper(str))
	fmt.Println("Lowercase:", strings.ToLower(str))
	fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))
	fmt.Println("Replace 'Awesome' with 'Great':", strings.ReplaceAll(str, "Awesome", "Great"))
}
