package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi("560"))
}

func BasicAtoi(s string) (resultat int) {
	longeur := len(s) - 1
	for i, v := range s {
		multiple := puissance(10, longeur-i)
		resultat += (int(v) - 48) * multiple
	}
	return
}

func puissance(n int, k int) (o int) {
	k -= 1
	val0 := n
	for i := 0; i != k; i++ {
		n = n * val0
	}
	o = n
	return
}
