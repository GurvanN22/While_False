package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		bytes := []byte(Intostr(n))
		len := len(bytes)
		for i := 0; i != len-1; {
			if int(bytes[i]) > int(bytes[i+1]) {
				trans := bytes[i]
				bytes[i] = bytes[i+1]
				bytes[i+1] = trans
				i = 0
			} else {
				i++
			}
		}
		for i := range bytes {
			z01.PrintRune(rune(int(bytes[i])))
		}
	}
}

// Cette fonction transforme un int en string
func Intostr(n int) string {
	utf8 := 48
	longeur := cont(n) - 1
	var str string
	if n > 0 {
		for i := 0; i != longeur; i++ {
			var diviseur int = puissance(10, longeur-i)
			aafficher := n / diviseur
			n -= aafficher * (diviseur)
			str += string(rune(aafficher + utf8))
		}
		str += string(rune(n + utf8))
	}
	return str
}

// Cette fonction sert de len pour une variable int
func cont(k int) (i int) {
	for k != 0 {
		k = k / 10
		i++
	}
	return
}

// Cette fonction mette le premier paramètre à la puissance du second paramètre
func puissance(n int, k int) (o int) {
	k -= 1
	val0 := n
	for i := 0; i != k; i++ {
		n = n * val0
	}
	o = n
	return
}
