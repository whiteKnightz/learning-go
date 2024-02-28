package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
  {"name":"orange","price":1.50,"quantity":8},
  {"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	fmt.Println(result)
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	err := json.Unmarshal([]byte(jsonString), &cart)
	if err != nil {
		panic(err)
	}
	return cart
}

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}
