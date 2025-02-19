package main

import (
	"fmt"
	"runtime"
	"time"
)

func allocate() {
	data := make([]byte, 1e6)
	_ = data
}

func main() {
	fmt.Println("Memory before allocation:", runtime.MemStats{}.Alloc)
	allocate()
	time.Sleep(time.Second)
	fmt.Println("Memory after allocation:", runtime.MemStats{}.Alloc)
}
