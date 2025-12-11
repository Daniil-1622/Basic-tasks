package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "104"
	num := 35

	s, _ := strconv.Atoi(str)
	fmt.Println("Приобразованная строка", s)

	strconv.Itoa(num)
	fmt.Println("Преобразованное число в строку", num)
}
