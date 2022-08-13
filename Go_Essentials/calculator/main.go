package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the calculator \n Please enter two numbers")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic("Sorry that wasn't a number")
	}

	fmt.Println("enter a second number: ")

	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		panic("Sorry that wasn't a number")
	}

	total := float1 + float2
	total = math.Round(total*100) / 100

	fmt.Printf("You entered %v and %v, the sum is  %v", float1, float2, total)

}
