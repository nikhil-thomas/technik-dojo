package area

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		result float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, result: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, result: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Altitude: 6}, result: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.result {
				t.Errorf("%#v got %.2f want %.2f\n", tt.shape, got, tt.result)
			}
		})
	}
}
