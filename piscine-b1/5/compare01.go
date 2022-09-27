package piscine

func Compare(a, b string) int {
	len1 := 0
	len2 := 0
	for i, _ := range a {
		len1 = i + 1
	}
	for i, _ := range b {
		len2 = i + 1
	}
	if len1 == len2 {
		return 0
	}
	if len1 > len2 {
		return -1
	}
	if len1 < len2 {
		return 1
	}
	return 666
}
