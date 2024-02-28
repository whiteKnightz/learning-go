package main

import "fmt"

func main() {
	poodle := Dog{"Poodle", 10}

	printDog(poodle)

	poodle.Weight = 15
	fmt.Println("\n\n")

	printDog(poodle)

}

func printDog(poodle Dog) {
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\n", poodle.Breed)
	fmt.Printf("Weight: %v\n", poodle.Weight)
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
