package main

import "fmt"

// & <- This is get the reference operator (memory location for value of variable)
// * <- dereference operator used to change value of variable using memory reference.
// Both operators work with all types - mutable, immutable

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {

	x := 7
	y := &x // assigning reference of x to variable y

	fmt.Println(x, y)
	fmt.Printf("%T\n", y) // Here, y is NOT type int, but *int

	*y = 8 // dereference operator used to change value of variable
	fmt.Println(x, y)

	s := "Hello"
	fmt.Println(s)
	changeValue(&s) // passing reference of s
	fmt.Println(s)  // It will change value of s

	s1 := "World"
	fmt.Println(s1)
	changeValue2(s1) // passing value of s1
	fmt.Println(s1)  // It won't change value of s1

	// Basically this is what it does
	s2 := "Hello World"
	var pointer *string = &s2
	fmt.Println(pointer, *pointer)   // value of pointer (ref of s2), dereferencing s2
	fmt.Println(&pointer, *&pointer) // reference of pointer, dereferencing pointer

}
