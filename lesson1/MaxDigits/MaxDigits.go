package Maxdigits

import (
	"fmt"

	"github.com/max007446/Algoritms/lesson1/Indexbody/Indexbody"
)

func main() {
	MaxNumber := 0
	a := 10
	b := 5004
	c := 5004
	d := 5004
	if a >= b && a >= c && a >= d {
		MaxNumber = a
	} else if b >= a && b >= c && b >= d {
		MaxNumber = b
	} else if c >= a && c >= b && c >= d {
		MaxNumber = c
	} else if d >= a && d >= b && d >= c {
		MaxNumber = d

	}

	fmt.Println("Максимальное число", MaxNumber)
	Indexbody.Indexbody()
}
