# Intermediate Plugin System

This example demonstrates how to use a plugin in Go. Follow the instructions below to get started.

## Building the Plugin

Build the plugin:

```sh
go build -buildmode=plugin -o plugin_example.so plugin_example.go
```

## Usage

To run the main program, use the following command:

```sh
go run main.go
```
