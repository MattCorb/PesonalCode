package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()

	poodle.SpeakThreeTime()
	poodle.Speak()
}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTime is when the dog sees a squirell
// when the dog is passed as a receiver a copy is passed
// thats why it goes back to normal
func (d Dog) SpeakThreeTime() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
