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
