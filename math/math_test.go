package mathOperations

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestSumAll(t *testing.T) {
	t.Run("should return {0,0,0} if 3 empty slices are provided", func(t *testing.T) {
		actual := SumAll([]int{}, []int{}, []int{})
		expected := []int{0, 0, 0}
		assert.Equal(t, expected, actual)
	})

	t.Run("should sum of all numbers of each slice then return a single slice", func(t *testing.T) {
		actual := SumAll([]int{1, 1, 1})
		expected := []int{3}
		assert.Equal(t, expected, actual)
	})

	t.Run("should sum of all numbers of each slice then return a single slice", func(t *testing.T) {
		actual := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
		expected := []int{6, 15}
		assert.Equal(t, expected, actual)
	})

}

func TestSumAllTails(t *testing.T) {

	//In this test we use 3 different variations for checking the result. Assert package provides much more
	//usable methods for testing. We can stick with this approach.
	//1- creating custom checkSlice function
	//2- Using assert package with custom message
	//3- Using assert package without custom message

	checkSlice := func(t testing.TB, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	}

	t.Run("should safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{})
		expected := []int{0, 0}
		checkSlice(t, expected, actual)
	})

	t.Run("should safely sum when some inputs are empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{1, 2, 3, 4}, []int{5, 6})
		expected := []int{0, 9, 6}
		assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v, but got %v", expected, actual))
	})

	t.Run("should return some of tails of given slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3, 4}, []int{5, 6})
		expected := []int{9, 6}
		assert.Equal(t, expected, actual)
	})

}
