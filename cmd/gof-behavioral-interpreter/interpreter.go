package main

import "fmt"

type Expression interface {
	Interpret() string
}

type TerminalExpression struct {
	value string
}

func (t *TerminalExpression) Interpret() string {
	return t.value
}

type NonTerminalExpression struct {
	left, right Expression
}

func (n *NonTerminalExpression) Interpret() string {
	return n.left.Interpret() + " " + n.right.Interpret()
}

func main() {
	left := &TerminalExpression{"Hello"}
	right := &TerminalExpression{"World"}
	expression := &NonTerminalExpression{left, right}

	fmt.Println(expression.Interpret())
}
