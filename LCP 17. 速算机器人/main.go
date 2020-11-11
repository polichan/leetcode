package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d", calculate("AB"))
}

func calculate(s string) int {
	x, y := 1, 0
	str := []rune(s)
	for _, item := range str{
		if string(item) == "A" {
			x  = calA(x, y)
		}else{
			y = calB(x, y)
		}
	}
	return x + y
}

func calA(x, y int) int {
	x = 2 * x + y
	return x
}
func calB(x, y int) int {
	y = 2 * y + x
	return y
}