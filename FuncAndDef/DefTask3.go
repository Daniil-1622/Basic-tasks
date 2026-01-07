package main

import "fmt"

func addOne() (result int) {
	result = 5
	defer func() {
		result = result + 1
	}()
	return
}

func main() {
	fmt.Println(addOne())
}
