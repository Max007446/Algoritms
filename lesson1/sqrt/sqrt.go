package main

import (
	"fmt"
	"math"
)

func main() {
	var d, x1, x2 float64
	a := 3.0
	b := 10.0
	c := 5.0
	d = b*b - 4*a*c
	if d > 0 {
		x1 = fn0(x1, b, d, a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
	}
	fmt.Println("Решение уравнения 3x2+10x+5=0 \n", d, x1, x2, math.Sqrt(d))
}

func fn0(x1 float64, b float64, d float64, a float64) float64 {
	x1 = (-b + math.Sqrt(d)) / (2 * a)
	return x1
}
