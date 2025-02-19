package main

import (
	"fmt"
)

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func mayFail(fail bool) error {
	if fail {
		return &MyError{"Something went wrong!"}
	}
	return nil
}

func main() {
	err := mayFail(true)
	if err != nil {
		fmt.Println("Error occurred:", err)
	}
}
