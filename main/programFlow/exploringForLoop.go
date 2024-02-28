package main

import "fmt"

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Printf("\nThe color on index %v is %v", i, colors[i])
	}

	fmt.Println("\n\n")
	for i := range colors {
		fmt.Printf("\nThe color on index %v is %v", i, colors[i])

	}

	fmt.Println("\n\n")
	for i, color := range colors {
		fmt.Printf("\nThe color on index %v is %v", i, color)

	}

	fmt.Println("\n\n")
	val := 1
	for val < 10 {
		fmt.Println(val)
		val++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("This is the end")
}
