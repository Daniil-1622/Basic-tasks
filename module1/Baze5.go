package main

import "fmt"

func main() {
	var str string
	reverseStr := ""
	fmt.Print("Введите строку: ")
	fmt.Scan(&str)

	for _, char := range str {
		reverseStr = string(char) + reverseStr
	}
	fmt.Println(reverseStr)
}
