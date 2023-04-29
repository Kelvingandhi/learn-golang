package main

import "fmt"

// Here, just acts like a type e.g. int
type Point struct {
	x int32
	y int32
}

func changeX(p *Point) {
	p.x = 100 // No need to use * to access val, instead just access directly for struct data type
}

func changeY(p Point) {
	p.y = 100
}

func main() {
	var p1 Point = Point{2, 5}

	fmt.Println(p1.x, p1.y)

	// Change value
	p1.x = 8
	fmt.Println(p1.x)

	p2 := Point{-3, 7} // can define just like another data type
	fmt.Println(p2)

	// Also define with one arg, takes default value when arg is not provided
	p3 := Point{y: -8} // here, it will assign default int value 0 to x
	fmt.Println(p3)

	a := 3
	b := &a
	c := &p3 // we can assign pointer of struct type to another var
	fmt.Println(a, b)
	fmt.Println(c)
	changeX(c) // passign pointer of struct type
	fmt.Println(c)

	p4 := Point{8, 9}
	fmt.Println(p4)
	changeY(p4) // It will pass copy of p4
	fmt.Println(p4)

}
