package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Printf("До добавления: len: %d, cap: %d\n", len(slice), cap(slice))

	for i := 1; len(slice) < cap(slice); i++ {
		slice = append(slice, i)
		fmt.Printf("После добавления: len: %d, cap: %d\n", len(slice), cap(slice))
	}
}
