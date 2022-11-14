package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
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

type Rectangle struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Record struct {
	ProcessID string `json:"process_id"`
	DeletedAt JSTime `json:"deleted_at"`
}

type JSTime time.Time

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

	s := `{
		"width": 5,
		"height": 10,
		"radius": 6
	}`

	var rect Rectangle
	d := json.NewDecoder(bytes.NewReader([]byte(s)))

	// 想定していないJSONフィールドが存在している場合はバリデーションとしてエラーをするための記述
	d.DisallowUnknownFields()

	if err := d.Decode(&rect); err != nil {
		// Error Handling
	}

	r := &Record{
		ProcessID: "001",
		DeletedAt: JSTime(time.Now()),
	}

	e, _ := json.Marshal(r)
	fmt.Println(string(e))

	var t *Record
	if err := json.Unmarshal(e, &t); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", time.Time(t.DeletedAt).Format(time.RFC3339Nano))

}

func Int(v int) *int {
	return &v
}

func (t JSTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if tt.IsZero() {
		return []byte("null"), nil
	}
	v := strconv.Itoa(int(tt.UnixMilli()))
	return []byte(v), nil
}

func (t *JSTime) UnmarshalJSON(data []byte) error {
	var jsonNumber json.Number
	err := json.Unmarshal(data, &jsonNumber)
	if err != nil {
		return err
	}
	unix, err := jsonNumber.Int64()
	if err != nil {
		return err
	}

	*t = JSTime(time.Unix(0, unix))
	return nil
}
