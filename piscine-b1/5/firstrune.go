package piscine

func FirstRune(s string) rune {
	var h rune
	for _, v := range s {
		h = v
		return v
	}
	return h
}
