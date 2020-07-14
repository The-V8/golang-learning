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
	got := Perimeter(10.0, 10.0)
	want := 40.0
	assert(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0
	assert(t, got, want)
}
