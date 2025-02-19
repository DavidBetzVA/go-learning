package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Predefined input string
	input := "Hello\nWorld\nThis is a test\nEND\nAnother line"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// Demonstrate reading multiple lines
	fmt.Println("Reading multiple lines from predefined input:")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "END" {
			break
		}
		fmt.Println("Read line:", line)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}
}
