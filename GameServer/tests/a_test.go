package main

import (
	"fmt"
	"testing"
	"ttmyth123/kit/pwdKit"
)

func TestAA(t *testing.T) {

	s := fmt.Sprintf("aa %02d", 0)
	fmt.Println(s)
	for i := 0; i < 999999; i++ {
		s1 := fmt.Sprintf("%06d", i)
		s = pwdKit.Sha1ToStr(s1)
		if s == "fEqNCco3Yq9h5ZUglD3CZJT4lBs=" {
			break
		}
	}

	s = pwdKit.Sha1ToStr("111111")
	fmt.Println(s)

	jackpots := "71"

	a1 := getEnd(jackpots, 3)
	fmt.Println("a1:", a1)
}

func getEnd(str string, n int) string {
	iLen := len(str)
	if iLen > n {
		i := iLen - n
		return str[i:]
	}
	return str
}
