package main

import "fmt"

func main() { // Les boucles range en Golang
	liste := []int{0, 1, 2, 3, 4} //On crée notre valeur liste
	for i, v := range liste {     //On initie la boucle range
		fmt.Print(i, v)
	}
	fmt.Println()
	phrase := "Hello, world!"
	for _, valeur := range phrase {
		fmt.Print(string(valeur))
	}
}
