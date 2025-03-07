package iteration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("should return word itself when no count is provided", func(t *testing.T) {
		actual := Repeat("a", 0)
		expected := "a"
		assert.Equal(t, expected, actual)
	})

	t.Run("should repeat word count times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"
		assert.Equal(t, expected, actual)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkImproveRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImprovedRepeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("chety", 2)
	fmt.Println(result)
	//Output: chetychety
}
