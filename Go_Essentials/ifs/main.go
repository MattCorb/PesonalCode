package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional Logic")

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "less than zero"
	} else if theAnswer == 0 {
		result = "equal to zero"
	} else {
		result = "greater than zero"
	}

	fmt.Println(result)

	if x := -42; x < 0 {
		result = "less than zero"
	} else if x == 0 {
		result = "equal to zero"
	} else {
		result = "greater than zero"
	}

	fmt.Println(result)

}
