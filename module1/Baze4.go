package main

import (
	"fmt"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	var x int
	fmt.Println("Введите x:", x)
	fmt.Scanln(&x)
	result := factorial(x)
	fmt.Println("Факториал числа x:", result)
}
