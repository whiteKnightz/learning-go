package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]string)
	key1 := "ON"
	key2 := "BC"
	key3 := "QC"
	key4 := "AB"

	m[key1] = "Ontario"
	m[key2] = "British Colombia"
	m[key3] = "Quebec"

	fmt.Println(m)
	fmt.Println(m[key1])
	fmt.Println(m[key2])
	fmt.Println(m[key3])

	delete(m, key3)
	fmt.Println(m)

	m[key4] = "Alberta"
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("\n%v -> %v", k, v)
	}

	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	fmt.Println("\n", keys)

	sort.Strings(keys)
	fmt.Println("\n", keys)

	for i := range keys {
		fmt.Printf("\n%v -> %v", keys[i], m[keys[i]])
	}
}
