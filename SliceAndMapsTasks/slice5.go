package main

import "fmt"

func addString(slice *[]string, s string) {
	*slice = append(*slice, s)
}

func main() {
	slice := []string{"a", "b"}
	fmt.Println("До модификации: ", slice)
	addString(&slice, "modified")
	fmt.Println("После модификации: ", slice)
}
