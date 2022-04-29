package main

import "fmt"

func main() {

	v := 10
	p := &v // assign address to pointer p

	fmt.Println(*p) // print value at p

	*p = 12 // update value at by 12

	fmt.Println(v)

	var another int = 30
	var aPointer *int = &another ///address of another

	fmt.Println(another)
	//value at
	*aPointer = 56
	fmt.Println(another)
}
