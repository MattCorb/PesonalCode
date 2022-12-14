package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().Unix())

	// fmt.Println("Day", dow)

	var result string

	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's sunday"
		// fallthrough
	case 2:
		result = "It's monday"
		// fallthrough
	default:
		result = "It's some other day"
	}

	fmt.Println(result)

}
