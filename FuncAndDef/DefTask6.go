package main

import "fmt"

func Defer(a, b int) (sum int) {
	defer func() {
		sum = a + b
	}()
	a = 100
	return sum
}

func main() {
	total := Defer(10, 6)
	fmt.Println(total)
}
