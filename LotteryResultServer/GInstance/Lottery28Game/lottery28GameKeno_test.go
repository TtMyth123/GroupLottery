package Lottery28Game

import (
	"fmt"
	"testing"
)

func Test_getNum(t *testing.T) {
	nums := []int{7, 8, 14, 16, 17, 22, 26, 34, 39, 41, 42, 48, 54, 58, 63, 64, 69, 72, 73, 79}
	aaa1 := getNum(1, nums)
	aaa2 := getNum(2, nums)
	aaa3 := getNum(3, nums)
	fmt.Println(aaa1, aaa2, aaa3)
}
