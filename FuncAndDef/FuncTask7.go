package main

import "fmt"

func FindUnique(slice []int) []int {
	count := make(map[int]int)
	for _, num := range slice {
		count[num]++
	}
	var result []int
	for _, num := range slice {
		if count[num] == 1 {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	slice := []int{1, 1, 2, 3, 4, 6, 6}
	result := FindUnique(slice)
	fmt.Println(result)
}
