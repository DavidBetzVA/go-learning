package main

import (
	"fmt"
	"math"
)

// Shape is an interface that defines behavior for different shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements the Shape interface
type Rectangle struct {
	Width, Height float64
}

// Circle implements the Shape interface
type Circle struct {
	Radius float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method for Circle (Circumference)
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// PrintShapeInfo takes any Shape and prints its details
func PrintShapeInfo(s Shape) {
	fmt.Printf("Shape Details:\n")
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
	fmt.Println()
}

func main() {
	// Create instances of Rectangle and Circle
	rect := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 7}

	// Pass them to PrintShapeInfo (demonstrating interface usage)
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)

	// Using an interface slice to store multiple shape types
	shapes := []Shape{rect, circle, Rectangle{Width: 3, Height: 4}}

	// Loop through and process different shapes
	fmt.Println("Looping through shapes:")
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
}
