package main

import (
	"fmt"
	"time"
)

func runMillion() {
	start := time.Now()
	defer func() {
		fmt.Printf("Цикл:%v\n", time.Since(start))
	}()
	var sum int
	for i := 0; i <= 100000000; i++ {
		sum = sum + i
	}
}

func main() {
	runMillion()
}
