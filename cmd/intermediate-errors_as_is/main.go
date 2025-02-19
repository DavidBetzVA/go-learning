package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

func getData(key string) error {
	if key == "" {
		return &CustomError{msg: "custom error: key is empty"}
	}
	if key == "not found" {
		return ErrNotFound
	}
	return nil
}

func main() {
	err := getData("not found")
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Handled not found error")
	}
	err = getData("")
	var customErr *CustomError
	if errors.As(err, &customErr) {
		fmt.Println("Handled custom error:", customErr)
	}
}
