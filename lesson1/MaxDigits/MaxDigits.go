package main

import (
	"fmt"
)

func main() {
	MaxNumber := 0
	a := 10
	b := 1
	c := 54
	d := 150
	if a > b && a > c && a > d {
		MaxNumber = a
	} else if b > a && b > c && b > d {
		MaxNumber = b
	} else if c > a && c > b && c > d {
		MaxNumber = c
	} else if d > a && d > b && d > c {
		MaxNumber = d

	}

	fmt.Println("Максимальное число", MaxNumber)
}