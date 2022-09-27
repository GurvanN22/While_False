package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		n = 9223372036854775807
		n += 1
	}
	longeur := cont(n) - 1
	utf8 := 48
	if n > 0 {
		for i := 0; i != longeur; i++ {
			var diviseur int = puissance(10, longeur-i)
			aafficher := n / diviseur
			n -= aafficher * (diviseur)
			z01.PrintRune(rune(aafficher + utf8))
		}
		z01.PrintRune(rune(n + utf8))
	} else {
		z01.PrintRune('-')
		n = n * -1
		for i := 0; i != longeur; i++ {
			var diviseur int = puissance(10, longeur-i)
			aafficher := n / diviseur
			n -= aafficher * (diviseur)
			z01.PrintRune(rune(aafficher + utf8))
		}
		z01.PrintRune(rune(n + utf8))
	}
}

// Cette fonction sert de len pour une variable int
func cont(k int) (i int) {
	for k != 0 {
		k = k / 10
		i++
	}
	return
}

func puissance(n int, k int) (o int) {
	k -= 1
	val0 := n
	for i := 0; i != k; i++ {
		n = n * val0
	}
	o = n
	return
}
