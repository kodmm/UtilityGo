pacakge main

// ポインターとしてのみ利用する構造体
type NoCopyStruct struct {
	self *NoCopyStruct
	Value *string
}

// 初期化時にポインターを保持しておく
func NewNoCopyStruct(value string) *NoCopyStruct {
	r := &NoCopyStruct{
		Value: &value
	}
	r.self = r
	return r
}
// メソッドの中でチェック
func (n *NoCopyStruct) String() {
	if n != n.self {
		panic("should not copy NoCopyStruct intstance without Copy() method")
	}
	return *n.Value
}

// 明示的なコピー用メソッド
func (n *NoCopyStruct) Copy() *NoCopyStruct {
	str := *n.Value
	p2 := &NoCopyStruct{
		Value: &str
	}
	p2.self = p2
	return p2
}
func main() {

}