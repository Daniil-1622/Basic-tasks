package main

import (
	"errors"
	"fmt"
)

func InvertMap(s map[string]int) (map[int]string, error) {
	result := make(map[int]string)
	for key, value := range s {
		if _, ok := result[value]; ok {
			return nil, errors.New("Ошибка")
		}
		result[value] = key
	}
	return result, nil
}

func main() {
	good := map[string]int{"apple": 5, "banana": 3}
	if inv, err := InvertMap(good); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inv)
	}
	bad := map[string]int{"apple": 5, "banana": 5}
	if inv, err := InvertMap(bad); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(inv)
	}
}
