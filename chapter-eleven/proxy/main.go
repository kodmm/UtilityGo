package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func FromEnvironment() *Config {
	return &Config{
		HTTPProxy:  getEnvAny("HTTP_PROXY", "http_proxy"),
		HTTPSProxy: getEnvAny("HTTPS_PROXY", "https_proxy"),
		NoProxy:    getEnvAny("NO_PROXY", "no_proxy"),
		CGI:        os.GetEnv("REQUEST_METHOD") != "",
	}
}

func main() {
	// 証明書を読み込む
	cert, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	certPoll := x509.NewCertPool()
	certPoll.AppendCertsFromPEM(cert)
	cfg := &tls.Config{
		RootCAs: certPoll,
	}
	cfg.BuildNameToCertificate()

	// クライアントを作成
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			}, // 検証無効化にする方法
			// TLSClientConfig: cfg,
			Proxy: http.ProxyFromEnvironment, //要設定
		},
	}

	resp, _ := hc.Get("https://www.oreilly.co.jp/index.html")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
