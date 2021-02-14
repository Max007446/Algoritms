package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	fmt.Print("Введите номер месяца: ")
	fmt.Fscan(os.Stdin, &number)

	if number == 1 || number == 2 || number == 12 {
		fmt.Println("зима")
	} else if number == 3 || number == 4 || number == 5 {
		fmt.Println("весна")
	} else if number == 6 || number == 7 || number == 8 {
		fmt.Println("лето")
	} else if number == 9 || number == 10 || number == 11 {
		fmt.Println("осень")

	} else {
		fmt.Println("Введите номер от 1 по 12")
	}
}
