package main

import (
	"fmt"
)

// constants can be infered outside function
const aConst int = 64

func main() {

	var aString string = "This is a string in go"

	fmt.Println(aString)
	fmt.Printf("The variable's type is a %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	//Infered values

	var anotherString = "This is another String"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Printf("the variable's type is %T\n", anotherInt)

	//can only be infered inside a function
	//myString := "this is also a string"
	fmt.Printf("the variable's type is %T\n", aConst)

}
