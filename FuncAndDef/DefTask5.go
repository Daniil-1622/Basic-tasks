package main

import "fmt"

func example() {
	for i := 1; i < 5; i++ {
		defer fmt.Println("defer", i)
	}
	fmt.Println("end")
}

func main() {
	example()
}
