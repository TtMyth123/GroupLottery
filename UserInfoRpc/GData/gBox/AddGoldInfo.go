package gBox

import "github.com/TtMyth123/UserInfoRpc/models/mconst"

type AddGoldInfo struct {
	UserId    int
	GroupId   int
	Gold      float64
	T         mconst.Account_Type
	Des       string
	ExtraData interface{}

	Des2  string
	DesMp string
}
