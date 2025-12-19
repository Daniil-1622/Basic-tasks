package main

import "fmt"

func insertSlice(s []int, value int, index int) []int {
	newSlice := make([]int, len(s)+1)

	copy(newSlice, s)
	newSlice[index] = index
	copy(newSlice[index+1:], s[index:])
	return newSlice
}

func main() {
	s := []int{1, 2, 4, 5}
	s = insertSlice(s, 2, 3)
	fmt.Println(s)
}
