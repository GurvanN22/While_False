package piscine

func IterativeFactorial(nb int) (resultat int) {
	resultat = 1
	if nb > 20 || nb < 0 {
		return
	} else {
		for i := nb; i > 0; i-- {
			resultat *= i
		}
		return
	}
}
