package main

import "fmt"

func helloWorld() {
	fmt.Println("Hello, world!")
}

// func with arguments
func add(a int, b int) {
	fmt.Println(a + b)
}

// func with same arguments, different way
func add1(a, b int) {
	fmt.Println(a + b)
}

// func with return type
func mathOps(a, b int) int {
	return a + b
}

// func with multiple return types
func helloAdd(a, b int) (string, int) {
	return "Hello", a + b
}

// Label the return type with var_names, which makes it easy to return them
// by just assigning value to them inside the function and lastly write return keyword
func calc(a, b int) (x, y int) {
	x = a + b
	y = a - b
	return
}

// defer keyword is used to specify commands/statements which won't be executed
// until it finds return keyword. It will execute at last when it finds return
// or it finds end of function if there is no return keyword
func testDefer(a, b int) (x, y int) {

	defer fmt.Println("Hello...I'm at with defer statement")

	x = a + b
	y = a - b

	fmt.Println("Before return")
	return
}

func testDefer1(a, b int) {

	defer fmt.Println("Hello...I'm at with defer statement")

	x := a + b
	fmt.Println("Sum = ", x)
	fmt.Println("There's no return")

}

func main() {
	helloWorld()
	add(5, 7)
	add1(5, 7)
	res := mathOps(8, 6)
	fmt.Println(res)

	ans1, ans2 := helloAdd(4, 3)
	fmt.Println(ans1, ans2)

	res1, res2 := calc(12, 8)
	fmt.Println(res1, res2)

	_, _ = testDefer(5, 7)

	testDefer1(8, 2)

}
