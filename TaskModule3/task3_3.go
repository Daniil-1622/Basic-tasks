package main

import "fmt"

func main() {
	week := []int{1, 2, 3, 4, 5}
	weekends := []int{6, 7}

	result := append(week, weekends...)

	fmt.Println("Неделя: ", result)

}
