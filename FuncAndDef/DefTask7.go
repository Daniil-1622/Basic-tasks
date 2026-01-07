package main

import "fmt"

func Defer() {
	x := 10
	fmt.Println(" x:", x)
	defer func() {
		fmt.Println("defer x:", x)
	}()
	x = 20
	fmt.Println("no defer x:", x)
}

func main() {
	Defer()
}
