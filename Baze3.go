package main

import (
	"fmt"
)

func Even(x int) {
	if x%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

func main() {
	var x int
	fmt.Println("Введите x:", x)
	fmt.Scanln(&x)
	Even(x)
}
