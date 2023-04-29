package main

import "fmt"

func main() {
	var ary [3]int // declare array

	ary[0] = 12
	fmt.Println(ary)

	nums := [3]int{10, 20, 50} // Another way to declare & initialize array

	fmt.Println(nums)
	fmt.Println(nums[2])
	fmt.Println(len(nums))

	// loop through array
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// 2D array
	ary2d := [3][2]int{{10, 20}, {50, 100}, {25, 60}}
	fmt.Println(ary2d)

}
