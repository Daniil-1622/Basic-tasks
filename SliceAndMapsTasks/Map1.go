package main

import "fmt"

func CountChar(s []string) map[string]int {
	counts := make(map[string]int)
	for _, str := range s {
		counts[str]++
	}
	return counts
}

func main() {
	slice := []string{"a", "b", "a", "c", "b"}
	counts := CountChar(slice)
	fmt.Println("Подсчет среза: ", counts)
}
