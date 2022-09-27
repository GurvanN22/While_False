package main

import "github.com/01-edu/z01"

func main() {

	for celt := '0'; celt <= '9'; celt++ {
		z01.PrintRune(celt)
	}

	z01.PrintRune('\n')
}
