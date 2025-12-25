package main

import (
	"fmt"
)

var magazine = map[int]int{
	1: 10,
	2: 5,
	3: 0,
}

func Chek(itemID int, quantity int) bool {
	return magazine[itemID] >= quantity
}

func placeOrder(itemID int, quantity int) error {
	if !Chek(itemID, quantity) {
		return fmt.Errorf("товара с ID: %d нет на складе", itemID)
	}
	magazine[itemID] -= quantity
	fmt.Printf("Товар ID:%d, quantity: %d\n", itemID, quantity)
	return nil
}

func main() {
	if err := placeOrder(1, 1); err != nil {
		fmt.Println(err)
	}
}
