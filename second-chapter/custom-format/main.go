package main

import (
	"encoding/json"
	"fmt"
)

type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c ConfidentialCustomer) String() string {
	return "xxxx-xxxx-xxxx"
}

func (c ConfidentialCustomer) GoString() string {
	return "xxxx-xxxx-xxxx"
}

func main() {
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "4111-1234-5678",
	}

	fmt.Println(c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)

	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes))
}
