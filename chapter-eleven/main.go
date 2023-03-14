package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type User struct {
	Name string
	Addr string
}

func main() {
	u := User{
		Name: "O'Reilly Japan",
		Addr: "Tokyo Pachioji",
	}

	payload, err := json.Marshal(u)
	if err != nil {
		// ....
	}

	resp, err := http.Post("http://example.com/", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		// ...
	}

	defer resp.Body.Close()
}
