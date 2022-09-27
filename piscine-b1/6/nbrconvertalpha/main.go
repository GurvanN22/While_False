package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	supp := 0
	Upper := 0
	lenght := len(arguments) - 1
	if lenght == 0 {
	} else {
		if arguments[1] == "--upper" {
			Upper = 1
			supp = 1
		}
		for i := 1 + supp; i <= lenght; i++ {
			resbytes := 0
			leng := len(arguments[i]) - 1
			for index, v := range arguments[i] {
				resbytes += (int(v) - 48) * puissance(10, leng-index)
			}
			resbytes += 96
			if resbytes >= 97 && resbytes <= 122 {
				if Upper == 1 {
					z01.PrintRune(rune(resbytes - 32))
				} else {
					z01.PrintRune(rune(resbytes))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func puissance(n int, k int) (o int) {
	if k == 0 {
		return 1
	}
	k -= 1
	val0 := n
	for i := 0; i != k; i++ {
		n = n * val0
	}
	o = n
	return
}
