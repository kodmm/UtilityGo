package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func findBook(isbn string) (*Book, error) {
	// ...
	// 値が取得できなかったため、ErrNotFoundを返す
	return nil, ErrNotFound
}

func validate(length int) error {
	if length <= 0 {
		return fmt.Errorf("length must be greater than0, length = %d", length)
	}
	// ...
	return nil
}

func main() {

}
