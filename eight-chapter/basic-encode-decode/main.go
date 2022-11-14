package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

type privateIp struct {
	Origin string `json:"origin"`
	// urlと非公開の変数にするとマッピングできない。
	url string `json:"url"`
}

type user struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	// TODO: json.Marshal()ではチャネル型や関数型、複素数型の値が含まれる場合はエラー発生する。
	// TODO: `json:"-"`とすることで、エラーを発生させずにエンコード可能。
	X func() `json:"-"`
}

func main() {
	f, err := os.Open("ip.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var resp ip
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
	// {Origin:255.255.255.255 URL:https://httpbin.org/get}

	s := `{
		"origin": "255.255.255.255",
		"url": "https://httpbin.org/get"
	}`

	// pattern 2: json.Unmarshal
	if err := json.Unmarshal([]byte(s), &resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)

	// TODO: io.Readerを満たしているストリーミングなJSONデータを扱う場合はDecoder()を、それ以外はUnmarshal()を使用する。

	var b bytes.Buffer
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	_ = json.NewEncoder(&b).Encode(u)
	fmt.Printf("%v\n", b.String())

	c, _ := json.Marshal(u)
	fmt.Printf("%v\n", string(c))

	var resprivate privateIp
	if err := json.Unmarshal([]byte(s), &resprivate); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resprivate)
	// {Origin:255.255.255.255 url:}

}
