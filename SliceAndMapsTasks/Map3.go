package main

import "fmt"

func Filter(map1 map[string]int, n int) map[string]int {
	result := make(map[string]int)

	for key, value := range map1 {
		if value > n {
			result[key] = value
		}
	}
	return result
}

func main() {
	var n int
	fmt.Print("Введите n: ")
	fmt.Scanln(&n)
	map1 := map[string]int{"a": 10, "b": 5, "c": 20}
	result := Filter(map1, n)
	fmt.Println("Отфильтрованная мапа:", result)
}
