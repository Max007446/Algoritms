package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, cube, sq float64
	fmt.Printf("Введите a :\n")
	fmt.Fscan(os.Stdin, &a)
	fmt.Printf("Введите b: \n")
	fmt.Fscan(os.Stdin, &b)

	for ; a < b; a++ {
		cube = math.Pow(a, 3)
		sq = math.Pow(a, 2)
		fmt.Println(cube, sq)
	}
}
