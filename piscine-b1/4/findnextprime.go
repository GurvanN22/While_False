package piscine

func FindNextPrime(nb int) int {
	if nb > 1 {
		for i := nb; i < nb*2; i++ {
			if IsPrime(i) == true {
				return i
			}
		}
	}
	return 2
}
