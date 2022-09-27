package piscine

func Capitalize(s string) string {
	bytes := []byte(s)
	for i, v := range s {
		if int(v) >= 65 && int(v) <= 90 || int(v) >= 97 && int(v) <= 122 {
			if int(v) >= 65 && int(v) <= 90 {
				if i != 0 {
					if bytes[i-1] >= 65 && bytes[i-1] <= 90 || bytes[i-1] >= 97 && bytes[i-1] <= 122 || bytes[i-1] >= 48 && bytes[i-1] <= 57 {
						bytes[i] += 32
					}
				}
			} else {
				if i == 0 {
					bytes[i] -= 32
				} else {
					if bytes[i-1] >= 65 && bytes[i-1] <= 90 || bytes[i-1] >= 97 && bytes[i-1] <= 122 || bytes[i-1] >= 48 && bytes[i-1] <= 57 {
					} else {
						bytes[i] -= 32
					}
				}
			}
		}
	}
	return string(bytes)
}
