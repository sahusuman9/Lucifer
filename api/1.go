package main

import "fmt"

func main() {
	adder(add)

}

func adder(add func(int, int) int) {
	val := add(3, 5)
	fmt.Println(val)
}

func add(i int, j int) int {
	return i + j
}
