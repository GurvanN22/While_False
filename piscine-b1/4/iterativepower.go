package piscine

func IterativePower(n int, k int) (o int) {
	if k < 0 {
		return 0
	}
	if k == 0 {
		return n
	}
	k -= 1
	val0 := n
	for i := 0; i != k; i++ {
		n = n * val0
	}
	o = n
	return
}
