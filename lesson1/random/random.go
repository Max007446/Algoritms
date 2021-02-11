package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min := 0
	max := 100
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(max-min) + min)

	fmt.Printf("Без rand \n")
	fmt.Println(time.Now().Nanosecond() / 10000000)
}
