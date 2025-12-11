package main

import (
	"fmt"
)

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	numbers := []int{2, 4, 1, 3, 5}
	max := findMax(numbers)
	fmt.Println("Максимальное значения массива: ", max)
}
