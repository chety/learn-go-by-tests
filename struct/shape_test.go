package shape

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("should return right perimeter for the Rectangle", func(t *testing.T) {
		rectangle := Rectangle{Width: 15.00, Height: 25.00}
		actual := rectangle.Perimeter()
		expected := 80.00

		assert.Equal(t, expected, actual, fmt.Sprintf("expected %.2f, got %.2f", expected, actual))
	})
}

func TestArea(t *testing.T) {

	type Test struct {
		name     string
		shape    Shape
		expected float64
	}

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		assert.Equal(t, expected, actual)
	}

	tests := []Test{
		{
			name:     "Rectangle area should return right value",
			shape:    Rectangle{Width: 10, Height: 25},
			expected: 250.00,
		},
		{
			name:     "Circle area should return right value",
			shape:    Circle{Radius: 10},
			expected: 314.00,
		},
		{
			name:     "Triangle area should return right value",
			shape:    Triangle{Base: 10, Height: 20},
			expected: 100.00,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.expected)
		})
	}
}
