package main

import "fmt"

func Anagrams(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	for _, r := range t {
		counts[r]--
	}
	for _, count := range counts {
		if count != 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "listen"
	s2 := "silent"
	fmt.Println("Результат: ", Anagrams(s1, s2))
}
