package slicesandarrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Sum(numbers)
		expected := 55

		if got != expected {
			t.Errorf("got %d, want %d, given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	t.Run("sum collection of EMPTY slices", func(t *testing.T) {
		got := SumAll([]int{})
		expected := []int{0}

		if slices.Compare(got, expected) != 0 {
			t.Errorf("got %d, want %d", got, expected)
		}
	})
	t.Run("sum collection of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
		expected := []int{6, 15}

		//!reflect.DeepEqual is not type-safe. If we change expected := "some string" it will still compile!
		//Instead, use the slices.Compare function
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %d, want %d", got, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if slices.Compare(got, want) != 0 {
			t.Errorf("got %d, want %d", got, want)
		}
	}

	t.Run("sum collection of tail empty slices", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		checkSum(t, got, want)
	})

	t.Run("safely sums empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSum(t, got, want)
	})

	t.Run("sum collection of tail slices with length 2", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4})
		want := []int{2, 4}

		checkSum(t, got, want)
	})

	t.Run("sum collection of tail slices with more than length 2 ", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 12}, []int{3, 4, 20})
		want := []int{14, 24}

		checkSum(t, got, want)
	})
}

var data = [][]int{
	{1, 2, 3, 4, 5},
	{6, 7, 8, 9, 10},
	{11, 12, 13, 14, 15},
	{16, 17, 18, 19, 20},
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll(data...)
	}
}

func BenchmarkSumAll1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll2(data...)
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails(data...)
	}
}
