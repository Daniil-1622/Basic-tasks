package main

import (
	"fmt"
	"math"
)

func main() {
	var length float64 = 35.0
	P2 := 6.28318
	var R float64 = length / P2
	var r1 *float64 = &R
	fmt.Println("Радиус окружности равен: ", &r1)

	p2 := 3.14159
	S := p2 * (*r1) * (*r1)
	rounded := (math.Round(S*100) / 100)
	fmt.Println("Округленная площадь окружности равна: ", rounded)
}
