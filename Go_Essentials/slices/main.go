package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	//slice is an abstraction layer that sits on an array
	//resizeable
	//like arrays all items are of the same type

	//declaration is the same as an array without a defined lenth
	var colors = []string{"red", "green", "blue"}
	fmt.Println("the colors are: ", colors)

	colors = append(colors, "purple")
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	//last 5 caps the length remove to make a slice
	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 135
	numbers[2] = 1
	numbers[3] = 14
	numbers[4] = 13

	numbers = append(numbers, 345)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
