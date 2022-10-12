package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:9001"
	}
	modifier := func(res *http.Response) error {
		body := make(map[string]interface{})
		dec := json.NewDecoder(res.Body)
		dec.Decode(&body)
		body["fortune"] = "大吉"
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(&body)
		res.Body = ioutil.NopCloser(&buf)
		res.Header.Set("Content-length", strconv.Itoa(buf.Len()))
		return nil
	}

	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}

	http.ListenAndServe(":9000", rp)

}
