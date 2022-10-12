package main

import (
	"fmt"
	"reflect"
	"sync"
)

// 巨大な構造体
type BigStruct struct {
	Member string
}

// Poolは初期化関数をNewフィールドに設定して作成する
var pool = &sync.Pool{
	New: func() interface{} {
		return &BigStruct{}
	},
}

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

	// インスタンスはGet()メソッドで取得
	// 在庫があればそれ、なければNew()を呼び出す
	b := pool.Get().(*BigStruct)

	// 使い終わったらPut()でPoolに戻して次回に備える
	pool.Put(b)
}

// BigStructのインスタンスを作成するファクトリー関数
// 内部でプールを利用
func NewBigStruct() *BigStruct {
	b := pool.Get().(*BigStruct)
	return b
}

// 自分自身を返却するメソッド
func (b *BigStruct) Release() {
	// 初期化してから格納
	b.Member = ""
	pool.Put(b)
}
