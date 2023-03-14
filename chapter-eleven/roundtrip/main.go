package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"
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

type retryableRoundTripper struct {
	base     http.RoundTripper
	attempts int
	waitTime time.Duration
}

func (rt *retryableRoundTripper) shouldRetry(resp *http.Response, err error) bool {
	// ネットワークによるリトライ
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) {
			return true
		}
	}

	// レスポンスコードによるリトライ
	if resp != nil {
		if resp.StatusCode == 429 || (500 <= resp.StatusCode && resp.StatusCode <= 504) {
			return true
		}
	}

	// リトライすべきではないため、リトライしない
	return false
}

func (rt *retryableRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var (
		resp *http.Response
		err  error
	)
	for count := 0; count < rt.attempts; count++ {
		resp, err = rt.base.RoundTrip(req)

		if !rt.shouldRetry(resp, err) {
			return resp, err
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(rt.waitTime): // リトライのため待機
		}
	}
	return resp, err
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
