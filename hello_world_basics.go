package main

import "fmt"

func main() {

	// var var_name var_type (explicit declaration)
	var name string

	name = "Kelu"

	var num int = -270

	var flag bool = true

	fmt.Println("Hello World!", name, "and", num, "with flag", flag)

	// we can let go guess the type (implicit declaration) - not a good idea always
	var number = 260 // We can't change its value type after giving specific type value
	var st = "Kelu"

	fmt.Printf("%T, %v\n", number, number)
	fmt.Printf("%T\n", st)

	// Expression assignment operator - Walrus operator
	age := 24 // It is equivalent to var age int = 24

	// You can't change its type later on age = "Kelu"

	fmt.Printf("%T, %v\n", age, age)

	// Declaration without assignment - assign default value
	var address string
	var zip_cd uint16
	var flg bool

	println(address, zip_cd, flg)

}
