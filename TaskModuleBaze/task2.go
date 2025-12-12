package main

import "fmt"

func main() {
	A := 50
	r1 := &A
	B := 20
	r2 := &B

	fmt.Println("A =", r1)
	fmt.Println("B = ", r2)

	*r1 = 15
	B = *r1
	fmt.Println("B = ", *r1)
}
