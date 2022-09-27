package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	utf8 := 48
	for dizaine2 := 0; dizaine2 < 10; dizaine2++ {
		for unite2 := 0; unite2 < 10; unite2++ {
			for dizaine1 := 0; dizaine1 < 10; dizaine1++ {
				for unite1 := 0; unite1 < 10; unite1++ {

					if dizaine2 != dizaine1 && unite2 != unite1 {
						z01.PrintRune(rune(dizaine2 + utf8))
						z01.PrintRune(rune(unite2 + utf8))
						z01.PrintRune(' ')
						z01.PrintRune(rune(dizaine1 + utf8))
						z01.PrintRune(rune(unite1 + utf8))
						z01.PrintRune(',')
					}
				}
			}
		}
	}
}
