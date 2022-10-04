package main

import (
	"container/list"
	"fmt"
	"io"
	"net/url"
	"os"
	"reflect"
)

func main() {
	// プリミティブ型
	// type SourceType int

	// 構造体
	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	// var int int = 1

	writeType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writeType))
	fmt.Println("status: ", StatusForbidden)
	fmt.Println("route: ", AizuNumataKaido)

	v1 := url.Values{}
	// v2 := make(url.Values, 2)

	v1.Add("key1", "value1")
	v1.Add("key2", "value2")

	for k, v := range v1 {
		fmt.Printf("%s: %s\n", k, v)
	}

	// container/listを使う
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// container/listはNext()で返り値がnilでない間はループ
	for ele := l.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele.Value)
	}

}

//go:generate enumer --type=HTTPStatus -json
type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

type NationalRoute int

const (
	NagasakiKaido   NationalRoute = 200
	AizuNumataKaido NationalRoute = 401
	HokurikuDo      NationalRoute = 402
	KurinokiBypass  NationalRoute = 403
)

func (n NationalRoute) String() string {
	switch n {
	case NagasakiKaido:
		return "長崎街道"
	case AizuNumataKaido:
		return "会津沼田街道"
	case HokurikuDo:
		return "北陸道"
	case KurinokiBypass:
		return "栗の木バイパス"
	default:
		return fmt.Sprintf("国道%d号線", n)
	}
}
