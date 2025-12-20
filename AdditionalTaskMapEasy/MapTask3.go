package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	if _, exits := ages["Bob"]; exits {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not exists")
	}
}
