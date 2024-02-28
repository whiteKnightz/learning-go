package main

import "fmt"

func main() {
	//fmt.Println(convertToMap([]string{"apple", "banana", "orange", "date"}))
	fmt.Println(convertToMap([]string{"banana", "orange", "apple"}))
	//fmt.Println(convertToMap([]string{"apple", "orange"}))
}

func convertToMap(items []string) map[string]float64 {

	result := make(map[string]float64)

	val := 100 / float64(len(items))

	for _, item := range items {
		result[item] = val
	}

	return result

}
