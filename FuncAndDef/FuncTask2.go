package main

import "fmt"

func FilterEven(s []int) []int {
	result := make([]int, 0)
	for _, v := range s {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	result := FilterEven(slice)
	fmt.Println(result)
}
