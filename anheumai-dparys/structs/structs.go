package structs

import (
	"math"
)

// Rectangle declares a rectangle (comments are just here to please the linter)
type Rectangle struct {
	width  float64
	height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Circle declares a circle (what a surprise)
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a rectangle
func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

// Area calculates the area of a rectangle
func Area(r Rectangle) float64 {
	return r.width * r.height
}
