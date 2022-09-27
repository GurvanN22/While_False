package piscine

func StrRev(s string) (reversed string) {
	for i := len(s) - 1; i != -1; i-- {
		reversed = reversed + string(s[i])
	}
	return
}
