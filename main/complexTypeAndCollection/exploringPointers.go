package main

import "fmt"

func main() {
	anInt := 48
	var p = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 63.78
	pointer1 := &value1
	fmt.Println("Value of pointer1:", *pointer1)

	*pointer1 = *pointer1 / 23
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value of pointer 1:", value1)
}
