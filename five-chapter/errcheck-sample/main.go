package main

import (
	"errors"
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
}

func a() (string, error) {
	return "", nil
}

func b() error {
	return errors.New("b error")
}
