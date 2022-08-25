package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files")

	content := "hello from go!"

	file, err := os.Create("./fromString.txt")
	checkError(err)

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("wrote a file with %v\n", length)

	defer file.Close()
	defer readFile("./fromString.txt")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)

	fmt.Println("Text read from file: ", string(data))
}
