package main

import "fmt"

func helloWorld(a int) {
	fmt.Println("Hello, world!", a)
}

func test2(myFunc func(int) int) { // function as argument
	fmt.Println(myFunc(7))
}

func returnFunc(a string) func() {
	return func() {
		fmt.Println(a) //using var a inside this function is called function closures
	}
}

func multiply(a int) func(int) int {
	return func(v int) int {
		return a * v
	}
}

// we can assign function in variable which reference it and we can call it
// whenever we want
func main() {

	x := helloWorld
	fmt.Printf("Type of var x is %T\n", x)
	x(5)

	// function inside a function
	y := func(a int) int {
		return a * -1
	}
	fmt.Println(y(8)) // calling function

	// calling function directly when defining
	z := func(a int) int {
		return a * -1
	}(4) // this will call function and save the result in z
	fmt.Println(z) // printing var z

	// Passing function to another function call
	test2(y)

	// calling function that will return a function
	r := returnFunc("Hiiii") // saving return val (as function) into var r
	r()                      // calling returned function

	returnFunc("Hola")() // another way to call returned function

	res := multiply(5)(6) // func return func with parameter and make use of closure
	fmt.Println(res)

}
