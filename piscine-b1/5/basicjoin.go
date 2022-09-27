package piscine

func BasicJoin(elems []string) string {
	var resultat string
	for _, v := range elems {
		resultat += v
	}
	return resultat
}
