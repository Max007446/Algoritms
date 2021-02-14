package main

import (
	"fmt"
)

func main() {
	m := 5
	x := 10
	c := 0
	_ = c
	c = m
	m = x
	x = c
	fmt.Println(m, x)

	a := 10
	b := 25
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
