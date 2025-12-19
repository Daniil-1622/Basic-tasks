package main

import "fmt"

func inPlace(slice []int) {
	for i := range slice {
		slice[i] = slice[i] * 2
	}
	fmt.Println(slice)
}

func main() {
	slice := []int{1, 2, 3}
	inPlace(slice)
}
