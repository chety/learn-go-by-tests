package mathOperations

// Add multiple numbers. Here we are using named parameters(naked return)
func Add(num ...int) (sum int) {
	for _, n := range num {
		sum += n
	}
	return
}

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numSlices ...[]int) []int {
	var sumSlice = make([]int, 0)
	for _, slice := range numSlices {
		sumSlice = append(sumSlice, Sum(slice))
	}
	return sumSlice
}

func SumAllTails(numSlices ...[]int) []int {
	var sumSlice = make([]int, 0)
	for _, slice := range numSlices {
		sum := 0
		if len(slice) > 0 {
			sum = Sum(slice[1:])
		}
		sumSlice = append(sumSlice, sum)
	}
	return sumSlice
}
