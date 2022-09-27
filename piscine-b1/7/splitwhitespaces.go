package piscine

func SplitWhiteSpaces(s string) []string {
	var resultat []string
	var res string
	for i, value := range s {
		if value > 32 || i+1 == len(s) && value > 32 {
			res += string(value)
			if i+1 == len(s) {
				resultat = append(resultat, res)
			}
		} else if res != "" {
			resultat = append(resultat, res)
			res = ""
		}
	}
	return resultat
}
