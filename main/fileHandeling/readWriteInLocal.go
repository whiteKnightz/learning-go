package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("")
	content := "Hello there!!"
	file, err := os.Create("./out/fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Printf("Wrote a file of length %v\n", length)
	defer file.Close()

	defer readFile("./out/fromString.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("Text from file:", string(data))
}
