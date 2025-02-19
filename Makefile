phony: cmd

collections:
	go run ./cmd/collections

control:
	go run ./cmd/control

datatypes:
	go run ./cmd/datatypes

functions:
	go run ./cmd/functions
	
io:
	go run ./cmd/io

packages:
	go run ./cmd/packages

pointers:
	go run ./cmd/pointers

structs:
	go run ./cmd/structs

variables:
	go run ./cmd/variables
	
slice:
	go run ./cmd/slice
