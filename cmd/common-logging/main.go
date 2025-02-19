package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Basic logging
	log.Println("This is a basic log message")

	// Logging with different log levels
	log.Print("This is a Print log message")
	log.Printf("This is a %s log message", "Printf")
	log.Fatal("This is a Fatal log message") // This will terminate the program
	// log.Panic("This is a Panic log message") // This will terminate the program with a panic

	// Custom logger
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()
	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("This is a custom log message")

	// Custom logger with different log levels
	infoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger.Println("This is an info log message")
	warnLogger.Println("This is a warning log message")
	errorLogger.Println("This is an error log message")

	// Logging to multiple destinations
	multiLogger := log.New(io.MultiWriter(file, os.Stdout), "MULTI: ", log.Ldate|log.Ltime|log.Lshortfile)
	multiLogger.Println("This is a log message to multiple destinations")
}
