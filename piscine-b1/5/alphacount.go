package piscine

func AlphaCount(s string) int {
	len := 0
	for _, v := range s {
		if int(v) >= 65 && int(v) <= 90 || int(v) >= 97 && int(v) <= 122 {
			len += 1
		}
	}
	return len
}
