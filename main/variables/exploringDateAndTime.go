package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("\n\nThe time now is %v", now)

	t := time.Date(2019, time.November, 14, 21, 30, 0, 0, time.UTC)
	fmt.Printf("\n\nThe custom time is %v", t)
	fmt.Println("\n\n")
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 06 15:04 MST"))

	parsedTime, _ := time.Parse(time.ANSIC, "Thu Nov 14 21:30:00 2019")
	fmt.Printf("\n\nThe type of parsedTime is %T\n\n", parsedTime)
}
