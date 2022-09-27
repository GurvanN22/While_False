package piscine

func StrLen(s string) (compteur int) {
	compteur = 0
	for _, v := range s {
		compteur += 1
		v += v
	}
	return
}
