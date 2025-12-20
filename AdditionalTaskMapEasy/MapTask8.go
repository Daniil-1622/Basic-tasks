package main

import "fmt"

func main() {
	text := "Hello"
	test := make(map[rune]int)

	for _, char := range text {
		test[char]++
	}
	for char, count := range test {
		fmt.Printf("%s: %d\n", string(char), count)
	}
}
