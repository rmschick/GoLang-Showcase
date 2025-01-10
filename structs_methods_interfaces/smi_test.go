package structs_methods_interfaces_test

import (
	"testing"

	smi "GoLang-Showcase/structs_methods_interfaces"
)

func TestPerimeter(t *testing.T) {
	rectangle := smi.Rectangle{
		Width:  10.0,
		Height: 10.0,
	}

	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape smi.Shape
		want  float64
	}{
		{name: "Rectangle", shape: &smi.Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{name: "Circle", shape: &smi.Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: &smi.Triangle{Base: 12.0, Height: 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%s: got %g want %g", tt.name, got, tt.want)
			}
		})
	}
}
