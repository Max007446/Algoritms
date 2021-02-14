package main

import (
	"fmt"
	"math"
)

func main() {
	var sh, n2 int
	var n, sq float64
	i := 0.0
	fmt.Printf("Введите предел поиска чисел: \n")
	fmt.Scanln(&n)
	n2 = int(n)
	for {
		n2 = n2 / 10
		if n2 == 0 {
			sh = sh + 1
			break
		}
		sh = sh + 1

	}
	fmt.Println("Количество разрядов в числе:", sh)
	for j := 0; j < sh; j++ {
		for ; i < n; i++ {

			sq = math.Pow(i, 2)
			if int(sq)%int(math.Pow10(j)) == int(i) {
				fmt.Println(i)
			}
			if i > math.Pow10(j) {
				break
			}

		}

	}
}
