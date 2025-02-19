package main

import "fmt"

type Animal struct {
	Species string
}

type Dog struct {
	Animal
	Name string
}

func main() {
	d := Dog{Animal: Animal{"Canine"}, Name: "Buddy"}
	fmt.Println(d.Name, "is a", d.Species)
}
