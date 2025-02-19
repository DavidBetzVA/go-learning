package main

import (
	"fmt"
)

func main() {
	str := "Go ❤️ Runes"
	runes := []rune(str)

	for i, r := range runes {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}
}
