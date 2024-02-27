package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	var colors [2]string
	colors[0] = "black"
	colors[1] = "white"

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
}
