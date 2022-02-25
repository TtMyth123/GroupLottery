package gBox

type SaveMoneyInfo struct {
	Id          int
	UserId      int
	Money       float64
	AuditorId   int
	AuditorName string
	State       int
	InfoState   SaveMoneyInfoState
	VoucherUrl  string
}
type SaveMoneyInfoState int

const (
	SaveMoneyInfo_State_1_Apply          SaveMoneyInfoState = 1
	SaveMoneyInfo_State_2                SaveMoneyInfoState = 2
	SaveMoneyInfo_State_3_AgreeSaveMoney SaveMoneyInfoState = 3 //同意充值
	SaveMoneyInfo_State_4                SaveMoneyInfoState = 4
)
