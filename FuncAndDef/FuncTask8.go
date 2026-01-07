package main

import "fmt"

func Average(slice []int) float64 {
	var total float64
	if slice == nil {
		return 0.0
	}
	for _, nums := range slice {
		total += float64(nums) / float64(len(slice))
	}
	return total
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	total := Average(slice)
	fmt.Println(total)
}
