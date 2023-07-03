package main

import (
	"fmt"
)

// With generics, you can declare and use functions or types that are written to work with
// any of a set of types provided by calling code.

func sumInts(a []int) int {

	s := 0

	for _, v := range a {
		s += v
	}

	return s
}

func sumFloats(a []float32) float32 {

	var s float32 = 0.0

	for _, v := range a {
		s += v
	}

	return s
}

func main() {

	fmt.Println("This is a generics program")

	input_int := []int{1, 2, 3}

	fmt.Println(sumInts(input_int))

	input_flt := []float32{2.3, 2.4, 4.1}
	fmt.Println(sumFloats(input_flt))

}
