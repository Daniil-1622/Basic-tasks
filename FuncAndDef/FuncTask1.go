package main

import "fmt"

func CountFrequency(s []string) map[string]int {
	count := make(map[string]int)
	fmt.Println(s)
	for _, str := range s {
		count[str] = count[str] + 1
	}
	return count
}

func main() {
	inputArray := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	result := CountFrequency(inputArray)
	fmt.Println(result)
}
