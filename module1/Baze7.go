package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	runes := []rune(str)
	left, right := 0, len(runes)-1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("шалаш"))
}
