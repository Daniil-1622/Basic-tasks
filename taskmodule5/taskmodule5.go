package main

import "fmt"

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func getMax(numbers ...int) int {
	result := numbers[0]
	for _, n := range numbers {
		if n > result {
			result = n
		}
	}
	return result
}

func main() {
	a := []string{"Яблоко", "Помидор", "Дота"}
	x := "Яблоко"
	fmt.Println(contains(a, x))
	fmt.Println(getMax(1, 4, 3, 2))
}
