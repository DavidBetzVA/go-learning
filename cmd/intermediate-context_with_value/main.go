package main

import (
	"context"
	"fmt"
	"time"
)

// Context key type to avoid collisions
type contextKey string

const userKey contextKey = "userID"

// Simulated function that retrieves user data using context
func processRequest(ctx context.Context) {
	userID, ok := ctx.Value(userKey).(string)
	if !ok {
		fmt.Println("No user ID found in context")
		return
	}

	fmt.Println("Processing request for user:", userID)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Request processed successfully for user:", userID)
	case <-ctx.Done():
		fmt.Println("Request canceled for user:", userID)
	}
}

func main() {
	// Create a base context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Store user ID in context
	ctx = context.WithValue(ctx, userKey, "12345")

	// Process the request with context
	processRequest(ctx)
}
