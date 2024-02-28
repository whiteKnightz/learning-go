package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}

	fmt.Println(colors)

	colors = append(colors, "Yellow")
	colors = append(colors, "Grey")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	//numbers := make([]int, 5)
	numbers[0] = 24
	numbers[1] = 78
	numbers[2] = 7
	numbers[3] = 8
	numbers[4] = 49

	fmt.Println(numbers)

	numbers = append(numbers, 19)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
