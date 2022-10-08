package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
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
}
