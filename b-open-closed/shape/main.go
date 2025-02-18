package main

import "fmt"

// This example demonstrates the Open/Closed Principle (OCP).
// Software entities (classes, modules, functions, etc.) should be open for
// extension, but closed for modification.

// Shape defines the common behavior for all shapes
type Shape interface {
	Area() float64
}

// Rectangle implements the Shape interface (implicitly)
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method as required by the interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements the Shape interface
type Circle struct {
	Radius float64
}

// pi is approximately...
const pi = 3.14159

// Area method as required by the interface
func (c Circle) Area() float64 {
	return pi * c.Radius * c.Radius
}

// Triangle implements the Shape interface
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// AreaCalculator calculates the total area of a slice of Shapes.
// This class is open for extension (can handle new shape types)
// but closed for modification (adding new shapes doesn't require changes)
type AreaCalculator struct{}

func (ac AreaCalculator) CalculateTotalArea(shapes []Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 5, Height: 10},
		Circle{Radius: 7},
		Triangle{Base: 4, Height: 6},
		// Add a new shape types without modifying AreaCalculator...
	}

	calculator := AreaCalculator{}
	totalArea := calculator.CalculateTotalArea(shapes)
	fmt.Printf("Total area: %.2f\n", totalArea)
}
