package main

import "fmt"

func main() {
	shop := map[string]int{
		"apple":  5,
		"banana": 10,
		"cherry": 15,
	}
	fmt.Println("Магазин до завоза: ", shop)
	shop["orange"] = 20
	fmt.Println("Магазин после завоза: ", shop)
}
