package main

import (
	"errors"
	"fmt"
	"math"
)

var NegativeNumberError = errors.New("negative number error")

func NegativeNumber(number int) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("Отрицательное число:%w", NegativeNumberError)
	}
	return math.Sqrt(float64(number)), nil
}

func main() {
	result, err := NegativeNumber(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
