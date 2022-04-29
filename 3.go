package main

import "fmt"

//var, const, type

var nameOfTheVariable int

const nameOfTheVariable = 12

func main() {
	nameOfTheVariable := 50
	fmt.Printf("%T\n", nameOfTheVariable)

	//zero value
	//int = 0
	//float64 = 0.0
	//string = ""
}
