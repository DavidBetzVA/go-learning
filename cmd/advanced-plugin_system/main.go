package main

import (
	"fmt"
	"plugin"
)

func main() {
	// Load the plugin
	p, err := plugin.Open("plugin_example.so")
	if err != nil {
		fmt.Println("Error loading plugin:", err)
		return
	}

	// Lookup the exported symbol
	sym, err := p.Lookup("Hello")
	if err != nil {
		fmt.Println("Error finding symbol:", err)
		return
	}

	// Assert function type and execute
	helloFunc, ok := sym.(func() string)
	if !ok {
		fmt.Println("Invalid function signature")
		return
	}

	fmt.Println(helloFunc())
}
