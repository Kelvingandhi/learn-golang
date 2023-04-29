package main

import "fmt"

// Embedded structs meaning using struct field inside another

type Point struct {
	x int32
	y int32
}

// If we don't use pointers then it will create copies every time

type Circle struct {
	radius float64
	center *Point //Usually use pointer of struct to derefence its value
}

// if var names are different than other struct, then no need to give name
type Circle1 struct {
	radius float64
	*Point // Not giving name here, access Point's var directly
}

type Circle2 struct {
	radius float64
	Point
}

func main() {
	p := &Point{3, 4} // pointer of type Point
	c := Circle{4.56, p}

	p.x = -10
	fmt.Println(c)                      // new value of x will reflect in c
	fmt.Println(c.center.x, c.center.y) // accessing Point variables

	c1 := Circle1{2.34, &Point{6, 8}}
	fmt.Println(c1.x, c1.y)

	// If we don't use reference & pointer
	p2 := Point{-5, -8}
	c2 := Circle2{3.45, p2}
	p2.x = -10
	fmt.Println(c2) // new value of x won't be reflect in c2
}
