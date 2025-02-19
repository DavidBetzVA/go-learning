package main

import (
	"fmt"
	"unicode"
)

func main() {
	r := 'G'

	if unicode.IsLetter(r) {
		fmt.Println(string(r), "is a letter")
	}
	if unicode.IsUpper(r) {
		fmt.Println(string(r), "is uppercase")
	}
}
