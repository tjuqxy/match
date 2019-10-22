package math

import (
	"fmt"
	"testing"
)

func TestGetNBitsDecimals(t *testing.T) {
	for i:=0; i<3; i++ {
		fmt.Println(GetNBitsDecimals(3))
	}
}

func TestGetTwoDecimals(t *testing.T) {
	for i:=0; i<3; i++ {
		fmt.Println(GetTwoDecimals(3, 2))
	}
}