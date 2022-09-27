package piscine

func IsPrime(nb int) bool {
	if nb > 1 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
