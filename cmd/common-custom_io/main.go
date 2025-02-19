package main

import (
	"fmt"
	"io"
	"strings"
)

// Custom Reader
type MyReader struct {
	data string
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	if len(r.data) == 0 {
		return 0, io.EOF
	}
	n = copy(p, r.data)
	r.data = r.data[n:]
	return n, nil
}

// Custom Writer
type MyWriter struct{}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println("Custom Writer Output:", string(p))
	return len(p), nil
}

func main() {
	// Using custom Reader
	reader := &MyReader{data: "Hello, Custom Reader!"}
	buf := make([]byte, 8)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Print(string(buf[:n]))
	}
	fmt.Println()

	// Using custom Writer
	writer := &MyWriter{}
	_, _ = io.Copy(writer, strings.NewReader("Hello, Custom Writer!"))
}
