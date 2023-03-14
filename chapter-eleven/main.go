package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
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

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: http.DefaultTransport,
	}

	// Getリクエストの生成
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://example.com", nil)
	if err != nil {
		// ...
	}
	// リクエストヘッダーにトークンを付与
	req.Header.Add("Authorization", "Bearer XXX ... XXX")

	// HTTPリクエストの発行
	response, err := client.Do(req)
}
