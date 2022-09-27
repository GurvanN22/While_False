package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	bitbybit := []byte(s)
	for i := 0; i != len(s); i++ {
		a := bitbybit[i]
		z01.PrintRune(rune(a))

	}

}
