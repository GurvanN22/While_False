package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var panic []int
		return panic
	}
	res := make([]int, max-min)
	index := 0
	for i := min; i < max; i++ {
		res[index] = i
		index++
	}
	return res
}
