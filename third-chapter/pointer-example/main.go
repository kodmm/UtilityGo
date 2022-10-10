package main

import "math/big"

// ポインターとしてのみ利用する構造体
type NoCopyStruct struct {
	self  *NoCopyStruct
	Value *string
}

// 初期化時にポインターを保持しておく
func NewNoCopyStruct(value string) *NoCopyStruct {
	r := &NoCopyStruct{
		Value: &value,
	}
	r.self = r
	return r
}

// メソッドの中でチェック
func (n *NoCopyStruct) String() string {
	if n != n.self {
		panic("should not copy NoCopyStruct intstance without Copy() method")
	}
	return *n.Value
}

// 明示的なコピー用メソッド
func (n *NoCopyStruct) Copy() *NoCopyStruct {
	str := *n.Value
	p2 := &NoCopyStruct{
		Value: &str,
	}
	p2.self = p2
	return p2
}

type MutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (m MutableMoney) Currency() Currency {
	return m.currency
}

func (m *MutableMoney) SetCurrency(c Currency) {
	m.Currency = c
}

type ImmutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (im ImmutableMoney) Currency() Currency {
	return im.currency
}

func (im ImmutableMoney) SetCurrency(c Currency) ImmutableMoney {
	return ImmutableMoney{
		currency: c,
		amount:   im.amount,
	}
}

type Node struct {
	name   string
	depth  int
	parent *Node
}

// 初期化しないと各フィールドはゼロ値になる
// ms.name は文字列 depthはゼロ parentはnil
var n Node

type Status int

const (
	DefaultStatus Status = iota
	ActiveStatus
	CloseStatus
)

type Visitor struct {
	Status Status // ゼロ値であるDefaultStatusが設定済みとなる
}

func main() {
}
