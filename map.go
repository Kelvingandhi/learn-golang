package main

import "fmt"

// Map doesn't preserve the order

func main() {

	var mp map[string]int = map[string]int{
		"apple":  1,
		"orange": 2,
		"pear":   3,
	}

	fmt.Println(mp)
	fmt.Println(mp["apple"])

	mp["apple"] = 10

	fmt.Println(mp["apple"])

	mp["grape"] = 25

	delete(mp, "apple")

	fmt.Println(mp)

	// Another way to create a map
	mp2 := map[string]int{"One": 1, "Two": 2}

	fmt.Println(mp2)

	// Another way to create a map
	mp3 := make(map[string]int) // initialize empty map

	fmt.Println(mp3)

	// Trick to check if key exists in map or not
	if _, ok := mp3["cherry"]; !ok {
		fmt.Println("Key not found in map")
	}

	val, ok := mp["orange"]
	fmt.Println(val, ok)

	fmt.Println(len(mp))

	// Loop through map
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
