package main

import (
	"fmt"
)

func main() {
	MaxNumber := 0
	a := 10
	b := 100
	c := -5
	if a >= b && a >= c {
		MaxNumber = a
	} else if b >= a && b >= c {
		MaxNumber = b
	} else if c >= a && c >= b {
		MaxNumber = c
	}

	fmt.Println("Максимальное число", MaxNumber)

}
