package piscine

func IsUpper(s string) bool {
	for _, v := range s {
		if v < 65 || v > 90 {
			return false
		}
	}
	return true
}
