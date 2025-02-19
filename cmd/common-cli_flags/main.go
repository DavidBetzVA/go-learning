package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "User", "Your name")
	age := flag.Int("age", 30, "Your age")

	flag.Parse()

	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
