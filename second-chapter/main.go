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

}
