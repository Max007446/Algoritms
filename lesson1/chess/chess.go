package main

import (
	"fmt"
	"os"
)

func main() {
	var x1, y1, x2, y2 int
	var point1, point2 string
	fmt.Printf("Введите координаты \n")
	fmt.Fscan(os.Stdin, &x1, &y1, &x2, &y2)
	if (x1%2 == 1 && y1%2 == 1) || (x1%2 == 0 && y1%2 == 0) {
		point1 = "black"

	} else {
		point1 = "white"
	}
	if (x2%2 == 1 && y2%2 == 1) || (x2%2 == 0 && y2%2 == 0) {
		point2 = "black"

	} else {
		point2 = "white"
	}
	if point1 == point2 {
		fmt.Println("Поля относятся к одному цвету ")
	} else {
		fmt.Println("Поля не относятся к одному цвету ")
	}
}
