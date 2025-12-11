package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 1

	for i := 0; i < 10; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println("10 Число Фибоначчи: ", b)
}
