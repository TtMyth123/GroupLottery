package Game28ResultKit

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/Game28ResultKit/Game28Const"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit/lotteryKit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/timeKit"
)

var (
	MpBs map[int]int
)

type Zg28Result struct {
	AwardNumbers string
	Nums         []int
	HS           int
	StrHS        string
}

type BetWinLost struct {
	IsWin    bool
	BetM     int
	WinM     float64
	Odds     float64
	OddsType int
}

func init() {
	MpBs = make(map[int]int)
	MpBs[1] = Game28Const.TO_G28_025
	MpBs[4] = Game28Const.TO_G28_025
	MpBs[7] = Game28Const.TO_G28_025
	MpBs[10] = Game28Const.TO_G28_025
	MpBs[16] = Game28Const.TO_G28_025
	MpBs[19] = Game28Const.TO_G28_025
	MpBs[22] = Game28Const.TO_G28_025
	MpBs[25] = Game28Const.TO_G28_025

	MpBs[2] = Game28Const.TO_G28_026
	MpBs[5] = Game28Const.TO_G28_026
	MpBs[8] = Game28Const.TO_G28_026
	MpBs[11] = Game28Const.TO_G28_026
	MpBs[17] = Game28Const.TO_G28_026
	MpBs[20] = Game28Const.TO_G28_026
	MpBs[23] = Game28Const.TO_G28_026
	MpBs[26] = Game28Const.TO_G28_026

	MpBs[3] = Game28Const.TO_G28_024
	MpBs[6] = Game28Const.TO_G28_024
	MpBs[9] = Game28Const.TO_G28_024
	MpBs[12] = Game28Const.TO_G28_024
	MpBs[15] = Game28Const.TO_G28_024
	MpBs[18] = Game28Const.TO_G28_024
	MpBs[21] = Game28Const.TO_G28_024
	MpBs[24] = Game28Const.TO_G28_024
}

func GetGame28AwardInfo(aLoAwardInfo models.LoAwardInfo, StopTime, l int) Game28AwardInfo {
	aGame28AwardInfo := Game28AwardInfo{}
	nums := lotteryKit.GetStrNum2Arr(aLoAwardInfo.ResultNums)
	if len(nums) < 3 {
		return aGame28AwardInfo
	}
	sum := GetAwardResultHS(nums)

	aGame28AwardInfo.LotteryStr = aLoAwardInfo.LotteryStr
	aGame28AwardInfo.NextLotteryStr = aLoAwardInfo.NextLotteryStr
	aGame28AwardInfo.ResultNums = aLoAwardInfo.ResultNums

	aGame28AwardInfo.StrLotteryTime = aLoAwardInfo.CurLotteryTime.Format(timeKit.DateTimeLayout)
	aGame28AwardInfo.StrNextLotteryTime = aLoAwardInfo.NextLotteryTime.Format(timeKit.DateTimeLayout)

	aGame28AwardInfo.ResultDS = GetAwardResultDS(sum)
	aGame28AwardInfo.ResultDX = GetAwardResultDX(sum)
	aGame28AwardInfo.ResultLH = GetAwardResultLH(nums)

	if l == 1 {
		aGame28AwardInfo.ResultGZH = fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s",
			aGame28AwardInfo.ResultDX,
			aGame28AwardInfo.ResultDS,
			aGame28AwardInfo.ResultLH,
			GetAwardResultBS(sum),
			GetAwardResultDX(sum)+GetAwardResultDS(sum),
			GetAwardResultJS(sum),
			GetAwardResultBDS(nums),
		)
	} else {
		aGame28AwardInfo.ResultGZH = GetAwardResultBS(sum)
	}

	t := time.Now()
	aGame28AwardInfo.Countdown = int(aLoAwardInfo.NextLotteryTime.Sub(t) / time.Second)
	aGame28AwardInfo.StopCountdown = aGame28AwardInfo.Countdown - StopTime

	return aGame28AwardInfo
}

func GetGame28HistoryResult(aLoAwardInfo models.LoAwardInfo) Game28HistoryResult {
	aGame28AwardInfo := Game28HistoryResult{}
	nums := lotteryKit.GetStrNum2Arr(aLoAwardInfo.ResultNums)
	sum := GetAwardResultHS(nums)

	aGame28AwardInfo.LotteryStr = aLoAwardInfo.LotteryStr
	aGame28AwardInfo.NextLotteryStr = aLoAwardInfo.NextLotteryStr
	aGame28AwardInfo.StrLotteryTime = aLoAwardInfo.CurLotteryTime.Format(timeKit.DateTimeLayout)
	aGame28AwardInfo.StrNextLotteryTime = aLoAwardInfo.NextLotteryTime.Format(timeKit.DateTimeLayout)

	aGame28AwardInfo.ResultNums = aLoAwardInfo.ResultNums
	aGame28AwardInfo.HS = sum
	aGame28AwardInfo.LH = GetAwardResultLH(nums)
	aGame28AwardInfo.BDS = GetAwardResultBDS(nums)

	aGame28AwardInfo.BS = GetAwardResultBS(sum)
	aGame28AwardInfo.DX = GetAwardResultDX(sum)
	aGame28AwardInfo.DS = GetAwardResultDS(sum)
	aGame28AwardInfo.JS = GetAwardResultJS(sum)

	return aGame28AwardInfo
}

