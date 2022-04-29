package main

import "fmt"

//gloabal level
var dec int = 11

func main() {
	var dec int = 10
	fmt.Println(dec)

	if anyVar := 9; anyVar > 10 {
		fmt.Println("X is greater than 10")
	} else {
		fmt.Println("X is smaller than 10")
	}
}
