package math

import (
	"testing"
)

// 普通的加减乘算式
func TestPrintNormalQuestion(t *testing.T) {
	for i:=0; i<3; i++ {
		PrintNormalQuestion(4, 2, "×")
	}
}

// 带余数的除法算式
func TestPrintNormalDivide(t *testing.T) {
	for i:=0; i<120; i++ {
		PrintNormalDivide(4, 2)
	}
}

// 整除算式
func TestPrintExactDivideQuestion(t *testing.T) {
	for i:=0; i<3; i++ {
		PrintExactDivideQuestion(4, 2)
	}
}

// 限制除数范围的除法算式
func TestPrintLimitDivisorQuestion(t *testing.T) {
	for i:=0; i<3; i++ {
		PrintLimitDivisorQuestion(4, 10, 30)
	}
}

// 小数加法算式
func TestPrintDecimalAddQuestion(t *testing.T) {
	for i:=0; i<3; i++ {
		PrintDecimalAddQuestion(3, 3)
	}
}

// 小数减法算式
func TestPrintDecimalSubtractQuestion(t *testing.T) {
	for i:=0; i<3; i++ {
		PrintDecimalSubtractQuestion(3, 2)
	}
}