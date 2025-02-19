package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"apple":  5,
		"banana": 3,
		"orange": 7,
	}

	fmt.Println("Map contents:", m)

	if val, exists := m["apple"]; exists {
		fmt.Println("Apple count:", val)
	}

	delete(m, "banana")
	fmt.Println("After deleting banana:", m)

	for key, val := range m {
		fmt.Printf("%s -> %d\n", key, val)
	}
}
