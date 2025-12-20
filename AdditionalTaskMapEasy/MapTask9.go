package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 4}
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}
	dublicate := []int{}
	for num, count := range count {
		if count > 1 {
			dublicate = append(dublicate, num)
		}
	}
	fmt.Println("Дубликаты:", dublicate)
}
