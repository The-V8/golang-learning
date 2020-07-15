package structsinterfaces

import "math"

// Shape can calculate its area
type Shape interface {
	Area() float64
}

// Rectangle definition
type Rectangle struct {
	Width  float64
	Height float64
}

// Area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle definition
type Circle struct {
	Radius float64
}

// Area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle definition
type Triangle struct {
	Base   float64
	Height float64
}

// Area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

//Perimeter returns the perimeter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
