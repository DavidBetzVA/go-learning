package main

import (
	"errors"
	"fmt"
)

// Divide performs division and returns an error if divisor is zero
// This function is exported (public) because it starts with an uppercase letter
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Another public function (exported, can be accessed from other packages)
func PublicFunction() {
	fmt.Println("I am a public function!")
}

// Private function (not exported, accessible only within this package)
func privateFunction() {
	fmt.Println("I am a private function!")
}

func main() {
	result, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	PublicFunction()
	privateFunction()
}
