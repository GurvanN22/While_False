package main

import "github.com/01-edu/z01"

func main() {
	n := 1
	PrintCombN(n)
}
func PrintCombN(n int) {
	var liste [10]int
	for liste[0] == 9 {
		liste[n-1] += 1
		for i := n; i == 0; i-- {
			if liste[i] == 9 {
				liste[i] = 0
				liste[i-1] += 1
			}
		}
		afficher(n, liste)
	}

}

func afficher(n int, l [10]int) {
	for i := 0; i != n; i++ {
		z01.PrintRune(rune(l[i] + 48))
	}
}
