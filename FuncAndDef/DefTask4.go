package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("Отложенное число:", x)
	x = 99
	fmt.Println("Текущее число:", x)
}
