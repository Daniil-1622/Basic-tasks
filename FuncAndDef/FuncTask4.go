package main

import "fmt"

func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, num := range slice {
		if !seen[num] {
			result = append(result, num)
			seen[num] = true
		}
	}
	return result
}

func main() {
	slice := []string{"a", "b", "c", "c", "b"}
	result := RemoveDuplicates(slice)
	fmt.Println("Готовый слайс:", result)
}
