package main

import "fmt"

func findTrulyUnique(arr []int) []int {
	count := make(map[int]int)
	for _, x := range arr {
		count[x]++
	}
	result := []int{}
	for _, x := range arr {
		if count[x] == 1 {
			result = append(result, x)
		}
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 3, 4, 4, 5}
	result := findTrulyUnique(arr)
	fmt.Println("Массив из уникальных элементов", result)
}
