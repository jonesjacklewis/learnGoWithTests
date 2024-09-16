package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}

	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		Shape Shape
		Want  float64
	}{
		{Shape: Rectangle{Width: 12, Height: 6}, Want: 72},
		{Shape: Circle{Radius: 10}, Want: 314.1592653589793},
		{Shape: Triangle{Base: 12, Height: 6}, Want: 36},
	}

	for _, tt := range areaTests {
		got := tt.Shape.Area()

		if got != tt.Want {
			t.Errorf("%#v - got %.2f want %.2f", tt.Shape, got, tt.Want)
		}
	}
}
