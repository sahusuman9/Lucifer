package main

import "fmt"

func main() {
	var added int
	added = add(1, 2)
	fmt.Println(added)

	var swapX, swapY int = 5, 7
	fmt.Printf("SwapX = %d | SwapY = %d\n", swapX, swapY)
	swapX, swapY = swap(swapX, swapY)
	fmt.Printf("SwapX = %d | SwapY = %d\n", swapX, swapY)
}

// int add(int i, int j)
func add(i, j int) int { //datatype can be removed if it's same
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}
