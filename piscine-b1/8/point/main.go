package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func setPoint(ptr *point) {
	ptr.x = "42"
	ptr.y = "21"
}

func main() {
	points := &point{}

	setPoint(points)
	res := "x = " + points.x + ", y = " + points.y
	printStr(res)
}
