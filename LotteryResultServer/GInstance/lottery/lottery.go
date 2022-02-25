package lottery

import "errors"

type ILotteryServer interface {
	GetNewAwardInfo() NewAwardInfo
	UpdateResultUrl(string) error
}

type BLotteryServer struct {
	ILotteryServer
}

func (this BLotteryServer) GetNewAwardInfo() NewAwardInfo {
	return NewAwardInfo{}
}

func (this BLotteryServer) UpdateResultUrl(url string) error {
	return errors.New("未实现")
}
