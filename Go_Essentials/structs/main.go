package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")

	poodle := Dog{"Poodle", 10}
	fmt.Printf("%+v\n", poodle)
}

// kinda like classes but without inheritance
// custom types
// uppercase first letter = "exported" can be used elsewhere
// no upper case = "non-exported" = private
// property names same rule as above
type Dog struct {
	Breed  string
	Weight int
}
