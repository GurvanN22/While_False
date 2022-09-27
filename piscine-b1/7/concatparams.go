package piscine

func ConcatParams(args []string) string {
	var resultat string
	lenght := len(args) - 1
	for i := range args {
		resultat += args[i]
		if i != lenght {
			resultat += "\n"
		}

	}
	return resultat
}
