package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	utf8 := 48
	for centaine := 0; centaine < 10; centaine++ {
		for dizaine := 0; dizaine < 10; dizaine++ {
			for unite := 0; unite < 10; unite++ {
				if centaine < dizaine && dizaine < unite {
					if centaine == 7 {
						z01.PrintRune(rune(centaine + utf8))
						z01.PrintRune(rune(dizaine + utf8))
						z01.PrintRune(rune(unite + utf8))
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(rune(centaine + utf8))
						z01.PrintRune(rune(dizaine + utf8))
						z01.PrintRune(rune(unite + utf8))
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
