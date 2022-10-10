package main

import (
	"fmt"
	"reflect"
)

func main() {
	emptyStructType := reflect.TypeOf(&struct{}{}).Elem()
	fmt.Println(emptyStructType.Size())

	wait := make(chan struct{}, 1)
	go func() {
		// 空の構造体のインスタンスを送信
		fmt.Println("送信")
		wait <- struct{}{}
	}()
	// データサイズゼロのインスタンスを受け取る
	fmt.Println("受信待ち")
	<-wait
	fmt.Println("受信完了")
}
