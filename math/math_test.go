package mathOperations

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("should return 0 when no numbers are provided", func(t *testing.T) {
		actual := Add()
		expected := 0

		assert.Equal(t, expected, actual)
	})

	t.Run("should add multiple numbers", func(t *testing.T) {
		actual := Add(1, 2, 3, 4, 5, 6, 7)
		expected := 28

		assert.Equal(t, expected, actual)
	})
}

func ExampleAdd() {
	result := Add(1, 2, 3, 4, 5)
	fmt.Println(result)
	// Output: 15
}

func TestSum(t *testing.T) {

	t.Run("should return sum of the numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		assert.Equal(t, expected, got)
	})

	t.Run("should return 0 if no numbers are provided", func(t *testing.T) {
		var numbers []int
		got := Sum(numbers)
		expected := 0

		assert.Equal(t, expected, got)
	})
}

func ExampleSum() {
	result := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(result)
	// Output: 15
}