func GetGame28HistoryResult2(aLoAwardInfo models.LoAwardInfo) Game28HistoryResult2 {
	aGame28AwardInfo := Game28HistoryResult2{}
	nums := lotteryKit.GetStrNum2Arr(aLoAwardInfo.ResultNums)
	sum := GetAwardResultHS(nums)

	aGame28AwardInfo.Period = aLoAwardInfo.LotteryNum
	aGame28AwardInfo.StrPeriod = aLoAwardInfo.LotteryStr

	aGame28AwardInfo.Nums = nums
	aGame28AwardInfo.SumNum = sum
	aGame28AwardInfo.DX = GetAwardResultDX(sum)
	aGame28AwardInfo.DXDS = aGame28AwardInfo.DX + GetAwardResultDS(sum)
	aGame28AwardInfo.JZ = GetAwardResultJS(sum)
	aGame28AwardInfo.BDS = GetAwardResultBDS(nums)
	aGame28AwardInfo.LH = GetAwardResultLH(nums)
	aGame28AwardInfo.BS = GetAwardResultBS(sum)

	return aGame28AwardInfo
}

/**
极大，极小
*/
func GetAwardResultJS(HS int) string {
	str := ""
	if HS <= 5 {
		str = "极小"
	} else if HS >= 22 {
		str = "极大"
	}

	return str
}
func GetAwardResultDS(HS int) string {
	str := "单"
	if HS%2 == 0 {
		str = "双"
	}
	return str
}

func GetAwardResultDX(HS int) string {
	str := "小"
	if HS >= 14 {
		str = "大"
	}
	return str
}

func isShunZ(Nums []int) bool {
	newNums := getSortNums(Nums)
	if newNums[0]+2 == newNums[2] && newNums[1]+1 == newNums[2] {
		return true
	} else if newNums[0] == 0 && newNums[1] == 1 && newNums[2] == 9 {
		return true
	} else if newNums[0] == 0 && newNums[1] == 8 && newNums[2] == 9 {
		return true
	}
	return false
}

/**
豹子，对子，顺子（顺暂无）
*/
func GetAwardResultBDS(Nums []int) string {
	str := ""
	if len(Nums) == 3 {
		if Nums[0] == Nums[1] && Nums[1] == Nums[2] {
			str = "豹子"
		} else if Nums[0] == Nums[1] || Nums[0] == Nums[2] || Nums[1] == Nums[2] {
			str = "对子"
		} else if isShunZ(Nums) {
			str = "顺子"
		}
	}
	return str
}

func GetAwardResultLH(Nums []int) string {
	str := "和"
	if Nums[0] < Nums[2] {
		str = "虎"
	} else if Nums[0] > Nums[2] {
		str = "龙"
	}
	return str
}

func GetZg28Result(AwardNumbers string) Zg28Result {
	aZg28Result := Zg28Result{AwardNumbers: AwardNumbers}

	aNums := lotteryKit.GetStrNum2Arr(AwardNumbers)
	aHS := GetAwardResultHS(aNums)
	aZg28Result.Nums = aNums
	aZg28Result.HS = aHS
	aZg28Result.StrHS = strconv.Itoa(aHS)

	return aZg28Result
}

func GetAwardResultBS(HS int) string {
	t := MpBs[HS]
	switch t {
	case Game28Const.TO_G28_024:
		return "红波"
	case Game28Const.TO_G28_026:
		return "蓝波"
	case Game28Const.TO_G28_025:
		return "绿波"
	}
	return ""
}

func GetAwardResultHS(Nums []int) int {
	s := 0
	for _, v := range Nums {
		s += v
	}
	return s
}
func GetAwardNumbers(AwardNumbers string) []int {
	arrStrNum := strings.Split(AwardNumbers, ",")
	arrIntNum := make([]int, len(arrStrNum))

	for i := 0; i < len(arrStrNum); i++ {
		arrIntNum[i] = strconvEx.StrTry2Int(arrStrNum[i], 0)
	}
	return arrIntNum
}

