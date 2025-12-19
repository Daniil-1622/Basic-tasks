package main

import "fmt"

func main() {
	ages := map[string]int{
		"Алиса": 30,
		"Боб":   25,
	}
	names := Keys(ages)
	scores := Values(ages)
	fmt.Println("Имена:", names)
	fmt.Println("Значения:", scores)
}

func Keys(m map[string]int) []string {
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

func Values(m map[string]int) []int {
	var result []int
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
