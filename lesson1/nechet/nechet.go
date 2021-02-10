package main

import (
	"fmt"
)

func main() {
	var N, ost int
	flag := false
	fmt.Printf("Введите N: \n")
	fmt.Scanln(&N)
	for N > 10 {
		N = N / 10
		ost = N % 2

		if ost == 1 {
			fmt.Println("TRUE")
			flag = true
			break
		}

	}
	if flag == false {
		fmt.Println("FALSE")
	}
}
