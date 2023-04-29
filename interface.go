package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rec struct {
	length float64
	width  float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rec) area() float64 {
	return r.length * r.width
}

func main() {
	c1 := circle{12.25}
	r1 := rec{25, 30}

	fmt.Println(c1.area())
	fmt.Println(r1.area())

	shapes := []shape{&c1, r1} // It is good practice to use reference operator, just in case we get pointer of obj in function arguments

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

}
