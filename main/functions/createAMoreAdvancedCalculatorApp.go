package main

import (
	"fmt"
	"strconv"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Println(result)
}

func calculate(input1 string, input2 string, operation string) float64 {
	val1 := convertInputToValue(input1)
	val2 := convertInputToValue(input2)
	switch operation {
	case "+":
		return addValues(val1, val2)
	case "-":
		return subtractValues(val1, val2)
	case "*":
		return multiplyValues(val1, val2)
	case "/":
		return divideValues(val1, val2)
	default:
		panic("Invalid operation")
	}
	return 0
}

func convertInputToValue(input string) float64 {
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		msg := fmt.Sprintf("%v cannot be parsed", input)
		panic(msg)
	}
	return val
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
