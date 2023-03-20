package simpletest

import (
	"fmt"
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
		t.Run(tt.name, func(t *testing.T) {
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
