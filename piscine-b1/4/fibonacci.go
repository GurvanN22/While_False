package piscine

func Fibonacci(index int) int {
	result := 0
	if index < 0 {
		return -1
	} else if index == 0 {
		result = 0
	} else if index == 1 {
		result = 1
	} else {
		result = Fibonacci(index-1) + Fibonacci(index-2)
	}
	return result
}
