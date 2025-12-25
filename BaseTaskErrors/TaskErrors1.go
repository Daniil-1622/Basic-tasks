package main

import (
	"errors"
	"fmt"
)

func Devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("На ноль делить нельзя")
	}
	return a / b, nil
}

func main() {
	a := 10
	b := 0
	result, err := Devide(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
