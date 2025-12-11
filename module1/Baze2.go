package main

import (
	"fmt"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	var x int
	fmt.Println("Введите x:", x)
	fmt.Scanln(&x)
	var y int
	fmt.Println("Введите y:", y)
	fmt.Scanln(&y)
	result := Sum(x, y)
	fmt.Println("Результат: ", result)

}
