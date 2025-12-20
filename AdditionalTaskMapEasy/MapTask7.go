package main

import (
	"fmt"
)

func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range map1 {
		result[k] = v
	}
	for k, v := range map2 {
		result[k] = v
	}
	return result
}

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	result := mergeMaps(m1, m2)

	fmt.Println("Результат: ", result)
}
