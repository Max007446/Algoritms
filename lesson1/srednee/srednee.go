package main

import (
	"fmt"
)

func main() {
	var average, ost, n, sum, kolvo int
	n = 1
	for n != 0 {
		fmt.Printf("Введите число (для окончания ввода ввидите 0) \n")
		fmt.Scanln(&n)
		ost = n % 10
		if ost == 8 && n > 0 {
			sum = sum + n
			kolvo++
		}
	}
	average = sum / kolvo
	fmt.Println("Среднее арифмитическое введенных чисел = ", average)

}
