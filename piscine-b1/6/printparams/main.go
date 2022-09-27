package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i, v := range arguments {
		if i != 0 {
			for _, s := range v {
				z01.PrintRune(rune(s))
			}
			z01.PrintRune('\n')

		}
	}
}
