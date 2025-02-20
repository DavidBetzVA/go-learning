package main

// import (
// 	"compress/gzip"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/rand"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// Open the input file
// 	inputFile, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening input file:", err)
// 		return
// 	}
// 	defer inputFile.Close()

// 	// Create the output file
// 	outputFile, err := os.Create("output.gz.enc")
// 	if err != nil {
// 		fmt.Println("Error creating output file:", err)
// 		return
// 	}
// 	defer outputFile.Close()

// 	// Create a gzip writer
// 	gzipWriter := gzip.NewWriter(outputFile)
// 	defer gzipWriter.Close()

// 	// Create an AES cipher
// 	key := []byte("example key 1234") // 16 bytes for AES-128
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		fmt.Println("Error creating AES cipher:", err)
// 		return
// 	}

// 	// Create a random IV
// 	iv := make([]byte, aes.BlockSize)
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		fmt.Println("Error creating IV:", err)
// 		return
// 	}

// 	// Write the IV to the output file
// 	if _, err := outputFile.Write(iv); err != nil {
// 		fmt.Println("Error writing IV to output file:", err)
// 		return
// 	}

// 	// Create a cipher stream writer
// 	stream := cipher.NewCFBEncrypter(block, iv)
// 	writer := &cipher.StreamWriter{S: stream, W: gzipWriter}

// 	// Copy the input file to the cipher stream writer
// 	if _, err := io.Copy(writer, inputFile); err != nil {
// 		fmt.Println("Error copying data:", err)
// 		return
// 	}

// 	fmt.Println("File successfully compressed and encrypted")
// }
