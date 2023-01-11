package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type record struct {
	Number  int    `csv:"number"`
	Message string `csv:"message"`
}

func main() {
	f, err := os.Open("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c := make(chan record)
	done := make(chan bool)
	go func() {
		if err := gocsv.UnmarshalToChan(f, c); err != nil {
			log.Fatal(err)
		}
		done <- true
	}()

	for {
		select {
		case v := <-c:
			fmt.Printf("%+v\n", v)
		case <-done:
			return
		}
	}
}
