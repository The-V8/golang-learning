package structs

import (
	"testing"
)

func assert(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	assert(t, got, want)
}

func TestArea(t *testing.T) {

	t.Run("Calculate area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		assert(t, got, want)
	})

	t.Run("Calculate area of a circle", func(t *testing.T) {
		cicle := Circle{10.0}
		got := cicle.Area()
		want := 314.16
		assert(t, got, want)
	})
}
