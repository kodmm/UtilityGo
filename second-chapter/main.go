package main

import (
	"fmt"
	"io"
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
