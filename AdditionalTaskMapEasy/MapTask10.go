package main

import "fmt"

func invertMap(origignal map[int]string) map[string]int {
	invert := make(map[string]int)
	for k, v := range origignal {
		invert[v] = k
	}
	return invert
}

func main() {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	result := invertMap(m)
	fmt.Println("Инвертированная мапа: ", result)
}
