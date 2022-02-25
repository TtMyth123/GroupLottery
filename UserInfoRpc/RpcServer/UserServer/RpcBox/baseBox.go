package RpcBox

type ReplyBox struct {
	Result bool
	ErrMsg string
}

type ArgsBox struct {
	UserId int
	GroupId int
}
