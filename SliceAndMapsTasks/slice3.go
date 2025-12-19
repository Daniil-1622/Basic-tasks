package main

import "fmt"

func deleteIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3}
	result := deleteIndex(slice, 2)
	fmt.Println("Срез без удаленного элемента", result)
}
