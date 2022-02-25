package gBox

type DrawMoneyInfo struct {
	Id             int
	UserId         int
	GroupId        int
	Money          float64
	AuditorId      int
	AuditorName    string
	DrawMoneyState DrawMoneyInfoState
	State          int
	VoucherUrl     string
	OrderId        string
}

type DrawMoneyInfoState int

const (
	DrawMoneyInfo_State_1 DrawMoneyInfoState = 1
	DrawMoneyInfo_State_2 DrawMoneyInfoState = 2
	DrawMoneyInfo_State_3 DrawMoneyInfoState = 3
	DrawMoneyInfo_State_4 DrawMoneyInfoState = 4
	DrawMoneyInfo_State_5 DrawMoneyInfoState = 5
	DrawMoneyInfo_State_6 DrawMoneyInfoState = 6
	DrawMoneyInfo_State_7 DrawMoneyInfoState = 7
	DrawMoneyInfo_State_8 DrawMoneyInfoState = 8
	DrawMoneyInfo_State_9 DrawMoneyInfoState = 9
)
