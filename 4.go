package main

import "fmt"

func main() {
	j := 0
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for j < 5 {
		fmt.Println(j)
	}

	for {
		fmt.Println(i)
	}
}
