package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "This is GO!!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T", aString)

	fmt.Println("\n\n")
	var anInt int = 42
	fmt.Println(anInt)
	fmt.Printf("The variable's type is %T", anInt)

	var defaultInt int
	fmt.Println(defaultInt)

	fmt.Println("\n\n")
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T", myString)

	fmt.Println("\n\n")
	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T", aConst)
}
