package simpletest_test

import (
	"fmt"
	"reflect"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	if got != 3 {
		t.Errorf("expect 3, but %d", got)
	}
}

func Calc(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("0の除算は未定義です")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("予期しない演算子 %v が設定されました。", operator)
}

type args struct {
	a        int
	b        int
	operator string
}

func TestCalc(t *testing.T) {

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// 正常系
		{
			name: "足し算",
			args: args{
				a:        10,
				b:        2,
				operator: "+",
			},
			want:    12,
			wantErr: false,
		},
		// 異常系(anomalous condition)
		{
			name: "不正な演算子を指定",
			args: args{
				a:        10,
				b:        2,
				operator: ")",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() err = %v, hasErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	// 1. テスト全体の実行前
	setup()

	// 6.テスト全体の実行後
	defer teardown()

	m.Run()
}

func setup() {}

func teardown() {}

func TestHoge(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// テストケース
	}

	// 2.テスト関数の実行前
	defer func() {
		// 5. テスト関数の実行後
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 3. テストケースの実行前
			defer func() {}(
			// 4. テストケースの実行後
			)

			got, err := Hoge(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("failed")
			}
			if got != tt.want {
				t.Errorf("Hoge() = %v, want %v", got, tt.want)
			}
		})
	}

}

type User struct {
	UserID    string
	UserName  string
	Languages []string
}

func TestTom(t *testing.T) {
	tom := User{
		UserID:    "0001",
		UserName:  "Tom",
		Languages: []string{"java", "go"},
	}

	tom2 := User{
		UserID:    "0001",
		UserName:  "Tom",
		Languages: []string{"java", "go"},
	}

	if !reflect.DeepEqual(tom, tom2) {
		t.Errorf("User tom is mismatch, tom=%v tom2=%v", tom, tom2)
	}
}

func TestByTestify(t *testing.T) {
	result, err := Calc(1, 2, "+")
	assert.Nil(t, err)
	assert.Equal(t, 3, result)
}