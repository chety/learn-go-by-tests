package slicesandarrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll this is more performant since only 1 memory allocation is done here
func SumAll(numbersSlices ...[]int) []int {
	sums := make([]int, len(numbersSlices))

	for i, slice := range numbersSlices {
		sums[i] = Sum(slice)
	}
	return sums
}

func SumAll2(numbersSlices ...[]int) []int {
	var sums []int
	for _, numberSlice := range numbersSlices {
		sums = append(sums, Sum(numberSlice))
	}
	return sums
}

func SumAllTails(numbersSlices ...[]int) []int {
	sums := make([]int, len(numbersSlices))
	if len(numbersSlices) == 1 {
		return sums
	}
	for i, slice := range numbersSlices {
		if len(slice) == 0 {
			continue
		}
		sums[i] = Sum(slice[1:])
	}
	return sums
}
