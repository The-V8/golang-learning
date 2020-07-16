package structs

import "math"

// Rectangle defined by width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle defined by radius
type Circle struct {
	Radius float64
}

// Area calculates area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of width by height
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// Shape defines an interface for shapes
type Shape interface {
	Area() float64
}

// Triangle defines a triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates area of triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
