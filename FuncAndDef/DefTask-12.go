package main

import "fmt"

func main() {
	x := 10
	y := 20

	defer func(val int) {
		fmt.Println("x:", val)
	}(x) // передача по значению и замарозка выведет 10

	defer func() {
		fmt.Println("y:", y)
	}() // выведет 200

	x = 100
	y = 200

	fmt.Println("Конец main")
}
