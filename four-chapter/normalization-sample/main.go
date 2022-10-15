package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

// コアとなる基本実装
func Normalize(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		s, err := br.ReadString('\n')
		if s != "" {
			io.WriteString(w, norm.NFKC.String(s))
		}
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

// ファイルとの読み書きがしやすいようにファイル名を受け取るようにしたラッパー
func NormalizeFile(input, output string) error {
	r, err := os.Open(input)
	if err != nil {
		return err
	}
	defer r.Close()
	w, err := os.Create(output)
	if err != nil {
		return err
	}
	defer w.Close()
	return Normalize(w, r)
}

// ユニットテストや他のコードから呼び出しやすいように文字列を渡すだけで利用できるラッパー
func NormalizeString(i string) (string, error) {
	r := strings.NewReader(i)
	var w strings.Builder
	err := Normalize(&w, r)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

// ジェネリクスが使われるcase
// TODO 返り値がanyで、受け取り側で毎回キャストが必要
// ポインターと値のうち、ポインターだけを受け取れるようにしたい
func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// interface{}型 == any型
var v any = "何でも入る変数"
// 整数も入る
v = 1
// 浮動小数点数も入る
v = 3.141592

// interface{}型のスライス
var slices = []any{
	"関ヶ原",
	1600,
}
// interface{}型のmap jsonの格納などで役立つ
var ieyasu = map[string]any{
	"名前": "徳川家康"
	"生まれ": 1543
}

func main() {
	// "favorite"というキーに対して、"銭形平次"という値を設定する。
	ctx := context.WithValue(context.Background(), "favorite", "銭形平次")

	// ctx.Value()はinterface{}なので変換が必要
	// 必ず、okで成功可否を確認すること
	if s, ok := ctx.Value("favorite").(string); ok {
		// sはstring型
		log.Printf("私の好きなものは%sです\n", s)
	}
	
	// あるいは型スイッチ
	// 同じvだが、case節ごとにvの型が変わる
	switch v := ctx.Value("favorite").(type) {
	case string:
		log.Printf("好きなものは: %s\n", v)
	case int:
		log.Printf("好きな数値は: %d\n", v)
	case complex128:
		log.Printf("好きな複素数: %f\n", v)
	default: // どれにもマッチしない場合
		log.Printf("好きなものは: %v\n", v)
	}
}
