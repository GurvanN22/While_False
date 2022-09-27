package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	compte := 0
	ind := -1
	sbyte := []byte(s)
	toFindbyte := []byte(toFind)
	for i := 0; i < len(sbyte); i++ {
		for y := 0 + compte; y < len(toFindbyte); y++ {
			if sbyte[i] == toFindbyte[y] {
				compte += 1
				if ind == -1 {
					ind = i
				}
				if compte == len(toFindbyte) {
					return ind
				}
				break
			} else {
				compte = 0
				break
			}
		}
	}
	return -1
}
