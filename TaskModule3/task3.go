package main

import "fmt"

func main() {
	week := []int{1, 2, 3, 4, 5, 6, 7}

	weekends := make([]int, 2)
	copy(weekends, week)
	week = week[0:5]

	fmt.Println("Будни: ", week)
	fmt.Println("Выходные", weekends)
}
