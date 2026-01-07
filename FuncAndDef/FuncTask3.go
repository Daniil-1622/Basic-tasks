package main

import "fmt"

func MergeMaps(s1, s2 map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range s1 {
		result[k] = v
	}
	for k, v := range s2 {
		result[k] = v
	}
	return result
}

func main() {
	s1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	s2 := map[string]int{
		"c": 3,
		"d": 4,
	}
	result := MergeMaps(s1, s2)
	fmt.Println("Обьединение двух мап: ", result)
}
