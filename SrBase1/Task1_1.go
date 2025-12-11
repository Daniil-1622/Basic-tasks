package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	euroKms := 120.4
	euroKmc := euroKms * 3.6
	euroSpeed := EuropeanVelocity(math.Round(euroKmc*100) / 100)
	fmt.Println("Европейская скорость: ", euroSpeed)

	usaMc := 130.0
	usaMvc := usaMc * 2.23694
	usaSpeed := AmericanVelocity(math.Round(usaMvc*100) / 100)
	fmt.Println("Американская скорость: ", usaSpeed)
}
