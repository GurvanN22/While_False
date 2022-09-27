package piscine

func ToLower(s string) string {
	bytes := []byte(s)
	len := len(bytes)
	for i := 0; i < len; i++ {
		if bytes[i] >= 65 && bytes[i] <= 90 {
			bytes[i] += 32
		}
	}
	return string(bytes)
}
