package main

import "fmt"

func Unific(map1 map[string]int, map2 map[string]int) map[string]int {

	result := make(map[string]int)

	for k, v := range map1 {
		result[k] = v
	}
	for k, v := range map2 {
		result[k] += v
	}
	return result
}

func main() {
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4}

	result := Unific(map1, map2)
	fmt.Println("Результат: ", result)
}
