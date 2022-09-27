package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	lenght := len(arguments) - 1
	for i := lenght; i > 0; i-- {
		for _, v := range arguments[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
