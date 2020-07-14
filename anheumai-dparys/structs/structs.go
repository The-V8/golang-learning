package structs

// Perimeter calculates the perimeter of a rectangle
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area calculates the area of a rectangle
func Area(width, height float64) float64 {
	return width * height
}
