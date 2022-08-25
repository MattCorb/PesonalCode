package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
	sum := addValues(5, 8)
	fmt.Println(sum)

	multisum, multicount := addAllValues(4, 7, 9, 45)
	fmt.Println(multisum, multicount)

}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0

	for _, v := range values {
		total += v
	}
	return total, len(values)
}
