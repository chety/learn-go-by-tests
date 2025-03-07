package interfaces

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Perimeter()
		if got != want {
			t.Errorf("%#v got %g, want %g", s, got, want)
		}
	}

	//table driven tests
	tests := []struct {
		name string
		s    Shape
		want float64
	}{
		{"Rectangle", Rectangle{Width: 10, Height: 20}, 60.0},
		{"Circle", Circle{Radius: 5}, 31.41592653589793},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkPerimeter(t, test.s, test.want)
		})
	}

	//t.Run("should return the perimeter of a rectangle", func(t *testing.T) {
	//	rect := Rectangle{Width: 10, Height: 20}
	//	checkPerimeter(t, rect, 60.0)
	//})

	//t.Run("should return the perimeter of a circle", func(t *testing.T) {
	//	circle := Circle{Radius: 5}
	//	checkPerimeter(t, circle, 31.41592653589793)
	//})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("%#v got %g, want %g", s, got, want)
		}
	}

	tests := []struct {
		name string
		s    Shape
		want float64
	}{
		{"Rectangle", Rectangle{Width: 10, Height: 20}, 200.0},
		{"Circle", Circle{Radius: 10}, 314.1592653589793},
		{name: "Triangle", s: Triangle{Base: 18, Height: 6}, want: 54.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.s, test.want)
		})
	}

	//t.Run("area of a rectangle", func(t *testing.T) {
	//	rect := Rectangle{Width: 18, Height: 20}
	//	checkArea(t, rect, 360)
	//})

	//t.Run("area of a circle", func(t *testing.T) {
	//	circle := Circle{Radius: 10}
	//	checkArea(t, circle, 314.1592653589793)
	//})
}
