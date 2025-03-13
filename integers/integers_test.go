package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("adds two integers", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6

		if sum != expected {
			t.Errorf("sum %d, expected %d", sum, expected)
		}
	})
}

// If we run tests, this function will be executed also
// If we remove the comment in the function, the function will not be executed
func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