func ComputeLoseWin(bet models.LoBetInfo, aZg28Result Zg28Result, mpLoOddsInfo map[int]models.LoOddsInfo) BetWinLost {
	aBetWinLost := BetWinLost{IsWin: false, WinM: 0, BetM: bet.BetM, OddsType: bet.OddsType, Odds: bet.Odds}
	fBetM := float64(bet.BetM)
	aLoOddsInfo := mpLoOddsInfo[bet.OddsType]
	switch aLoOddsInfo.OddsType {
	case Game28Const.TO_G28_011: //大
		{
			if aZg28Result.HS >= 14 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_012: //小
		{
			if aZg28Result.HS < 14 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_013: //单
		{
			if aZg28Result.HS%2 == 1 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_014: //双
		{
			if aZg28Result.HS%2 == 0 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_015: //小单
		{
			if aZg28Result.HS < 14 && aZg28Result.HS%2 == 1 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_016: //小双
		{
			if aZg28Result.HS < 14 && aZg28Result.HS%2 == 0 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_017: //大单
		{
			if aZg28Result.HS >= 14 && aZg28Result.HS%2 == 1 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_018: //大双
		{
			if aZg28Result.HS >= 14 && aZg28Result.HS%2 == 0 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_019: //极大
		{
			if aZg28Result.HS >= 22 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_020: //极小
		{
			if aZg28Result.HS <= 5 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_021: //对子
		{
			if (aZg28Result.Nums[0] == aZg28Result.Nums[1] && aZg28Result.Nums[0] != aZg28Result.Nums[2]) ||
				(aZg28Result.Nums[0] == aZg28Result.Nums[2] && aZg28Result.Nums[0] != aZg28Result.Nums[1]) ||
				(aZg28Result.Nums[1] == aZg28Result.Nums[2] && aZg28Result.Nums[1] != aZg28Result.Nums[0]) {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_022: //顺子
		{
			newNums := getSortNums(aZg28Result.Nums)
			if newNums[0]+2 == newNums[2] && newNums[1]+1 == newNums[2] {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			} else if newNums[0] == 0 && newNums[1] == 1 && newNums[2] == 9 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			} else if newNums[0] == 0 && newNums[1] == 8 && newNums[2] == 9 {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}
	case Game28Const.TO_G28_023: //豹子
		{
			if aZg28Result.Nums[0] == aZg28Result.Nums[1] && aZg28Result.Nums[1] == aZg28Result.Nums[2] {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}
	case Game28Const.TO_G28_024, Game28Const.TO_G28_025, Game28Const.TO_G28_026: //红波,绿波,蓝波
		{
			if MpBs[aZg28Result.HS] == bet.OddsType {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_027: //龙
		{
			if aZg28Result.Nums[0] > aZg28Result.Nums[2] {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_028: //虎
		{
			if aZg28Result.Nums[0] < aZg28Result.Nums[2] {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_029: //和
		{
			if aZg28Result.Nums[0] == aZg28Result.Nums[2] {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_031: //闲1
		{
			ZNum := aZg28Result.HS % 10
			if aZg28Result.Nums[0] > ZNum {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_032: //闲2
		{
			ZNum := aZg28Result.HS % 10
			if aZg28Result.Nums[1] > ZNum {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_033: //闲3
		{
			ZNum := aZg28Result.HS % 10
			if aZg28Result.Nums[2] > ZNum {
				aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
				aBetWinLost.IsWin = true
			}
		}

	case Game28Const.TO_G28_034: //庄
		{
			ZNum := aZg28Result.HS % 10
			for i := 0; i < 3; i++ {
				newBet := fBetM / 3
				if aZg28Result.Nums[i] <= ZNum {
					aBetWinLost.WinM += strconvEx.Decimal(newBet * aBetWinLost.Odds)
				}
			}
		}

	default: //数字:0~27
		num := bet.OddsType % 100
		if num == aZg28Result.HS {
			aBetWinLost.WinM = strconvEx.Decimal(fBetM * aBetWinLost.Odds)
			aBetWinLost.IsWin = true
		}
	}

	return aBetWinLost
}

func getSortNums(nums []int) []int {
	iLen := len(nums)
	newNums := make([]int, iLen)
	for i := 0; i < iLen; i++ {
		newNums[i] = nums[i]
	}
	for i := 0; i < iLen; i++ {
		for j := 0; j < iLen-i-1; j++ {
			if newNums[j] > newNums[j+1] {
				t := newNums[j]
				newNums[j] = newNums[j+1]
				newNums[j+1] = t
			}
		}
	}
	return newNums
}
