package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {

	var x []int = []int{2, 3, 4}

	y := x
	y[0] = 100 // change in y reflects to value of x as well, since both are pointers to that slice
	fmt.Println(x, y)

	var z map[string]int = map[string]int{"hello": 1}
	w := z
	w["world"] = 2 // change in w reflects to value of z as well, since both are pointers to that map
	fmt.Println(z, w)

	z["welcome"] = 3 // change in z reflects to value of y as well
	fmt.Println(z, w)

	// Another example using function
	var p []int = []int{1, 2, 3}
	fmt.Println("Before function call:", p)
	changeFirst(p)
	fmt.Println("After function call:", p)

	// Array is also mutable DT, BUT it doesn't have same behavior.
	var a [2]int = [2]int{3, 4}
	b := a
	b[0] = 100 // change in b WON'T reflects change to value of a, because b is separate copy from a
	fmt.Println(a, b)

	// This isn't true with immutable types like int, string, etc.
	// It always creates a copy when assing to another variable.

}
