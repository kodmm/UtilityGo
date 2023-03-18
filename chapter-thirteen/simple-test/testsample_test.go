package testsample

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
