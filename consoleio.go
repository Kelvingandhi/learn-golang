package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type year number: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("Input type %T and value %v\n", input, input)
	fmt.Printf("Current year minus given year = %d\n", 2022-input)
}
