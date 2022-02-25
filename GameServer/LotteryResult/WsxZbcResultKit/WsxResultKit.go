package WsxZbcResultKit

import (
	"strings"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxBox"
	"ttmyth123/GroupLottery/GameServer/LotteryResult/WsxConst"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/stringKit"
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

func GetMapNum(Nums string) map[string]int {
	mapNums := make(map[string]int)
	arrN := strings.Split(Nums, ",")
	for _, n := range arrN {
		mapNums[n] = 1
	}
	return mapNums
}
func GetMapNumByArr(Nums []string) map[string]int {
	mapNums := make(map[string]int)
	for _, n := range Nums {
		mapNums[n] = 1
	}
	return mapNums
}

func GetBetOkCount(betNum, openNum map[string]int) int {
	s := 0
	for n, _ := range betNum {
		if openNum[n] == 1 {
			s++
		}
	}
	return s
}

func ComputeLoseWin(bet models.LoBetInfo, aOpenCodeInfo WsxBox.OpenCodeZbcInfo) (float64, float64) {
	win := 0.0
	Odds := 0.0
	switch bet.BigOddsType {
	case WsxConst.TO_BigType_ZB01, WsxConst.TO_BigType_ZB02, WsxConst.TO_BigType_ZB03, WsxConst.TO_BigType_ZB04,
		WsxConst.TO_BigType_ZB05, WsxConst.TO_BigType_ZB06, WsxConst.TO_BigType_ZB07, WsxConst.TO_BigType_ZB08,
		WsxConst.TO_BigType_ZB09, WsxConst.TO_BigType_ZB10:
		{
			mpOdds := make(map[int]float64)
			stringKit.GetJsonObj(bet.StrOdds, &mpOdds)

			betNum := GetMapNum(bet.Nums)

			openNum := GetMapNumByArr(aOpenCodeInfo.Nums)
			s := GetBetOkCount(betNum, openNum)
			if Odds1, ok := mpOdds[s]; ok {
				win = strconvEx.Decimal(float64(bet.BetM) * Odds1)
				Odds = Odds1
			}
		}
	case WsxConst.TO_BigType_ZB11: //大小
		{
			dCount := 0 //大的数量
			xCount := 0 //小的数量
			for _, num := range aOpenCodeInfo.Nums {
				iNum := strconvEx.StrTry2Int(num, 0)
				if iNum <= 40 {
					xCount++
				} else {
					dCount++
				}
			}
			if bet.Nums == "1" && xCount > dCount {
				Odds = bet.Odds
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			} else if bet.Nums == "2" && xCount < dCount {
				Odds = bet.Odds
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}

			//iLen := len(aOpenCodeInfo.Nums)
			//Odds = bet.Odds
			//if iLen > 0 {
			//	iNum := strconvEx.StrTry2Int(aOpenCodeInfo.Nums[iLen-1], 0)
			//	iNum = iNum % 10
			//	if iNum <= 4 && bet.Nums == "1" {
			//		Odds = bet.Odds
			//		win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			//	} else if iNum > 4 && bet.Nums == "2" {
			//		Odds = bet.Odds
			//		win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			//	}
			//}
		}
	case WsxConst.TO_BigType_ZB12: //单双
		{
			dCount := 0 //单的数量
			sCount := 0 //双的数量
			for _, num := range aOpenCodeInfo.Nums {
				iNum := strconvEx.StrTry2Int(num, 0)
				if iNum%2 == 0 {
					sCount++
				} else {
					dCount++
				}
			}
			if bet.Nums == "1" && sCount < dCount {
				Odds = bet.Odds
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			} else if bet.Nums == "2" && sCount > dCount {
				Odds = bet.Odds
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}

			//iLen := len(aOpenCodeInfo.Nums)
			//Odds = bet.Odds
			//if iLen > 0 {
			//	iNum := strconvEx.StrTry2Int(aOpenCodeInfo.Nums[iLen-1], 0)
			//	iNum = iNum % 2
			//	if iNum == 1 && bet.Nums == "1" {
			//		Odds = bet.Odds
			//		win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			//	} else if iNum == 0 && bet.Nums == "2" {
			//		Odds = bet.Odds
			//		win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			//	}
			//}
		}

	case WsxConst.TO_BigType_ZB13: //单双
		{
			dCount := 0 //大的数量
			xCount := 0 //小的数量
			for _, num := range aOpenCodeInfo.Nums {
				iNum := strconvEx.StrTry2Int(num, 0)
				if iNum <= 40 {
					xCount++
				} else {
					dCount++
				}
			}
			if xCount == dCount {
				Odds = bet.Odds
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}

	case WsxConst.TO_BigType_ZB14: //单双和
		{
			dCount := 0 //单的数量
			sCount := 0 //双的数量
			for _, num := range aOpenCodeInfo.Nums {
				iNum := strconvEx.StrTry2Int(num, 0)
				if iNum%2 == 0 {
					sCount++
				} else {
					dCount++
				}
			}
			if sCount == dCount {
				Odds = bet.Odds
				win = strconvEx.Decimal(float64(bet.BetM) * bet.Odds)
			}
		}
	}

	return win, Odds
}

func GetResultNums(aOpenCodeInfo WsxBox.HttpOpenCodeZbc) map[string]interface{} {
	aResultNums := make(map[string]interface{})
	aResultNums["Nums"] = aOpenCodeInfo.Nums
	return aResultNums
}

func GetStr2ResultNums(strJson string) map[string]interface{} {
	aResultNums := make(map[string]interface{})
	stringKit.GetJsonObj(strJson, &aResultNums)

	return aResultNums
}
