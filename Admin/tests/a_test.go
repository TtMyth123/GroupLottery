package main

import (
	"fmt"
	"strings"
	"testing"
	"ttmyth123/kit/pwdKit"
)

func TestAA(t *testing.T) {
	s := fmt.Sprintf("aa %02d", 0)
	fmt.Println(s)

	s = pwdKit.Sha1ToStr("515888")
	fmt.Println(s)
}

func TestBetStr(t *testing.T) {
	BetStr := `;头特 大:$10;头特 小:$10`
	BetStr = strings.Replace(BetStr, " ", "_", -1)
	BetStr = strings.Replace(BetStr, "$", "", -1)
	arr := strings.Split(BetStr, ";")
	ArrBetDetail := make([]string, len(arr)-1)
	for i := 1; i < len(arr); i++ {
		ArrBetDetail[i-1] = arr[i] + "元"
	}
	fmt.Println(ArrBetDetail)
}
