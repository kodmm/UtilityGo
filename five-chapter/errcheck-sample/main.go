package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// エラーハンドリングを忘れている
	_, err := a()
	if err != nil {
		log.Fatal(err)
	}
	b()
	// このerrはa()から返ってきたエラーでb()の戻り値のエラーではない。
	if err != nil {
		log.Fatal(err)
	}

	err = errors.New("some error")
	//! %wは fmt.Errorf独自の機能
	fmt.Printf("%w", err)
}

func a() (string, error) {
	return "", nil
}

func b() error {
	return errors.New("b error")
}
