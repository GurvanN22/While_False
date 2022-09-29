package piscine

func CountIf(f func(string) bool, tab []string) int {
	compteur := 0
	for _, v := range tab {
		if f(v) == true {
			compteur++
		}
	}
	return compteur
}
