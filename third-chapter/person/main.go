package main

import (
	"fmt"
	"log"
)

type Person struct {
	FirstName string
	LastName  string
}

// メソッドを追加する型
// 構造体でなくてもプリミティブ型に対する型も可能
type Struct struct {
	v int
}

// レシーバを持つ関数としてメソッドを定義。 レシーバに指定できる型は値型, ポインター型がある。
func (s Struct) PrintStatus() {
	log.Println("Struct:", s.v)
}

// インスタンスのレシーバー
// 値を変更してもインスタンスのフィールドは変更されない
func (s Struct) SetValue(v int) {
	s.v = v
}

// ポインタのレシーバー
// 値を変更できる
func (s *Struct) SetValue(v int) {
	s.v = v
}

func main() {
	// new で作成(あまり使われない)
	p1 := new(Person)
	// var変数宣言で作成
	var p2 Person
	// 複合リテラルで初期化
	p3 := &Person{
		FirstName: "三成",
		LastName:  "石田",
	}
	//! var変数宣言でも、ポインター型の場合はインスタンス作成されない。　X: p4.FirstName
	var p4 *Person
	fmt.Println(p4)

	//! このようなコードは書いては行けない
	s := StructWithPointer{}
	i := 1
	s.v = &i
	s.Modify()
	fmt.Println(*s.v)

}

// NewPerson ファクトリ関数
func NewPerson(first, last string) *Person {
	return &Person{
		FirstName: first,
		LastName:  last,
	}
}

//! イディオムに反する実装例
type StructWithPointer struct {
	v *int
}

//! このメソッドはインスタンスレシーバーだが変更できてしまう。
//! reject案件
func (a StructWithPointer) Modify() {
	(*a.v) = 10
}
