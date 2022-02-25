package LotteryUscGame

import (
	"fmt"
	"strings"
	"time"
	"ttmyth123/GroupLottery/LotteryResultServer/GInstance/lottery"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/timeKit"
)

type UscGameR struct {
	timeI time.Duration
	t     int

	CurNums       string
	aNewAwardInfo lottery.NewAwardInfo
	lottery.BLotteryServer
}

func NewUscGameR(timeI time.Duration, t int) *UscGameR {
	aUscGameR := new(UscGameR)
	aUscGameR.timeI = timeI
	aUscGameR.t = t

	aUscGameR.aNewAwardInfo, aUscGameR.CurNums = aUscGameR.getAwardInfo()
	aUscGameR.run()
	return aUscGameR
}

func (this *UscGameR) GetNewAwardInfo() lottery.NewAwardInfo {
	return this.aNewAwardInfo
}

func (this *UscGameR) getAwardInfo() (lottery.NewAwardInfo, string) {
	Nums := GetNums(this.t)
	t := time.Now()
	newT := t.Add(this.timeI * time.Minute)
	sNewNum := newT.Format("060102150405")
	sLotteryNum := t.Format("060102150405")

	aNewAwardInfo := lottery.NewAwardInfo{
		LotteryNum:      strconvEx.StrTry2Int64(sLotteryNum, 0),
		LotteryStr:      sLotteryNum,
		ResultNums:      Nums,
		NextLotteryStr:  sNewNum,
		NextLotteryTime: newT,
		CurLotteryTime:  t,
	}
	return aNewAwardInfo, Nums
}
func (this *UscGameR) run() {
	ticker1 := time.NewTicker(this.timeI * time.Minute)
	go func() {
		for {
			select {
			case <-ticker1.C:
				this.aNewAwardInfo, this.CurNums = this.getAwardInfo()

				//this.CurNums = GetNums(this.t)
				//t := time.Now()
				//newT := t.Add(this.timeI * time.Minute)
				//sNewNum := newT.Format("060102150405")
				//sLotteryNum := t.Format("060102150405")
				//
				//this.aNewAwardInfo = UscModel.NewAwardInfo{
				//	NewNum:strconvEx.StrTry2Int64(sNewNum,0),
				//	NewNumTime:newT,
				//	NumList:GetArrStr2Nums(this.CurNums),
				//	LotteryNum:strconvEx.StrTry2Int64(sLotteryNum,0),
				//	LotteryTime:t,
				//	ServerTime:time.Now(),
				//}
			}
		}
	}()
}

func GetArrStr2Nums(nums string) []int {
	arr := strings.Split(nums, ",")
	arrN := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		arrN[i] = strconvEx.StrTry2Int(arr[i], 0)
	}
	return arrN
}

func GetNums(t int) string {
	if t == 5 {
		return Get5Result()
	} else if t == 10 {
		return Get10Result()
	} else if t == 20 {
		return Get20Result()
	}
	return ""
}

func Get5Result() string {
	nStr := ""
	for i := 0; i < 5; i++ {
		n := timeKit.GlobaRand.Intn(10)
		nStr = nStr + fmt.Sprintf(",%d", n)
	}

	return nStr[1:]
}

func Get10Result() string {
	nStr := ""
	mpN := make(map[int]int)
	for i := 0; i < 10; i++ {
		n := timeKit.GlobaRand.Intn(10) + 1
		for j := 0; j < 10; j++ {
			if mpN[n] == 0 {
				mpN[n] = 1
				break
			} else {
				n++
				if n == 11 {
					n = 1
				}
			}
		}
		nStr = nStr + fmt.Sprintf(",%d", n)
	}
	return nStr[1:]
}

func Get20Result() string {
	nStr := ""
	mpN := make(map[int]int)
	for i := 0; i < 8; i++ {
		n := timeKit.GlobaRand.Intn(20) + 1
		for j := 0; j < 20; j++ {
			if mpN[n] == 0 {
				mpN[n] = 1
				break
			} else {
				n++
				if n == 21 {
					n = 1
				}
			}
		}
		nStr = nStr + fmt.Sprintf(",%02d", n)
	}
	return nStr[1:]
}
