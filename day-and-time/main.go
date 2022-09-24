package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))

	now2 := time.Date(1995, time.October, 26, 9, 0, 0, 0, time.Local)
	past := time.Date(1995, time.November, 12, 6, 38, 0, 0, time.UTC)

	fmt.Println(now2)
	fmt.Println(past)
}
