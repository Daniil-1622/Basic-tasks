package main

import "fmt"

func sliceAdd(n int) {
	slice := []int{}
	for i := 1; i <= n; i++ {
		slice = append(slice, i)
	}
	fmt.Println("Исходный: ", slice)

	for i := n + 1; i <= 2*n; i++ {
		slice = append(slice, i)
	}
	fmt.Println("После append: ", slice)
}

func main() {
	n := 3
	sliceAdd(n)
}
