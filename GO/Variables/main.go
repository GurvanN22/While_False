package main

import "fmt"

func main() { //Les Variables

	nomber := 0 // variable int

	fmt.Println(nomber)

	text := "Hello World!" //variable string

	fmt.Println(text)

	boll := true //variable bool

	fmt.Println(boll)

	liste := []int{1, 2, 3} // variable liste

	liste = append(liste, 4)

	fmt.Println(liste)
	//////////////////////////Fonctionnes avec la table ascii (commande console: $man ascii)
	MyFirstByte := []byte(text) //variable byte

	fmt.Println(MyFirstByte)

	MyFirstRune := 'a' //variable rune

	fmt.Println(MyFirstRune)
}
