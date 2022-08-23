package WsxBbcResultKit

import (
	"github.com/TtMyth123/GameServer/LotteryResult/WsxBox"
	"github.com/TtMyth123/GameServer/LotteryResult/WsxConst"
	"github.com/TtMyth123/GameServer/models"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"strings"
)

var (
	mpBsNum map[string]string
)

func init() {
	mpBsNum = make(map[string]string)
	strBsNum0 := ",00,01,08,09,16,17,24,25,32,33,40,41,48,49,56,57,64,65,72,73,80,81,88,89,96,"
	strBsNum1 := ",02,03,10,11,18,19,26,27,34,35,42,43,50,51,58,59,66,67,74,75,82,83,90,91,97,"
	strBsNum2 := ",04,05,12,13,20,21,28,29,36,37,44,45,52,53,60,61,68,69,76,77,84,85,92,93,98,"
	strBsNum3 := ",06,07,14,15,22,23,30,31,38,39,46,47,54,55,62,63,70,71,78,79,86,87,94,95,99,"

	mpBsNum["0"] = strBsNum0
	mpBsNum["1"] = strBsNum1
	mpBsNum["2"] = strBsNum2
	mpBsNum["3"] = strBsNum3
}

func GetEndStr(str string, n int) string {
	iLen := len(str)
	if iLen > n {
		i := iLen - n
		return str[i:]
	}
	return str
}

func GetBeginStr(str string, n int) string {
	iLen := len(str)
	if iLen > n {
		return str[:n]
	}
	return str
}

