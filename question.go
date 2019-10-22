package math

import "fmt"

// 整数普通的加减乘算式
func PrintNormalQuestion(num1Bits, num2Bits int, symbol string) {
	num1 := GetNBitsNumber(num1Bits)
	num2 := GetNBitsNumber(num2Bits)
	result := num1 + num2
	switch symbol {
	case "-":
		result = num1 - num2
	case "×", "*":
		result = num1 * num2
	}
	fmt.Printf("%d%s%d=%d\n", num1, symbol, num2, result)
}

// 整数除法带余数
func PrintNormalDivide(num1Bits, num2Bits int) {
	num1 := GetNBitsNumber(num1Bits)
	num2 := GetNBitsNumber(num2Bits)
	result := num1 / num2
	remainder := num1 - result * num2
	fmt.Printf("%d÷%d=%d￿……%d\n", num1, num2, result,remainder)
}

// 小数的加法算式
func PrintDecimalAddQuestion(num1Bits, num2Bits int) {
	num1 := GetNBitsDecimals(num1Bits)
	num2 := GetNBitsDecimals(num2Bits)
	fmt.Printf("%v + %v = %v\n", num1, num2, num1+num2)
}

// 小数的减法算式
func PrintDecimalSubtractQuestion(num1Bits, num2Bits int) {
	num1, num2 := GetTwoDecimals(num1Bits, num2Bits)
	fmt.Printf("%v - %v\n", num1, num2)
}

// 整除算式
func PrintExactDivideQuestion(num1Bits, num2Bits int) {
	num2 := GetNBitsNumber(num2Bits)
	minDividend := 1
	// 构造最小的N位数
	for i:=1; i<num1Bits; i++ {
		minDividend *= 10
	}
	// 最小的乘数
	minMultiply := minDividend / num2 + 1
	// 最大的乘数
	maxMultiply := minDividend * 10 / num2
	num1 := GetNumberBetween(minMultiply, maxMultiply) * num2
	fmt.Printf("%d ÷ %d\n", num1, num2)
}

// 限制除数范围的算式
func PrintLimitDivisorQuestion(num1Bits, minNum2, maxNum2 int) {
	num1 := GetNBitsNumber(num1Bits)
	num2 := GetNumberBetween(minNum2, maxNum2)
	fmt.Printf("%d ÷ %d\n", num1, num2)
}