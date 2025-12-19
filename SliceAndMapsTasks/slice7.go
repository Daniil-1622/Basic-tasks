package main

import "fmt"

func Even(slice []int) []int {
	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 != 0 {
			slice[j] = slice[i]
			j++
		}
	}
	return slice[:j]
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	result := Even(data)
	fmt.Println("Результат: ", result)
}
