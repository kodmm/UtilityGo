package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title      string
	Author     Author
	Publisher  string
	ReleasedAt time.Time
	ISBN       string
}

type Author struct {
	FirstName string
	LastName  string
}

func main() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	book := Book{
		Title:      "Real World HTTP",
		Author:     Author{FirstName: "渋川", LastName: "しぶや"},
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4873119030",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book.Title)

	// testコードなどで一つの関数にしか使用しない一時的な構造体を宣言する場合
	book1 := struct {
		Title      string
		Author     Author
		Publisher  string
		ReleasedAt time.Time
		ISBN       string
	}{
		Title:      "Real World HTTP",
		Author:     Author{FirstName: "渋川", LastName: "しぶや"},
		Publisher:  "オライリー・ジャパン",
		ISBN:       "4873119030",
		ReleasedAt: time.Date(2017, time.June, 14, 0, 0, 0, 0, jst),
	}
	fmt.Println(book1.Title)

	b := &Book{
		Title: "Mithril",
	}
	fmt.Println(b.Title)    // OK
	fmt.Println((*b).Title) // 上記と同じ意味

	// デリファレンス: ポインターから値を取り出すこと。
	b2 := &b
	// fmt.Println(b2.Title) // NG
	fmt.Println((**b2).Title) // 1 or 2つデリファレンスすればOK
}
