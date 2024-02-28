package main

import "fmt"

func main() {
	var result string

	result = setResult(42)
	fmt.Println(result)

	result = setResult(0)
	fmt.Println(result)

	result = setResult(-2)
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "It is less than zero"
	} else if x == 0 {
		result = "It is zero"
	} else {
		result = "It is grater than zero"
	}

	fmt.Println(result)
}

func setResult(theAns int) string {
	if theAns < 0 {
		return "It is less than zero"
	} else if theAns == 0 {
		return "It is zero"
	} else {
		return "It is grater than zero"
	}
}
