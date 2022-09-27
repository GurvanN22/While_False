package main

import "github.com/01-edu/z01"

func main() {

	for celt := 'a'; celt <= 'z'; celt++ {
		z01.PrintRune(celt)
	}

	z01.PrintRune('\n')
}
