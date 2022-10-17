package main

import "errors"

var ErrNotFound = errors.New("not found")

func findBook(isbn string) (*Book, error) {
	// ...
	// 値が取得できなかったため、ErrNotFoundを返す
	return nil, ErrNotFound
}

func main() {

}
