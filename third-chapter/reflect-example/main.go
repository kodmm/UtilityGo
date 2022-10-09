package main

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type MapStruct struct {
	Str     string  `map:"str"`
	StrPtr  *string `map:"str"`
	Bool    bool    `map:"bool"`
	BoolPtr bool    `map:"bool"`
	Int     int     `map:"int"`
	IntPtr  *int    `map:"int"`
}

func main() {
	src := map[string]string{
		"str":  "string data",
		"bool": "true",
		"int":  "12345",
	}
	var ms MapStruct = MapStruct{
		Str: "test",
	}
	Decode(&ms, src)
	log.Printf("%#v\n", ms)
}

func Decode(target interface{}, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	fmt.Printf("v: %#v\n", v)
	fmt.Printf("e: %#v\n", e)
	return decode(e, src)
}

func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 埋め込まれた構造体は再帰処理
		if f.Anonymous {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// 子供が構造体だったら再帰処理
		if f.Type.Kind() == reflect.Struct {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// タグがなければフィールド名をそのまま使う
		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		// 元データになければスキップ
		sv, ok := src[key]
		if !ok {
			continue
		}

		// フィールドの型の取得
		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Pointer {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()
			// ポインターのポインターは無視
			if k == reflect.Pointer {
				continue
			}
			isP = true
		}
		switch k {
		case reflect.String:
			if isP {
				e.Field(i).Set(reflect.ValueOf(&sv))
			} else {
				e.Field(i).SetString(sv)
			}
		case reflect.Bool:
			b, err := strconv.ParseBool(sv)
			if err == nil {
				if isP {
					e.Field(i).Set(reflect.ValueOf(&b))
				} else {
					e.Field(i).SetBool(b)
				}
			}
		case reflect.Int:
			n64, err := strconv.ParseInt(sv, 10, 64)
			if err == nil {
				if isP {
					n := int(n64)
					e.Field(i).Set(reflect.ValueOf(&n))
				} else {
					e.Field(i).SetInt(n64)
				}
			}
		}

	}
	return nil
}
