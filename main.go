package main

import (
	"fmt"
	"otameshi-golang/foo"
)

func main() {
	fmt.Println("Hello, World!")

	PrintEqualLine()

	x := 5
	str := fmt.Sprintf("The number is... %d", x)
	fmt.Println(str)

	PrintEqualLine()

	var n string
	if n == "" { // n is empty
		fmt.Println("n is empty")
	}

	PrintEqualLine()

	foo.Hello()
}
