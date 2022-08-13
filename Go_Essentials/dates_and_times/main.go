package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("dates and times")

	n := time.Now()
	fmt.Println("I practiced go at", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("go launced at ", t)

	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")

	fmt.Printf("the type of parsedTime is %T\n", parsedTime)
}
