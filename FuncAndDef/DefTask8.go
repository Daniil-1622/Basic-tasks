package main

import "fmt"

func One() {
	fmt.Println("Начало")
	defer fmt.Println("3")
	fmt.Println("2")
}

func Two() {
	fmt.Println("4")
	One()
	defer fmt.Println("Конец")
}

func main() {
	One()
	Two()
}
