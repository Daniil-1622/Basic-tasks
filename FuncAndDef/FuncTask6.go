package main

import "fmt"

func ChunkSlice(slice []int, size int) [][]int {
	if size <= 0 {
		return nil
	}
	result := make([][]int, 0)
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	size := 3
	result := ChunkSlice(slice, size)
	fmt.Println(result)
}
