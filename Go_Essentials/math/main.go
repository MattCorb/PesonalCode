package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("here we will learn about math")

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3

	fmt.Println("Integer sum:", intSum, "second half")

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3

	fmt.Println("Float sum:", floatSum)

	//use = because the variable is being reassigned not intialiezed
	floatSum = math.Round(floatSum*100) / 100

	fmt.Println("the rounded sum is", floatSum)

	circleRadius := 15.5
	circumfrence := circleRadius * 2 * math.Pi
	fmt.Printf("circumfrence: %.2f\n", circumfrence)

}
