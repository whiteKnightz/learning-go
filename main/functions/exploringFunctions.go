package main

import "fmt"

func main() {

	sum, length := addAllValues(1, 2, 8, 7, 9, 3)
	fmt.Println(sum)
	fmt.Println(length)
}

func addAllValues(values ...int) (int, int) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, len(values)
}
