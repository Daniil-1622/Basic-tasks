package main

import "fmt"

func GroupByLength(slice []string) map[int][]string {
	mapper := make(map[int][]string)
	for _, value := range slice {
		len := len(value)
		mapper[len] = append(mapper[len], value)
	}
	return mapper
}

func main() {
	slice := []string{"apple", "banana", "cherry", "orange"}
	mapper := GroupByLength(slice)
	fmt.Println("Группировка: ", mapper)
}
