package main

import (
	"os"
	"text/template"
)

const tmpl = `Hello, {{.Name}}! You are {{.Age}} years old.`

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	t, _ := template.New("greeting").Parse(tmpl)
	_ = t.Execute(os.Stdout, p)
}
