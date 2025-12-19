package main

import "fmt"

func removeDublicate(slice []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, value := range slice {
		if !seen[value] {
			result = append(result, value)
			seen[value] = true
		}
	}
	return result
}

func main() {
	original := []string{"a", "b", "a", "c"}
	unique := removeDublicate(original)
	fmt.Println(unique)
}
