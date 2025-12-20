package main

import "fmt"

func main() {
	scores := map[string]int{
		"math":      90,
		"physics":   85,
		"chemistry": 78,
	}
	for k := range scores {
		score := scores[k]
		fmt.Println(score)
	}

}
