package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() { //Liste et utilisation de os.Args
	var arrayBool []bool
	arguments := os.Args[1:]
	for _, value := range arguments {
		valueInt, _ := strconv.Atoi(value)
		arrayBool = append(arrayBool, odd(valueInt))
	}
	fmt.Println(arguments)
	fmt.Println(arrayBool)

	//------------------------------------
	argumentsstring := os.Args[1]
	var liste []string
	for _, v := range argumentsstring {
		liste = append(liste, string(v))
	}
	fmt.Println(liste)
	nbEspace := 0
	for _, value := range liste {
		if value == " " {
			nbEspace++
		}
	}
	fmt.Println(nbEspace + 1)
}

func odd(value int) bool {
	if value%2 == 0 {
		return false
	} else {
		return true
	}
}
