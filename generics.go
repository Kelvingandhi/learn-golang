package main

import (
	"fmt"
)

// With generics, you can declare and use functions or types that are written to work with
// any of a set of types provided by calling code.

// Declaring interface of multiple types we want to pass input params to generic function.
type Number interface {
	int64 | float64
}

// Generic Function - Declaring it can take any param of type Number as input and
// we have used same type as return type as well.
func sumIntsOrFloats[T Number](a []T) T {

	var s T

	for _, v := range a {
		s += v
	}

	return s
}

func main() {

	fmt.Println("This is a generics program")

	input_int := []int64{1, 2, 3}
	fmt.Println(sumIntsOrFloats(input_int))

	input_flt := []float64{2.3, 2.4, 4.1}
	fmt.Println(sumIntsOrFloats(input_flt))

}
