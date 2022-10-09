package main

import (
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"strPtr"`
	Bool    bool    `map:"bool"`
	BoolPtr *bool   `map:"boolPtr"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"intPtr"`
}

func main() {
	src := MapStruct{
		Str:     "string-value",
		StrPtr:  &[]string{"string-value"}[0],
		Bool:    true,
		BoolPtr: &[]bool{true}[0],
		Int:     12345,
		IntPtr:  &[]int{6789}[0],
	}
	dest := map[string]string{}
	Encode(dest, &src)
	log.Println(dest)
}

func Encode(target map[string]string, src interface{}) error {
	v := reflect.ValueOf(src)
	e := v.Elem()
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 埋め込まれた構造体は再帰処理
		if f.Anonymous {
			if err := Encode(target, e.Field(i).Addr().Interface()); err != nil {
				return err
			}
			continue
		}
		key := f.Tag.Get("map")
		// タグがなければフィールド名をそのまま使う
		if key == "" {
			key = f.Name
		}
		// 子供が構造体だったら再帰処理(名前は引き継ぐ)
		if f.Type.Kind() == reflect.Struct {
			Encode(target, e.Field(i).Addr().Interface())
			continue
		}
		// フィールドの型を取得
		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Pointer {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			isP = true
			// ポインターのポインターは無視
			if k == reflect.Pointer {
				continue
			}
		}
		switch k {
		case reflect.String:
			if isP {
				// nilならデータは読み込まない
				if e.Field(i).Pointer() != 0 {
					target[key] = *(*string)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				target[key] = e.Field(i).String()
			}
		case reflect.Bool:
			var b bool
			if isP {
				if e.Field(i).Pointer() != 0 {
					b = *(*bool)(unsafe.Pointer(e.Field(i).Pointer()))
				}
			} else {
				b = e.Field(i).Bool()
			}
			target[key] = strconv.FormatBool(b)
		case reflect.Int:
			var n int64
			if isP {
				if e.Field(i).Pointer() != 0 {
					n = int64(*(*int)(unsafe.Pointer(e.Field(i).Pointer())))
				}
			} else {
				n = e.Field(i).Int()
			}
			target[key] = strconv.FormatInt(n, 10)
		}
	}

	return nil

}
