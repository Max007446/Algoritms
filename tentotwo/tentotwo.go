package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Введите число в десятичной системе \n")
	fmt.Scanln(&n)
	perevod(n)

}
func perevod(a int) {
	if a >= 1 {
		perevod(a / 2)
		fmt.Println(a % 2)
	}
}
