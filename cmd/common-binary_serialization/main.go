// Package main demonstrates advanced binary serialization using the encoding/gob package.
// This example shows how to encode and decode a custom data structure.
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Data is a custom data structure that we will serialize and deserialize.
type Data struct {
	Name  string
	Value int
}

func main() {
	// Create a buffer to hold the encoded data.
	var buf bytes.Buffer

	// Create an encoder that writes to the buffer.
	enc := gob.NewEncoder(&buf)

	// Create a decoder that reads from the buffer.
	dec := gob.NewDecoder(&buf)

	// Create an instance of the Data structure to be serialized.
	original := Data{Name: "Example", Value: 42}

	// Encode the original data structure into the buffer.
	_ = enc.Encode(original)

	// Create a variable to hold the decoded data.
	var decoded Data

	// Decode the data from the buffer into the decoded variable.
	_ = dec.Decode(&decoded)

	// Print the decoded data to verify it matches the original.
	fmt.Println("Decoded:", decoded)
}
