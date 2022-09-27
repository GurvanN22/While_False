package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}
	if estpremier(nb) == true {
		return 0
	}
	for i := nb; i > 0; i-- {
		if i*i == nb {
			return i
		}
	}
	return 0
}

func estpremier(nb int) (res bool) {
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			res = false
			return res
		}
	}
	res = true
	return res
}
