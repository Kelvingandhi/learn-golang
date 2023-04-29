package main

import "fmt"

/* slice represents a slice from an array
Slice itself is a Data Type in golang
Array and Slice has two attributes: length and capacity
Slice can be extended up until the capacity of array (resliced)
Ofc, we can slice off of the slice
*/

func main() {

	var ary [6]int = [6]int{12, 15, 8, 18, 25, 28} // initialize array

	var s []int = ary[1:4] // declare slice off of array

	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println("Length of Slice:", len(s))   // represent number of elements currently in slice
	fmt.Println("Capacity of Slice:", cap(s)) // capacity of slice is computed from starting index of array to end of array

	fmt.Println("Extend Slice: ", s[:cap(s)])

	// Make a slice
	var a []int = []int{2, 5, 8, 10, 15} // It internally creates an array and slice off of it
	fmt.Println(cap(a))
	fmt.Println(cap(a[:3]))

	// To add new elements in slice (extend it)
	b := append(a, 25) // It creates a new slice with 25 added in that
	fmt.Println(b)     // Usually devs assign to the same var

	// 3rd way to create a slice using make() method
	c := make([]int, 5)
	fmt.Printf("%T, %v\n", c, c)

	v := []int{1, 2, 3}
	n := v[:2] // It just references the slice v since it is under capacity of v
	fmt.Println("Slice n=", n, " and v=", v)
	n = append(n, 10) // This will change value of actual slice v
	fmt.Println("Slice n=", n, " and v=", v)

	w := []int{1, 2, 3}
	p := w[:2:2] // It just references the slice v and set capacity to 2
	fmt.Println("Slice p=", p, " and w=", w)
	p = append(p, 10) // It will now creates a new slice since it is beyond capacity of v & won't change value of v
	fmt.Println("Slice p=", p, " and w=", w)

}
