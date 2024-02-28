package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10, "Woof!"}

	poodle.Speak()

	poodle.Sound = "Arrf!"
	poodle.Speak()
	poodle.SpeakThreeTime()
	poodle.SpeakThreeTime()
	poodle.Speak()
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how does the Dog speak
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTime is how does the Dog speaks trice
func (d Dog) SpeakThreeTime() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
