package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		arguments := os.Args
		if len(arguments) >= 2 {
			for i := range arguments {
				if i != 0 {
					file, err := os.ReadFile(string(arguments[i]))
					if err != nil {
						text := "ERROR: open " + string(arguments[i]) + ": no such file or directory"
						printStr(text)
						z01.PrintRune('\n')
						os.Exit(1)
					} else {
						printStr(string(file))
					}
				}
			}
		} else if len(arguments) == 1 {
			printStr("File name missing")
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
