package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for i, v := range name {
		if i >= 2 {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
