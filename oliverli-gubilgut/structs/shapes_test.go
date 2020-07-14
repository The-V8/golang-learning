package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Triangle", shape: Triangle{height: 6, base: 12}, hasArea: 36.0},
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
