package main

import "fmt"

func example() (result int) {
	defer func() {
		result = result * 2
	}()
	if true {
		result = 21
		return
	}
	result = 0
	return
}

func main() {
	fmt.Println(example())
}
