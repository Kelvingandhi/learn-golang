package main

import "fmt"

func main() {

	var a []int = []int{2, 3, 6, 4, 3, 7, 9}

	// loop through the slices
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("Another way using range")

	// loop through the slices using range
	for i, element := range a { // similar to enumerate in python
		fmt.Printf("%d: %d\n", i, element)
	}

	// if we don't wanna use i, then we can replace with _ (underscore)
	for _, element := range a {
		fmt.Println(element)
	}

	// Let's solve a problem using range and for loop to print only duplicate in slice
	fmt.Println("Problem Example")
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			element2 := a[j]

			if element == element2 {
				fmt.Println(element)
			}
		}
	}
}
