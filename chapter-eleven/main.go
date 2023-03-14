package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type customRoundTripper struct {
	transport http.RoundTripper
	logger    func(string, ...interface{})
}

func (t customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// << リクエスト前の実施したい処理を追加 >>
	if t.logger == nil {
		t.logger = log.Printf
	}
	start := time.Now()

	resp, err := t.transport.RoundTrip(req)

	// << リクエスト後に実施したい処理を追加 >>
	if resp != nil {
		t.logger("%s %s %d %s, duration: %d", req.Method, req.URL.String(), resp.StatusCode, http.StatusText(resp.StatusCode),
			time.Since(start))
	}

	return resp, err
}

type basicAuthRoundTripper struct {
	username string
	password string
	base     http.RoundTripper
}

func (rt *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(rt.username, rt.password)
	return rt.base.RoundTrip(req)
}

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
		Timeout: 10 * time.Second,
		Transport: &customRoundTripper{
			transport: http.DefaultTransport,
		},
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
