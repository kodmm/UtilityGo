package main

import (
	"fmt"
	"time"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range items {
		v2 := v
		go func() {
			fmt.Printf("v = %d\n", v2)
		}()
	}
	time.Sleep(time.Second)

	for _, v := range items {
		go func(v int) {
			fmt.Printf("v = %d\n", v)
		}(v)
	}
}
