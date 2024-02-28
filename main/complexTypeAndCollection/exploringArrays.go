package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[0])
	fmt.Println(colors[1])
	fmt.Println(colors[2])

	var numbers = [5]int{6, 8, 1, 5, 3}
	fmt.Println(numbers)
	fmt.Printf("\nNumbers of colors: %v", len(colors))
	fmt.Printf("\nNumbers of numbers: %v", len(numbers))
}
