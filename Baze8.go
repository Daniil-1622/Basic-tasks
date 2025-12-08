package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 4, 5, 6}

	sum := 0
	for _, v := range arr {
		sum = sum + v
	}
	fmt.Println("Сумма массива: ", sum)
}
