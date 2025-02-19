package main

import (
	"fmt"
	"time"
)

func main() {
	rateLimiter := time.Tick(500 * time.Millisecond)

	for i := 1; i <= 5; i++ {
		<-rateLimiter
		fmt.Println("Processed request", i, "at", time.Now())
	}
}
