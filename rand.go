package math

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 返回一个取值范围在[min, max)的整数
func GetNumberBetween(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min)+min
}

// 返回一个n位的正整数
func GetNBitsNumber(n int) int {
	if n == 0 || n > 8 {
		return 0
	}

	// 先随机最高位（不能是0）
	number := GetNumberBetween(1, 10)
	for ; n>1; n-- {
		number = number * 10 + GetNumberBetween(0, 10)
	}
	return number
}

// 返回一个n位的小数部分(0, 1)
func GetNBitsDecimalsPart(n int) float64 {
	// 先随机最低位（不能是0）
	number := GetNumberBetween(1, 10)
	divisor := 10
	for ; n>1; n-- {
		number = number + GetNumberBetween(0, 10) * 10
		divisor *= 10
	}
	return float64(number)/float64(divisor)
}

// 返回一个n位的小数（包含整数部分）
func GetNBitsDecimals(n int) float64 {
	// 整数位数
	integerBitNum := GetNumberBetween(0, n)
	// 整数部分
	integerPart := GetNBitsNumber(integerBitNum)
	// 小数部分
	decimalPart := GetNBitsDecimalsPart(n-integerBitNum)
	return float64(integerPart)+decimalPart
}

// 返回两个指定位数的小数，要求第一个大于第二个
func GetTwoDecimals(num1Bits, num2Bits int) (float64, float64) {
	for {
		num1 := GetNBitsDecimals(num1Bits)
		num2 := GetNBitsDecimals(num2Bits)
		if num1 > num2 {
			return num1, num2
		}
	}
}