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

type FormInput struct {
	Name        string `json:"name"`
	CompanyName string `json:"company_name,omitempty"`
}

type Bottle struct {
	Name  string `json:"name"`
	Price int    `json:"price,omitempty"`
	KCal  *int   `json:"kcal,omitempty"` // *intのポインター型で宣言する
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

	in := FormInput{
		Name: "Nissy",
	}
	b, _ = json.Marshal(in)
	fmt.Println(string(b))

	bottle := Bottle{
		Name:  "ミネラルウォーター",
		Price: 0,
		KCal:  Int(0),
	}
	out, _ := json.Marshal(bottle)
	fmt.Println(string(out))
}

func Int(v int) *int {
	return &v
}
