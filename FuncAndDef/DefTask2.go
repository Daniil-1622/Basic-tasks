package main

import "fmt"

func demoLifo() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}
func main() {
	demoLifo()
}
