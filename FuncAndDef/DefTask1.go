package main

import "fmt"

func demoDeffer() {
	fmt.Println("Start")
	defer fmt.Println("Defer")
	fmt.Println("End")
}
func main() {
	demoDeffer()
}
