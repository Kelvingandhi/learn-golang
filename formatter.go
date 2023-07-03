package main

import "fmt"

func main() {

	fmt.Printf("Hello %T, %v\n", 10, 10)
	var s string = fmt.Sprintf("Number %v %T", "world", "world")
	fmt.Println("Value of s=", s)

	// integers
	fmt.Printf("Number: %b, %o, %d, %x, %X\n", 1025, 1025, 1025, 3435, 3435)

	// floating point numbers
	fmt.Printf("Float: %e, %f, %F, %g\n", 23.648027465908, 23.648027465908, 23.648027465908, 23.648027465908)

	// boolean
	fmt.Printf("Value: %t, %t\n", true, false)

	var f bool = true
	fmt.Printf("%v\n", f)

	// string
	fmt.Printf("Hello: %s, %q\n", "world", "welcome")

	// width and padding
	fmt.Printf("Hello, %8s is cool\n", "kelu")
	fmt.Printf("Hello, %-8s is cool\n", "kelu")
	fmt.Printf("Number: %f, %.2f, %3.2f, %.f\n", 3.5384930, 3.5384930, 3.5384930, 3.5384930)

	fmt.Printf("Number: %07f\n", 35.47)
	fmt.Printf("Number: %07d,\t,%2d", 3547, 3547)

	fmt.Println()
}
