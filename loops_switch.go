package main

import "fmt"

func main() {
	x := 2

	// first way
	for x < 5 { // Similar to while loop
		fmt.Println("x:", x)
		x += 1 // also write x++
	}

	// second way
	for y := 0; y < 5; y += 2 {
		fmt.Println("y:", y)
	}

	k := 5

	switch k {
	case 1, -1: // check multiple case evaluate to same execution inside it
		fmt.Println("Multiple case checking => k:", k)

	case 3:
		fmt.Println("case 3 => k:", k)

	case 4:
		fmt.Println("case 4 => k:", k)

	case 5:
		fmt.Println("case 5 => k:", k)

	default:
		fmt.Println("Default case: => k:", k)
	}

}
