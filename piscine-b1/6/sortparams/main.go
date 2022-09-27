package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argmuments := os.Args
	lenght := len(argmuments)
	if lenght != 1 {
		for i := 1; i <= lenght-2; i++ {
			if argmuments[i] > argmuments[i+1] {
				yes := argmuments[i]
				argmuments[i] = argmuments[i+1]
				argmuments[i+1] = yes
				i = 0

			}
		}
		for i, v := range argmuments {
			if i != 0 {
				for _, y := range v {
					z01.PrintRune(y)
				}
				z01.PrintRune('\n')
			}
		}
	}
}
