package main

import "fmt"

type Point struct {
	X int
	Y int
}

var p Point

func main() {

	fmt.Println(p)
	p.X = 34
	p.Y = 90
	fmt.Println(p)

	py := Point{12, 13}
	fmt.Println(py)

	p2 := Point{Y: 20}
	fmt.Println(p2)

	p3 := Point{
		X: 20,
		Y: 25,
	}
	some := p3.add()
	fmt.Println(some)

	p3.printCoordinates()

}

/*
	func add(p Point) int {f
	return (*p).X + (*p).Y
}
*/

func (p Point) add() int {
	return (*p).X + (*p).Y
}

func (p Point) printCoordinates() {
	fmt.Println(p.X, p.Y)
}
