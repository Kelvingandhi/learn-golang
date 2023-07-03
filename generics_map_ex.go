package main

import "fmt"

// Here, comparable type can be specified as any type that can be used as operands in conditions
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, val := range m {
		s += val
	}

	return s
}

func main() {

	// Initialize a map for the integer values
	ints := map[string]int64{"first": 23, "second": 45}

	// Initialize a map for the float values
	floats := map[string]float64{"first": 21.34, "second": 35.62}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
}
