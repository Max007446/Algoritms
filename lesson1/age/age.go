package main

import (
	"fmt"
	"os"
)

func main() {
	var age, ost int
	fmt.Printf("Введите возврат человека от 1 до 150 \n")
	fmt.Fscan(os.Stdin, &age)
	ost = age % 10
	if ost == 1 {
		println("Возраст человека ", age, "год")
	} else if ost > 1 && ost < 5 {
		println("Возраст человека ", age, "года")
	} else if ost >= 5 {
		println("Возраст человека ", age, "лет")
	}
}
