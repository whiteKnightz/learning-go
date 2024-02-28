package main

import "fmt"

func main() {
	var cart []cartItem
	var apples = cartItem{"apple", 2.99, 3}
	var oranges = cartItem{"orange", 1.50, 8}
	var bananas = cartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)

	fmt.Println(result)
}

func calculateTotal(cart []cartItem) float64 {
	var sum float64
	for _, item := range cart {
		sum += item.price * float64(item.quantity)
	}
	return sum
}

type cartItem struct {
	name     string
	price    float64
	quantity int
}