func ComputeLoseWin(bet models.LoBetInfo, aOpenCodeInfo WsxBox.OpenCodeInfo, mpLoOddsInfo map[int]models.LoOddsInfo) float64 {
	//BigType_01 = 1//头特大小
	//BigType_02 = 2//头特单双
	//BigType_03 = 3//一等奖大小
	//BigType_04 = 4//一等奖单双
	//BigType_05 = 5//头特
	//BigType_06 = 6//尾特
	//BigType_07 = 7//头等特码
	//BigType_08 = 8//一等特码
	//BigType_09 = 9//二等特码
	//BigType_10 = 10//二连位
	//BigType_11 = 11//三连位
	//BigType_12 = 12//平码两位区
	//BigType_13 = 13//平码三位区
	//BigType_14 = 14//波色
	win := 0.0
	aLoOddsInfo := mpLoOddsInfo[bet.OddsType]
	switch aLoOddsInfo.BigType {
	case WsxConst.BigType_01: //头特大小
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 2)
			n := strconvEx.StrTry2Int(strN, 0)
			if n <= 49 && aLoOddsInfo.SNum == "1" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			} else if n > 49 && aLoOddsInfo.SNum == "2" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_02: //头特单双
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 1)
			n := strconvEx.StrTry2Int(strN, 0)
			if n%2 == 1 && aLoOddsInfo.SNum == "1" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			} else if n%2 == 0 && aLoOddsInfo.SNum == "2" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_03: //一等奖大小
		{
			strN := GetEndStr(aOpenCodeInfo.FirstNum, 2)
			n := strconvEx.StrTry2Int(strN, 0)
			if n <= 49 && aLoOddsInfo.SNum == "1" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			} else if n > 49 && aLoOddsInfo.SNum == "2" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_04: //一等奖单双
		{
			strN := GetEndStr(aOpenCodeInfo.FirstNum, 1)
			n := strconvEx.StrTry2Int(strN, 0)
			if n%2 == 1 && aLoOddsInfo.SNum == "1" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			} else if n%2 == 0 && aLoOddsInfo.SNum == "2" {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_05, WsxConst.BigType_05B: //头等-头特
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 2)
			strN = GetBeginStr(strN, 1)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}
	case WsxConst.BigType_06, WsxConst.BigType_06B: //头等-尾特
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 1)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}
	case WsxConst.BigType_07, WsxConst.BigType_07B: //头等特码
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 2)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_08, WsxConst.BigType_08B: //一等特码
		{
			strN := GetEndStr(aOpenCodeInfo.FirstNum, 2)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_09, WsxConst.BigType_09B: //二等特码
		{
			nums := aOpenCodeInfo.SecondNum
			for _, v := range nums {
				strN := GetEndStr(v, 2)
				if aLoOddsInfo.SNum == strN {
					win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
					break
				}
			}
		}

	case WsxConst.BigType_10: //二连位
		{
			nums := make([]string, 0)
			nums = append(nums, aOpenCodeInfo.Jackpots, aOpenCodeInfo.FirstNum)
			nums = append(nums, aOpenCodeInfo.SecondNum...)
			nums = append(nums, aOpenCodeInfo.ThirdNum...)
			nums = append(nums, aOpenCodeInfo.ForthNum...)
			nums = append(nums, aOpenCodeInfo.FifthNum...)
			nums = append(nums, aOpenCodeInfo.SixthNum...)
			nums = append(nums, aOpenCodeInfo.SeventhNum...)
			//nums = append(nums,aOpenCodeInfo.EighthNum...)

			win = 0
			for _, v := range nums {
				strN := GetEndStr(v, 2)
				if aLoOddsInfo.SNum == strN {
					win += strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
					//break
				}
			}
		}
	case WsxConst.BigType_11: //头等 三连位
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 3)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}
	case WsxConst.BigType_22: //一等 三连位
		{
			strN := GetEndStr(aOpenCodeInfo.FirstNum, 3)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_12, WsxConst.BigType_13: //平码两位区
		{
			BetNums := strings.Split(bet.Nums, ",")
			mpBetNum := make(map[string]int)
			for _, betN := range BetNums {
				mpBetNum[betN] = 1
			}

			nums := make([]string, 0)
			nums = append(nums, aOpenCodeInfo.Jackpots, aOpenCodeInfo.FirstNum)
			nums = append(nums, aOpenCodeInfo.SecondNum...)
			nums = append(nums, aOpenCodeInfo.ThirdNum...)
			nums = append(nums, aOpenCodeInfo.ForthNum...)
			nums = append(nums, aOpenCodeInfo.FifthNum...)
			nums = append(nums, aOpenCodeInfo.SixthNum...)
			nums = append(nums, aOpenCodeInfo.SeventhNum...)
			//nums = append(nums,aOpenCodeInfo.EighthNum...)

			betNumLen := len(BetNums)
			sumI := 0

			for betN, _ := range mpBetNum {
				for _, v := range nums {
					strN := GetEndStr(v, 2)
					if betN == strN {
						sumI++
						break
					}
				}
				if betNumLen == sumI {
					win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
					break
				}
			}
			//if betNumLen==sumI {
			//	win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			//}

		}

	case WsxConst.BigType_15: //头等特码 波色
		{
			strN := GetEndStr(aOpenCodeInfo.Jackpots, 2)
			if strings.Contains(mpBsNum[aLoOddsInfo.SNum], strN) {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}
	case WsxConst.BigType_16: //一等特码 波色
		{

			strN := GetEndStr(aOpenCodeInfo.FirstNum, 2)
			if strings.Contains(mpBsNum[aLoOddsInfo.SNum], strN) {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.BigType_17: //二等特码 波色
		{
			for _, Nums := range aOpenCodeInfo.SecondNum {
				strN := GetEndStr(Nums, 2)
				if strings.Contains(mpBsNum[aLoOddsInfo.SNum], strN) {
					win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
					break
				}
			}
		}
	case WsxConst.BigType_18, WsxConst.BigType_18B: //一等-头特
		{
			strN := GetEndStr(aOpenCodeInfo.FirstNum, 2)
			strN = GetBeginStr(strN, 1)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}
	case WsxConst.BigType_19, WsxConst.BigType_19B: //一等-尾特
		{
			strN := GetEndStr(aOpenCodeInfo.FirstNum, 1)
			if aLoOddsInfo.SNum == strN {
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

		//case WsxConst.BigType_20,WsxConst.BigType_20B: //二等-头特
		//	{
		//		strN := GetBeginStr(aOpenCodeInfo.SecondNum,1)
		//		if aLoOddsInfo.SNum==strN {
		//			win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
		//		}
		//	}

		//case WsxConst.BigType_21,WsxConst.BigType_21B: //二等-尾特
		//	{
		//		strN := GetEndStr(aOpenCodeInfo.SecondNum,1)
		//		if aLoOddsInfo.SNum==strN {
		//			win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
		//		}
		//	}

	}

	return win
}

func GetResultNums(aOpenCodeInfo WsxBox.HttpOpenCodeBbc) map[string]interface{} {
	aResultNums := make(map[string]interface{})
	aResultNums["jackpots"] = aOpenCodeInfo.Jackpots
	aResultNums["firstNum"] = aOpenCodeInfo.FirstNum
	aResultNums["secondNum"] = aOpenCodeInfo.SecondNum
	aResultNums["thirdNum"] = aOpenCodeInfo.ThirdNum
	aResultNums["forthNum"] = aOpenCodeInfo.ForthNum
	aResultNums["fifthNum"] = aOpenCodeInfo.FifthNum
	aResultNums["sixthNum"] = aOpenCodeInfo.SixthNum
	aResultNums["seventhNum"] = aOpenCodeInfo.SeventhNum
	//aResultNums["eighthNum"] = aOpenCodeInfo.EighthNum
	return aResultNums
}

func GetStr2ResultNums(strJson string) map[string]interface{} {
	aResultNums := make(map[string]interface{})
	stringKit.GetJsonObj(strJson, &aResultNums)

	return aResultNums
}
