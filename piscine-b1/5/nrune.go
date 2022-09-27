package piscine

func NRune(s string, n int) rune {
	var h rune
	for i, v := range s {
		h = v
		if i+1 == n {
			return h
		}
	}
	return rune('\x00')
}
