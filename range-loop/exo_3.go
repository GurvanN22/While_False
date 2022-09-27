package main

import "fmt"

//level 3/3
//Dans cette exercice vous devez crée une fonction qui crée une liste et qui lui fait apprendre chaque mot et l'afficher

func main() {

	text := "Hello world we love golang !"
	liste := SplitWhiteSpaces(text)
	fmt.Print(text)

}

func SplitWhiteSpaces(s string) []string {

}
