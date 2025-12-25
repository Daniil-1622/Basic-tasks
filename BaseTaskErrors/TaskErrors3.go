package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Println("Введите 5 символов: ")
	fmt.Scanln(&str)

	if num, err := strconv.Atoi(str); err != nil || len(str) > 5 {
		fmt.Println("Не смог прочитать число", err)
	} else {
		fmt.Println("Число: ", num)
	}
}
