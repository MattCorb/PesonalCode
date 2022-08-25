package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning about memory")

	//to create objects like maps, slices, and channels
	//certain keywords need to be used make() or new()

	//new() - allocates, but doesn't intialize memory
	//returns a location but with 0 memory storage
	//can't add a key value pair

	//make() - allocates and initializes memory
	//allocates non-zeroed memory and returns a storage address
	//add key value pairs right away

	m := make(map[string]int)
	m["theAnswer"] = 42
	m["key"] = 43
	fmt.Println(m)

	//map[key:43 theAnswer:42]

	//memory is deallocated by garbage collector
}
