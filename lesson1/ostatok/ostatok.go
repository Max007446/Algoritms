package main

import (
	"fmt"
)

func main() {
	var N, K, ch, ost int

	fmt.Printf("Введите N: \n")
	fmt.Scanln(&N)
	fmt.Printf("Введите K: \n")
	fmt.Scanln(&K)
	for {
		N = N - K
		ch = ch + 1
		if N < K {
			break
		}
	}
	ost = N
	fmt.Println("Частное = ", ch)
	fmt.Println("Остаток от деления = ", ost)
}
