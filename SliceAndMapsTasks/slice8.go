package main

import "fmt"

func reverse(slice []int) []int {
	result := make([]int, 0, len(slice))

	for i := len(slice) - 1; i >= 0; i-- {
		result = append(result, slice[i])
	}
	return result
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	result := reverse(data)
	fmt.Println("Результат: ", result)
}
