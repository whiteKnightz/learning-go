package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	anInt := 7
	aFloat := 42.7

	sum := float64(anInt) + aFloat

	fmt.Printf("\n\nThe orignal sum value is:%v \nThe rounded sum value is:%v\n\n",
		sum, math.Round(sum))

	i1, i2, i3 := 12, 45, 68

	intSum := i1 + i2 + i3

	fmt.Printf("The integer sum is: %v \n\n", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Printf("The float sum is: %v \n\n", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Printf("The float sum is now: %v \n\n", floatSum)

	circleRadius := 15.5
	circumfrence := circleRadius * 2 * math.Pi
	fmt.Printf("The circumfrence: %.2f \n\n", circumfrence)

	value1 := "10"
	value2 := "5.5"
	val1F, err1 := strconv.ParseFloat(value1, 64)
	if err1 != nil {
		fmt.Println(err1)
		panic("Value 1 not a number")
	}
	val2F, err2 := strconv.ParseFloat(value2, 64)
	if err2 != nil {
		fmt.Println(err2)
		panic("Value 2 not a number")
	}
	sumF := math.Round((val1F+val2F)*10) / 10
	fmt.Printf("The value should be 15.500000: %v \n\n", sumF)
}
