package structs

import "math"

type Shape interface {
	Area() float64
}

// Rectangle func
type Rectangle struct {
	width  float64
	height float64
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return (t.base * t.height) * 0.5
}

// Arena func from Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

// Area func from Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter func
func perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
