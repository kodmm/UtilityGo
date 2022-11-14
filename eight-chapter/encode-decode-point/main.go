package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	UserID    string   `json:"user_id"`
	UserName  string   `json:"user_name"`
	Languages []string `json:"languages"`
}

func main() {
	u := user{
		UserID:   "001",
		UserName: "gopher",
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
	// {"user_id":"001","user_name":"gopher","languages":null}

	u = user{
		UserID:   "",
		UserName: "",
		// TODO: スライスの値が存在しないメンバーを、JSONでは空配列[]としてエンコードしたい場合、
		// TODO: 明示的にからのスライスを渡す必要がある。
		Languages: []string{},
	}
	b, _ = json.Marshal(u)
	fmt.Println(string(b))
	// {"user_id":"","user_name":"","languages":[]}
}
