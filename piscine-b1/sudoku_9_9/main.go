package main

import (
	"fmt"
	"os"
)

func main() { //On met les arguments au format souhaité et on lance RecusiveResolve
	if EntranceTest() == false {
		fmt.Println("Error")
		os.Exit(2)
	}
	var arguments [][9]string
	for index, value := range os.Args {
		if index != 0 {
			var table [9]string
			for index2, value2 := range value {
				if value2 == '.' {
					value2 = '0'
				}
				table[index2] = string(value2)
			}
			arguments = append(arguments, table)
		}
	}
	for index, colones := range arguments {
		for index2 := range colones {
			if ForbidenTest(arguments, [2]int{index, index2}) == true {
				fmt.Println("Error")
				os.Exit(2)
			}
		}
	}
	var path [][2]int
	RecusiveResolve(arguments, [2]int{0, 0}, path)
}
func Forbiden(aurray [][9]string, cell [2]int) []string { // On prent en entrer le tableau et une position et on en ressort toutes les valeurs interdit
	var forbidenNumbers []string
	for _, value := range aurray[cell[0]] {
		if value != "0" {
			forbidenNumbers = append(forbidenNumbers, value)
		}
	}
	for i := 0; i < 8; i++ {
		if aurray[i][cell[1]] != "0" {
			forbidenNumbers = append(forbidenNumbers, aurray[i][cell[1]])
		}
	}
	for i := 0; i < 3; i++ {
		y := (cell[1]/3)*3 + i
		for j := 0; j < 3; j++ {
			x := (cell[0]/3)*3 + j
			if aurray[x][y] != "0" {
				forbidenNumbers = append(forbidenNumbers, aurray[x][y])
			}
		}
	}
	for i := 0; i < len(forbidenNumbers)-1; i++ {
		if forbidenNumbers[i] > forbidenNumbers[i+1] {
			transition := forbidenNumbers[i]
			forbidenNumbers[i] = forbidenNumbers[i+1]
			forbidenNumbers[i+1] = transition
			i = -1
		}
	}
	for index := 0; index < len(forbidenNumbers)-1; index++ {
		if forbidenNumbers[index] == forbidenNumbers[index+1] {
			copy(forbidenNumbers[index:], forbidenNumbers[index+1:]) // Shift a[i+1:] left one index.
			forbidenNumbers[len(forbidenNumbers)-1] = ""             // Erase last element (write zero value).
			forbidenNumbers = forbidenNumbers[:len(forbidenNumbers)-1]
			index = -1
		}
	}
	return forbidenNumbers
}
func NextEmpty(aurray [][9]string, point [2]int) (next [2]int) { //On prent en entrer le tableau et une position et on en ressort la prochaine position vide. C'est aussi une porte de sortie quand il n'y a plus de position suivante
	for i := point[0]; i < len(aurray); i++ {
		ind := 0
		if i == 0 {
			ind = point[1]
		}
		for j := ind; j < len(aurray[i]); j++ {
			if aurray[i][j] == "0" {
				next[0] = i
				next[1] = j
				return next
			}
		}
	}
	Affichage(aurray)
	return next
}
func RecusiveResolve(array [][9]string, position [2]int, path [][2]int) [][9]string { //C'est la fonction qui relite toutes les fonctions en utlisant un tableau une position et un chemin
	nomberOfPath := Occurence(path, position)
	path = append(path, position)
	if position == [2]int{8, 8} && array[8][8] != "0" {
		Affichage(array)
		return array
	} else {
		array[position[0]][position[1]] = "0"
		cont := -1
		forbidden := Forbiden(array, position)
		for i := 1; i < 10; i++ {
			if IfNumberIn(forbidden, i) == false {
				cont++
				if cont == nomberOfPath {
					array[position[0]][position[1]] = string(i + 48)
					return RecusiveResolve(array, NextEmpty(array, position), path)
				}
			}
		}
		array[position[0]][position[1]] = "0"
		return RecusiveResolve(array, LastPoint(path, position), Sup(path, position))
	}
	return array
}
func Occurence(a [][2]int, b [2]int) (len int) { //On prend deux arguments et on regarde le nombre de fois où le second argument est dans le premier
	for _, value := range a {
		if value == b {
			len++
		}
	}
	return
}
func LastPoint(a [][2]int, b [2]int) [2]int { //C'est l'inverse de NextEmpty
	for i := len(a) - 1; i != 0; i-- {
		if a[i][0]*10+a[i][1] < b[0]*10+b[1] {
			if a[i] != b {
				return a[i]
			}
		}
	}
	return [2]int{0, 0}
}
func IfNumberIn(interdit []string, nb int) bool { //On regarde si une valeur est dans une liste
	for _, v := range interdit {
		if v == string(nb+48) {
			return true
		}
	}
	return false
}
func Sup(liste [][2]int, valeur [2]int) [][2]int { //On supprime toutes les occurances d'un éléments dans une liste
	for i := 0; i < len(liste); i++ {
		if liste[i] == valeur {
			transposition := liste[i+1:]
			liste = liste[:i]
			i--
			for _, v := range transposition {
				liste = append(liste, v)
			}
		}
	}
	return liste
}
func Affichage(tableau [][9]string) { //On affiche et on termine la fonction par un Exit(1)
	for index, value := range tableau {
		if index == 0 {
			for i := 0; i < 20; i++ {
				fmt.Print("-")
			}
			fmt.Println()
		}
		fmt.Println(value)
		if index == 8 {
			for i := 0; i < 20; i++ {
				fmt.Print("-")
			}
			fmt.Println()
		}
	}
	os.Exit(1)
}
func ForbidenTest(aurray [][9]string, cell [2]int) bool { // On prent en entrer le tableau et une position et on en ressort toutes les valeurs interdit
	var forbidenNumbers []string
	for index, value := range aurray[cell[0]] {
		if value != "0" && index != cell[1] {
			forbidenNumbers = append(forbidenNumbers, value)
		}
	}
	for i := 0; i < 8; i++ {
		if aurray[i][cell[1]] != "0" && i != cell[0] {
			forbidenNumbers = append(forbidenNumbers, aurray[i][cell[1]])
		}
	}
	for i := 0; i < 3; i++ {
		y := (cell[1]/3)*3 + i
		for j := 0; j < 3; j++ {
			x := (cell[0]/3)*3 + j
			if aurray[x][y] != "0" && x != cell[0] && y != cell[1] {
				forbidenNumbers = append(forbidenNumbers, aurray[x][y])
			}
		}
	}
	for i := 0; i < len(forbidenNumbers)-1; i++ {
		if forbidenNumbers[i] > forbidenNumbers[i+1] {
			transition := forbidenNumbers[i]
			forbidenNumbers[i] = forbidenNumbers[i+1]
			forbidenNumbers[i+1] = transition
			i = -1
		}
	}
	for index := 0; index < len(forbidenNumbers)-1; index++ {
		if forbidenNumbers[index] == forbidenNumbers[index+1] {
			copy(forbidenNumbers[index:], forbidenNumbers[index+1:]) // Shift a[i+1:] left one index.
			forbidenNumbers[len(forbidenNumbers)-1] = ""             // Erase last element (write zero value).
			forbidenNumbers = forbidenNumbers[:len(forbidenNumbers)-1]
			index = -1
		}
	}
	return IfStrIn(forbidenNumbers, aurray[cell[0]][cell[1]])
}
func IfStrIn(interdit []string, nb string) bool { //On regarde si une valeur est dans une liste
	for _, v := range interdit {
		if v == nb {
			return true
		}
	}
	return false
}
func EntranceTest() bool {
	reply := true
	if len(os.Args) == 10 { // validity of the number of input arguments
		reply = true
	} else {
		return false
	}
	for v := 1; v < len(os.Args)-1; v++ { // loops that test if the arguments are not the same twice or +
		testingParameter := os.Args[v]
		if testingParameter == " " {
			return false
		}
		for b := v + 1; b < len(os.Args); b++ {
			if testingParameter == os.Args[b] {
				return false
			}
		}
	}
	for i := 1; i < len(os.Args); i++ { // test to see if argument is valid, numbers and dots
		listEntrance := []rune(os.Args[i])
		if len(listEntrance) == 9 {
			reply = true
		} else {
			return false
		}
		for iNbr := 0; iNbr < 9; iNbr++ {
			if string(listEntrance[iNbr]) <= "9" && string(listEntrance[iNbr]) >= "1" || string(listEntrance[iNbr]) == "." {
				reply = true
			} else {
				return false
			}
		}
		for nbr := 0; nbr < len(listEntrance); nbr++ { // loop test of non-repetition of numbers in the argument
			nbrTest := listEntrance[nbr]
			if string(nbrTest) != "." {
				for nbr2 := nbr + 1; nbr2 < len(listEntrance); nbr2++ {
					if nbrTest == listEntrance[nbr2] {
						return false
					}
				}
			}
		}
	}
	return reply
}
