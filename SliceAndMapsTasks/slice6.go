package main

import "fmt"

func Unification(slice []int, s []int) []int {

	total := len(slice) + len(s)

	result := make([]int, total)

	copy(result, slice)
	copy(result[len(slice):], s)
	return result
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	result := Unification(a, b)
	fmt.Println("Готовый срез: ", result)
}
