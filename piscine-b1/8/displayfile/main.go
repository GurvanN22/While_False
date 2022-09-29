package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 2 {
		file, err := os.ReadFile(string(arguments[1]))
		if err == nil {
		}
		fmt.Print(string(file))
	} else if len(arguments) == 1 {
		fmt.Println("File name missing")
	} else if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	}
}
