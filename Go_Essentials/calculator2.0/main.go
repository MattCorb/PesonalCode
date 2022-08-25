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
	fmt.Println("Calculator 2.0")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		panic("Sorry that wasn't number")
	}

	fmt.Println("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		panic("sorry that wan't a  number")
	}

	fmt.Println("Operation: (+, -, *, /)")

	input3, _ := reader.ReadString('\n')
	operation := strings.TrimSpace(input3)

	oCalc := Calculator{Value1: float1, Value2: float2, Operation: operation}

	var r string

	switch operation {
	case "+":
		r = strconv.FormatFloat(oCalc.Addition(), 'g', 5, 64)

	case "-":
		r = strconv.FormatFloat(oCalc.Subtraction(), 'g', 5, 64)

	case "*":
		r = strconv.FormatFloat(oCalc.Multiplication(), 'g', 5, 64)

	case "/":
		r = strconv.FormatFloat(oCalc.Division(), 'g', 5, 64)

	default:
		r = "sorry this wasn't a proper operation"

	}

	fmt.Println("Result: ", r)

}

// Calculator struct
type Calculator struct {
	Value1    float64
	Value2    float64
	Operation string
}

// Addition is a method to add two values
func (c Calculator) Addition() float64 {
	total := c.Value1 + c.Value2
	total = math.Round(total*100) / 100
	return total

}

// Subtraction is a method to subtract two values
func (c Calculator) Subtraction() float64 {
	total := c.Value1 - c.Value2
	total = math.Round(total*100) / 100
	return total

}

func (c Calculator) Multiplication() float64 {
	total := c.Value1 * c.Value2
	total = math.Round(total*100) / 100
	return total
}

func (c Calculator) Division() float64 {
	total := c.Value1 / c.Value2
	total = math.Round(total*100) / 100
	return total
}
