package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// Example GET request
	getExample()

	// Example POST request
	postExample()

	// Example PUT request
	putExample()

	// Example DELETE request
	deleteExample()

	// Example POST Form request
	postFormExample()
}

func getExample() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("GET Response:", string(body))
}

func postExample() {
	jsonData := []byte(`{"title":"foo","body":"bar","userId":1}`)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("POST Response:", string(body))
}

func putExample() {
	client := &http.Client{}
	jsonData := []byte(`{"id":1,"title":"foo","body":"bar","userId":1}`)
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("PUT Response:", string(body))
}

func deleteExample() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("DELETE Response:", string(body))
}

func postFormExample() {
	formData := url.Values{
		"name":  {"John Doe"},
		"email": {"john.doe@example.com"},
	}

	resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", formData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("POST Form Response:", string(body))
}
