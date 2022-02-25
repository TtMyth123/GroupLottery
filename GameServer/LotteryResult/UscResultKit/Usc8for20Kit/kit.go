package Usc8for20Kit

import (
	"errors"
	"fmt"
	"strconv"
	"ttmyth123/kit/lotteryKit"
)

type Award struct {
	NumList []int
	StrNums string
	ZH      int
}

func GetZHDX(ZH int) int {
	if ZH >= 85 && ZH <= 132 {
		return 2
	} else if ZH >= 36 && ZH <= 83 {
		return 1
	} else {
		return 0
	}
}

func GetResultGZH(NumList []int) string {
	gzh := 0
	for i := 0; i < len(NumList); i++ {
		gzh += NumList[i]
	}
	t := GetZHDX(gzh)
	dx := "小"
	if t == 2 {
		dx = "大"
	}

	ds := "单"
	if GetZHDS(gzh) == 2 {
		ds = "双"
	}

	return fmt.Sprintf("%d,%s,%s", gzh, dx, ds)
}

func GetZHDS(ZH int) int {
	if ZH%2 == 0 {
		return 2
	} else {
		return 1
	}
}

func getSum(NumList []int) int {
	s := 0
	for _, v := range NumList {
		s += v
	}
	return s
}

/**
尾数大小，1：小。2：大
*/
func getWSDX(num int) int {
	n := num % 10
	if n <= 4 {
		return 1
	} else {
		return 2
	}
}

/**
大小 1：小。2：大
*/
func GetNumDX(num int) int {
	if num > 10 {
		return 2
	} else {
		return 1
	}
}
func GetResultDXs(NumList []int) string {
	ResultDX := ""
	for i := 0; i < len(NumList); i++ {
		t := GetNumDX(NumList[i])
		s := "小"
		if t == 2 {
			s = "大"
		}
		ResultDX = ResultDX + "," + s
	}
	if len(ResultDX) > 0 {
		ResultDX = ResultDX[1:]
	}
	return ResultDX
}

/**
合 单双，1：小。2：大
*/
func getHDS(num int) int {
	n1 := int(num / 10)
	n2 := num % 10
	n := n1 + n2
	return GetNumDS(n)
}

/**
单双
*/
func GetNumDS(num int) int {
	if num%2 == 0 {
		return 2
	} else {
		return 1
	}
}

func GetResultDSs(NumList []int) string {
	result := ""
	for i := 0; i < len(NumList); i++ {
		t := GetNumDS(NumList[i])
		s := "单"
		if t == 2 {
			s = "双"
		}
		result = result + "," + s
	}
	if len(result) > 0 {
		result = result[1:]
	}
	return result
}

func GetResultLHs(NumList []int) string {
	result := ""
	iLen := len(NumList)
	for i := 0; i < iLen/2; i++ {
		s := ""
		if NumList[i] > NumList[iLen-1-i] {
			s = "龙"
		} else if NumList[i] < NumList[iLen-1-i] {
			s = "虎"
		}
		result = result + "," + s
	}
	if len(result) > 0 {
		result = result[1:]
	}
	return result
}

/**
●中：开出之号码为01、02、03、04、05、06、07
中
*/
func GetNumZ(num int) int {
	if num >= 1 && num <= 1 {
		return 2
	} else {
		return 1
	}
}

/**
●发：开出之号码为08、09、10、11、12、13、14
发
*/
func GetNumF(num int) int {
	if num >= 8 && num <= 14 {
		return 2
	} else {
		return 1
	}
}

/**
●白：开出之号码为15、16、17、18、19、20
白
*/
func GetNumB(num int) int {
	if num >= 15 && num <= 20 {
		return 2
	} else {
		return 1
	}
}

/**
●东：开出之号码为01、05、09、13、17
东
*/
func GetNumDF(num int) int {
	if num%4 == 1 {
		return 2
	} else {
		return 1
	}
}

/**
●南：开出之号码为02、06、10、14、18
南
*/
func GetNumNF(num int) int {
	if num%4 == 2 {
		return 2
	} else {
		return 1
	}
}

/**
●西：开出之号码为03、07、11、15、19
西
*/
func GetNumXF(num int) int {
	if num%4 == 3 {
		return 2
	} else {
		return 1
	}
}

/**
●北：开出之号码为04、08、12、16、20
北
*/
func GetNumBF(num int) int {
	if num%4 == 0 {
		return 2
	} else {
		return 1
	}
}
func GetAward(strNums string) (Award, error) {
	var NumList = lotteryKit.GetStrNum2Arr(strNums)
	aAward := Award{
		NumList: NumList,
	}
	StrNums := ""
	for i := 0; i < len(NumList); i++ {
		StrNums = StrNums + "," + strconv.Itoa(NumList[i])
	}
	aAward.StrNums = StrNums

	aAward.ZH = getSum(NumList)
	return aAward, nil
}

func GetWinMoney(aAward Award, oddsType, betM int, odds float64) float64 {
	winM := float64(0)
	switch oddsType {
	case OT_001:
		if aAward.NumList[0] == 1 {
			winM = float64(betM) * odds
		}
	case OT_002:
		if aAward.NumList[0] == 2 {
			winM = float64(betM) * odds
		}
	case OT_003:
		if aAward.NumList[0] == 3 {
			winM = float64(betM) * odds
		}
	case OT_004:
		if aAward.NumList[0] == 4 {
			winM = float64(betM) * odds
		}
	case OT_005:
		if aAward.NumList[0] == 5 {
			winM = float64(betM) * odds
		}
	case OT_006:
		if aAward.NumList[0] == 6 {
			winM = float64(betM) * odds
		}
	case OT_007:
		if aAward.NumList[0] == 7 {
			winM = float64(betM) * odds
		}
	case OT_008:
		if aAward.NumList[0] == 8 {
			winM = float64(betM) * odds
		}
	case OT_009:
		if aAward.NumList[0] == 9 {
			winM = float64(betM) * odds
		}
	case OT_010:
		if aAward.NumList[0] == 10 {
			winM = float64(betM) * odds
		}
	case OT_011:
		if aAward.NumList[0] == 11 {
			winM = float64(betM) * odds
		}
	case OT_012:
		if aAward.NumList[0] == 12 {
			winM = float64(betM) * odds
		}
	case OT_013:
		if aAward.NumList[0] == 13 {
			winM = float64(betM) * odds
		}
	case OT_014:
		if aAward.NumList[0] == 14 {
			winM = float64(betM) * odds
		}
	case OT_015:
		if aAward.NumList[0] == 15 {
			winM = float64(betM) * odds
		}
	case OT_016:
		if aAward.NumList[0] == 16 {
			winM = float64(betM) * odds
		}
	case OT_017:
		if aAward.NumList[0] == 17 {
			winM = float64(betM) * odds
		}
	case OT_018:
		if aAward.NumList[0] == 18 {
			winM = float64(betM) * odds
		}
	case OT_019:
		if aAward.NumList[0] == 19 {
			winM = float64(betM) * odds
		}
	case OT_020:
		if aAward.NumList[0] == 20 {
			winM = float64(betM) * odds
		}
	case OT_021:
		if GetNumDS(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_022:
		if GetNumDS(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_023:
		if GetNumDX(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_024:
		if GetNumDX(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_025:
		if getHDS(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_026:
		if getHDS(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_027:
		if getWSDX(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_028:
		if getWSDX(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_029:
		if GetNumDF(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_030:
		if GetNumNF(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_031:
		if GetNumXF(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_032:
		if GetNumBF(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_033:
		if GetNumZ(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_034:
		if GetNumF(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_035:
		if GetNumB(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_036:
		if aAward.NumList[0] > aAward.NumList[7] {
			winM = float64(betM) * odds
		}
	case OT_037:
		if aAward.NumList[0] < aAward.NumList[7] {
			winM = float64(betM) * odds
		}

	case OT_038:
		if aAward.NumList[1] == 1 {
			winM = float64(betM) * odds
		}
	case OT_039:
		if aAward.NumList[1] == 2 {
			winM = float64(betM) * odds
		}
	case OT_040:
		if aAward.NumList[1] == 3 {
			winM = float64(betM) * odds
		}
	case OT_041:
		if aAward.NumList[1] == 4 {
			winM = float64(betM) * odds
		}
	case OT_042:
		if aAward.NumList[1] == 5 {
			winM = float64(betM) * odds
		}
	case OT_043:
		if aAward.NumList[1] == 6 {
			winM = float64(betM) * odds
		}
	case OT_044:
		if aAward.NumList[1] == 7 {
			winM = float64(betM) * odds
		}
	case OT_045:
		if aAward.NumList[1] == 8 {
			winM = float64(betM) * odds
		}
	case OT_046:
		if aAward.NumList[1] == 9 {
			winM = float64(betM) * odds
		}
	case OT_047:
		if aAward.NumList[1] == 10 {
			winM = float64(betM) * odds
		}
	case OT_048:
		if aAward.NumList[1] == 11 {
			winM = float64(betM) * odds
		}
	case OT_049:
		if aAward.NumList[1] == 12 {
			winM = float64(betM) * odds
		}
	case OT_050:
		if aAward.NumList[1] == 13 {
			winM = float64(betM) * odds
		}
	case OT_051:
		if aAward.NumList[1] == 14 {
			winM = float64(betM) * odds
		}
	case OT_052:
		if aAward.NumList[1] == 15 {
			winM = float64(betM) * odds
		}
	case OT_053:
		if aAward.NumList[1] == 16 {
			winM = float64(betM) * odds
		}
	case OT_054:
		if aAward.NumList[1] == 17 {
			winM = float64(betM) * odds
		}
	case OT_055:
		if aAward.NumList[1] == 18 {
			winM = float64(betM) * odds
		}
	case OT_056:
		if aAward.NumList[1] == 19 {
			winM = float64(betM) * odds
		}
	case OT_057:
		if aAward.NumList[1] == 20 {
			winM = float64(betM) * odds
		}
	case OT_058:
		if GetNumDS(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_059:
		if GetNumDS(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_060:
		if GetNumDX(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_061:
		if GetNumDX(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_062:
		if getHDS(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_063:
		if getHDS(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_064:
		if getWSDX(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_065:
		if getWSDX(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_066:
		if GetNumDF(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_067:
		if GetNumNF(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_068:
		if GetNumXF(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_069:
		if GetNumBF(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_070:
		if GetNumZ(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_071:
		if GetNumF(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_072:
		if GetNumB(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_073:
		if aAward.NumList[1] > aAward.NumList[6] {
			winM = float64(betM) * odds
		}
	case OT_074:
		if aAward.NumList[1] < aAward.NumList[6] {
			winM = float64(betM) * odds
		}

	case OT_075:
		if aAward.NumList[2] == 1 {
			winM = float64(betM) * odds
		}
	case OT_076:
		if aAward.NumList[2] == 2 {
			winM = float64(betM) * odds
		}
	case OT_077:
		if aAward.NumList[2] == 3 {
			winM = float64(betM) * odds
		}
	case OT_078:
		if aAward.NumList[2] == 4 {
			winM = float64(betM) * odds
		}
	case OT_079:
		if aAward.NumList[2] == 5 {
			winM = float64(betM) * odds
		}
	case OT_080:
		if aAward.NumList[2] == 6 {
			winM = float64(betM) * odds
		}
	case OT_081:
		if aAward.NumList[2] == 7 {
			winM = float64(betM) * odds
		}
	case OT_082:
		if aAward.NumList[2] == 8 {
			winM = float64(betM) * odds
		}
	case OT_083:
		if aAward.NumList[2] == 9 {
			winM = float64(betM) * odds
		}
	case OT_084:
		if aAward.NumList[2] == 10 {
			winM = float64(betM) * odds
		}
	case OT_085:
		if aAward.NumList[2] == 11 {
			winM = float64(betM) * odds
		}
	case OT_086:
		if aAward.NumList[2] == 12 {
			winM = float64(betM) * odds
		}
	case OT_087:
		if aAward.NumList[2] == 13 {
			winM = float64(betM) * odds
		}
	case OT_088:
		if aAward.NumList[2] == 14 {
			winM = float64(betM) * odds
		}
	case OT_089:
		if aAward.NumList[2] == 15 {
			winM = float64(betM) * odds
		}
	case OT_090:
		if aAward.NumList[2] == 16 {
			winM = float64(betM) * odds
		}
	case OT_091:
		if aAward.NumList[2] == 17 {
			winM = float64(betM) * odds
		}
	case OT_092:
		if aAward.NumList[2] == 18 {
			winM = float64(betM) * odds
		}
	case OT_093:
		if aAward.NumList[2] == 19 {
			winM = float64(betM) * odds
		}
	case OT_094:
		if aAward.NumList[2] == 20 {
			winM = float64(betM) * odds
		}
	case OT_095:
		if GetNumDS(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_096:
		if GetNumDS(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_097:
		if GetNumDX(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_098:
		if GetNumDX(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_099:
		if getHDS(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_100:
		if getHDS(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_101:
		if getWSDX(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_102:
		if getWSDX(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_103:
		if GetNumDF(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_104:
		if GetNumNF(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_105:
		if GetNumXF(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_106:
		if GetNumBF(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_107:
		if GetNumZ(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_108:
		if GetNumF(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_109:
		if GetNumB(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_110:
		if aAward.NumList[2] > aAward.NumList[5] {
			winM = float64(betM) * odds
		}
	case OT_111:
		if aAward.NumList[2] < aAward.NumList[5] {
			winM = float64(betM) * odds
		}

	case OT_112:
		if aAward.NumList[3] == 1 {
			winM = float64(betM) * odds
		}
	case OT_113:
		if aAward.NumList[3] == 2 {
			winM = float64(betM) * odds
		}
	case OT_114:
		if aAward.NumList[3] == 3 {
			winM = float64(betM) * odds
		}
	case OT_115:
		if aAward.NumList[3] == 4 {
			winM = float64(betM) * odds
		}
	case OT_116:
		if aAward.NumList[3] == 5 {
			winM = float64(betM) * odds
		}
	case OT_117:
		if aAward.NumList[3] == 6 {
			winM = float64(betM) * odds
		}
	case OT_118:
		if aAward.NumList[3] == 7 {
			winM = float64(betM) * odds
		}
	case OT_119:
		if aAward.NumList[3] == 8 {
			winM = float64(betM) * odds
		}
	case OT_120:
		if aAward.NumList[3] == 9 {
			winM = float64(betM) * odds
		}
	case OT_121:
		if aAward.NumList[3] == 10 {
			winM = float64(betM) * odds
		}
	case OT_122:
		if aAward.NumList[3] == 11 {
			winM = float64(betM) * odds
		}
	case OT_123:
		if aAward.NumList[3] == 12 {
			winM = float64(betM) * odds
		}
	case OT_124:
		if aAward.NumList[3] == 13 {
			winM = float64(betM) * odds
		}
	case OT_125:
		if aAward.NumList[3] == 14 {
			winM = float64(betM) * odds
		}
	case OT_126:
		if aAward.NumList[3] == 15 {
			winM = float64(betM) * odds
		}
	case OT_127:
		if aAward.NumList[3] == 16 {
			winM = float64(betM) * odds
		}
	case OT_128:
		if aAward.NumList[3] == 17 {
			winM = float64(betM) * odds
		}
	case OT_129:
		if aAward.NumList[3] == 18 {
			winM = float64(betM) * odds
		}
	case OT_130:
		if aAward.NumList[3] == 19 {
			winM = float64(betM) * odds
		}
	case OT_131:
		if aAward.NumList[3] == 20 {
			winM = float64(betM) * odds
		}
	case OT_132:
		if GetNumDS(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_133:
		if GetNumDS(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_134:
		if GetNumDX(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_135:
		if GetNumDX(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_136:
		if getHDS(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_137:
		if getHDS(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_138:
		if getWSDX(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_139:
		if getWSDX(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_140:
		if GetNumDF(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_141:
		if GetNumNF(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_142:
		if GetNumXF(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_143:
		if GetNumBF(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_144:
		if GetNumZ(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_145:
		if GetNumF(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_146:
		if GetNumB(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_147:
		if aAward.NumList[3] > aAward.NumList[4] {
			winM = float64(betM) * odds
		}
	case OT_148:
		if aAward.NumList[3] < aAward.NumList[4] {
			winM = float64(betM) * odds
		}

	case OT_149:
		if aAward.NumList[4] == 1 {
			winM = float64(betM) * odds
		}
	case OT_150:
		if aAward.NumList[4] == 2 {
			winM = float64(betM) * odds
		}
	case OT_151:
		if aAward.NumList[4] == 3 {
			winM = float64(betM) * odds
		}
	case OT_152:
		if aAward.NumList[4] == 4 {
			winM = float64(betM) * odds
		}
	case OT_153:
		if aAward.NumList[4] == 5 {
			winM = float64(betM) * odds
		}
	case OT_154:
		if aAward.NumList[4] == 6 {
			winM = float64(betM) * odds
		}
	case OT_155:
		if aAward.NumList[4] == 7 {
			winM = float64(betM) * odds
		}
	case OT_156:
		if aAward.NumList[4] == 8 {
			winM = float64(betM) * odds
		}
	case OT_157:
		if aAward.NumList[4] == 9 {
			winM = float64(betM) * odds
		}
	case OT_158:
		if aAward.NumList[4] == 10 {
			winM = float64(betM) * odds
		}
	case OT_159:
		if aAward.NumList[4] == 11 {
			winM = float64(betM) * odds
		}
	case OT_160:
		if aAward.NumList[4] == 12 {
			winM = float64(betM) * odds
		}
	case OT_161:
		if aAward.NumList[4] == 13 {
			winM = float64(betM) * odds
		}
	case OT_162:
		if aAward.NumList[4] == 14 {
			winM = float64(betM) * odds
		}
	case OT_163:
		if aAward.NumList[4] == 15 {
			winM = float64(betM) * odds
		}
	case OT_164:
		if aAward.NumList[4] == 16 {
			winM = float64(betM) * odds
		}
	case OT_165:
		if aAward.NumList[4] == 17 {
			winM = float64(betM) * odds
		}
	case OT_166:
		if aAward.NumList[4] == 18 {
			winM = float64(betM) * odds
		}
	case OT_167:
		if aAward.NumList[4] == 19 {
			winM = float64(betM) * odds
		}
	case OT_168:
		if aAward.NumList[4] == 20 {
			winM = float64(betM) * odds
		}
	case OT_169:
		if GetNumDS(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_170:
		if GetNumDS(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_171:
		if GetNumDX(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_172:
		if GetNumDX(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_173:
		if getHDS(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_174:
		if getHDS(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_175:
		if getWSDX(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_176:
		if getWSDX(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_177:
		if GetNumDF(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_178:
		if GetNumNF(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_179:
		if GetNumXF(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_180:
		if GetNumBF(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_181:
		if GetNumZ(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_182:
		if GetNumF(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_183:
		if GetNumB(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_184:
		if aAward.NumList[4] > aAward.NumList[3] {
			winM = float64(betM) * odds
		}
	case OT_185:
		if aAward.NumList[4] < aAward.NumList[3] {
			winM = float64(betM) * odds
		}

	case OT_186:
		if aAward.NumList[5] == 1 {
			winM = float64(betM) * odds
		}
	case OT_187:
		if aAward.NumList[5] == 2 {
			winM = float64(betM) * odds
		}
	case OT_188:
		if aAward.NumList[5] == 3 {
			winM = float64(betM) * odds
		}
	case OT_189:
		if aAward.NumList[5] == 4 {
			winM = float64(betM) * odds
		}
	case OT_190:
		if aAward.NumList[5] == 5 {
			winM = float64(betM) * odds
		}
	case OT_191:
		if aAward.NumList[5] == 6 {
			winM = float64(betM) * odds
		}
	case OT_192:
		if aAward.NumList[5] == 7 {
			winM = float64(betM) * odds
		}
	case OT_193:
		if aAward.NumList[5] == 8 {
			winM = float64(betM) * odds
		}
	case OT_194:
		if aAward.NumList[5] == 9 {
			winM = float64(betM) * odds
		}
	case OT_195:
		if aAward.NumList[5] == 10 {
			winM = float64(betM) * odds
		}
	case OT_196:
		if aAward.NumList[5] == 11 {
			winM = float64(betM) * odds
		}
	case OT_197:
		if aAward.NumList[5] == 12 {
			winM = float64(betM) * odds
		}
	case OT_198:
		if aAward.NumList[5] == 13 {
			winM = float64(betM) * odds
		}
	case OT_199:
		if aAward.NumList[5] == 14 {
			winM = float64(betM) * odds
		}
	case OT_200:
		if aAward.NumList[5] == 15 {
			winM = float64(betM) * odds
		}
	case OT_201:
		if aAward.NumList[5] == 16 {
			winM = float64(betM) * odds
		}
	case OT_202:
		if aAward.NumList[5] == 17 {
			winM = float64(betM) * odds
		}
	case OT_203:
		if aAward.NumList[5] == 18 {
			winM = float64(betM) * odds
		}
	case OT_204:
		if aAward.NumList[5] == 19 {
			winM = float64(betM) * odds
		}
	case OT_205:
		if aAward.NumList[5] == 20 {
			winM = float64(betM) * odds
		}
	case OT_206:
		if GetNumDS(aAward.NumList[5]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_207:
		if GetNumDS(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_208:
		if GetNumDX(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_209:
		if GetNumDX(aAward.NumList[5]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_210:
		if getHDS(aAward.NumList[5]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_211:
		if getHDS(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_212:
		if getWSDX(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_213:
		if getWSDX(aAward.NumList[5]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_214:
		if GetNumDF(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_215:
		if GetNumNF(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_216:
		if GetNumXF(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_217:
		if GetNumBF(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_218:
		if GetNumZ(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_219:
		if GetNumF(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_220:
		if GetNumB(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_221:
		if aAward.NumList[5] > aAward.NumList[2] {
			winM = float64(betM) * odds
		}
	case OT_222:
		if aAward.NumList[5] < aAward.NumList[2] {
			winM = float64(betM) * odds
		}

	case OT_223:
		if aAward.NumList[6] == 1 {
			winM = float64(betM) * odds
		}
	case OT_224:
		if aAward.NumList[6] == 2 {
			winM = float64(betM) * odds
		}
	case OT_225:
		if aAward.NumList[6] == 3 {
			winM = float64(betM) * odds
		}
	case OT_226:
		if aAward.NumList[6] == 4 {
			winM = float64(betM) * odds
		}
	case OT_227:
		if aAward.NumList[6] == 5 {
			winM = float64(betM) * odds
		}
	case OT_228:
		if aAward.NumList[6] == 6 {
			winM = float64(betM) * odds
		}
	case OT_229:
		if aAward.NumList[6] == 7 {
			winM = float64(betM) * odds
		}
	case OT_230:
		if aAward.NumList[6] == 8 {
			winM = float64(betM) * odds
		}
	case OT_231:
		if aAward.NumList[6] == 9 {
			winM = float64(betM) * odds
		}
	case OT_232:
		if aAward.NumList[6] == 10 {
			winM = float64(betM) * odds
		}
	case OT_233:
		if aAward.NumList[6] == 11 {
			winM = float64(betM) * odds
		}
	case OT_234:
		if aAward.NumList[6] == 12 {
			winM = float64(betM) * odds
		}
	case OT_235:
		if aAward.NumList[6] == 13 {
			winM = float64(betM) * odds
		}
	case OT_236:
		if aAward.NumList[6] == 14 {
			winM = float64(betM) * odds
		}
	case OT_237:
		if aAward.NumList[6] == 15 {
			winM = float64(betM) * odds
		}
	case OT_238:
		if aAward.NumList[6] == 16 {
			winM = float64(betM) * odds
		}
	case OT_239:
		if aAward.NumList[6] == 17 {
			winM = float64(betM) * odds
		}
	case OT_240:
		if aAward.NumList[6] == 18 {
			winM = float64(betM) * odds
		}
	case OT_241:
		if aAward.NumList[6] == 19 {
			winM = float64(betM) * odds
		}
	case OT_242:
		if aAward.NumList[6] == 20 {
			winM = float64(betM) * odds
		}
	case OT_243:
		if GetNumDS(aAward.NumList[6]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_244:
		if GetNumDS(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_245:
		if GetNumDX(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_246:
		if GetNumDX(aAward.NumList[6]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_247:
		if getHDS(aAward.NumList[6]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_248:
		if getHDS(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_249:
		if getWSDX(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_250:
		if getWSDX(aAward.NumList[6]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_251:
		if GetNumDF(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_252:
		if GetNumNF(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_253:
		if GetNumXF(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_254:
		if GetNumBF(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_255:
		if GetNumZ(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_256:
		if GetNumF(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_257:
		if GetNumB(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_258:
		if aAward.NumList[6] > aAward.NumList[1] {
			winM = float64(betM) * odds
		}
	case OT_259:
		if aAward.NumList[6] < aAward.NumList[1] {
			winM = float64(betM) * odds
		}

	case OT_260:
		if aAward.NumList[7] == 1 {
			winM = float64(betM) * odds
		}
	case OT_261:
		if aAward.NumList[7] == 2 {
			winM = float64(betM) * odds
		}
	case OT_262:
		if aAward.NumList[7] == 3 {
			winM = float64(betM) * odds
		}
	case OT_263:
		if aAward.NumList[7] == 4 {
			winM = float64(betM) * odds
		}
	case OT_264:
		if aAward.NumList[7] == 5 {
			winM = float64(betM) * odds
		}
	case OT_265:
		if aAward.NumList[7] == 6 {
			winM = float64(betM) * odds
		}
	case OT_266:
		if aAward.NumList[7] == 7 {
			winM = float64(betM) * odds
		}
	case OT_267:
		if aAward.NumList[7] == 8 {
			winM = float64(betM) * odds
		}
	case OT_268:
		if aAward.NumList[7] == 9 {
			winM = float64(betM) * odds
		}
	case OT_269:
		if aAward.NumList[7] == 10 {
			winM = float64(betM) * odds
		}
	case OT_270:
		if aAward.NumList[7] == 11 {
			winM = float64(betM) * odds
		}
	case OT_271:
		if aAward.NumList[7] == 12 {
			winM = float64(betM) * odds
		}
	case OT_272:
		if aAward.NumList[7] == 13 {
			winM = float64(betM) * odds
		}
	case OT_273:
		if aAward.NumList[7] == 14 {
			winM = float64(betM) * odds
		}
	case OT_274:
		if aAward.NumList[7] == 15 {
			winM = float64(betM) * odds
		}
	case OT_275:
		if aAward.NumList[7] == 16 {
			winM = float64(betM) * odds
		}
	case OT_276:
		if aAward.NumList[7] == 17 {
			winM = float64(betM) * odds
		}
	case OT_277:
		if aAward.NumList[7] == 18 {
			winM = float64(betM) * odds
		}
	case OT_278:
		if aAward.NumList[7] == 19 {
			winM = float64(betM) * odds
		}
	case OT_279:
		if aAward.NumList[7] == 20 {
			winM = float64(betM) * odds
		}
	case OT_280:
		if GetNumDS(aAward.NumList[7]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_281:
		if GetNumDS(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_282:
		if GetNumDX(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_283:
		if GetNumDX(aAward.NumList[7]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_284:
		if getHDS(aAward.NumList[7]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_285:
		if getHDS(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_286:
		if getWSDX(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_287:
		if getWSDX(aAward.NumList[7]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_288:
		if GetNumDF(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_289:
		if GetNumNF(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_290:
		if GetNumXF(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_291:
		if GetNumBF(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_292:
		if GetNumZ(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_293:
		if GetNumF(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_294:
		if GetNumB(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_295:
		if aAward.NumList[7] > aAward.NumList[0] {
			winM = float64(betM) * odds
		} else if aAward.NumList[7] == aAward.NumList[0] {
			winM = float64(betM)
		}
	case OT_296:
		if aAward.NumList[7] < aAward.NumList[0] {
			winM = float64(betM) * odds
		} else if aAward.NumList[7] == aAward.NumList[0] {
			winM = float64(betM)
		}
	case OT_317:
		if GetNumDS(aAward.ZH) == 1 {
			winM = float64(betM) * odds
		}
	case OT_318:
		if GetNumDS(aAward.ZH) == 2 {
			winM = float64(betM) * odds
		}
	case OT_319:
		if GetZHDX(aAward.ZH) == 2 {
			winM = float64(betM) * odds
		} else if GetZHDX(aAward.ZH) == 0 {
			winM = float64(betM)
		}
	case OT_320:
		if GetZHDX(aAward.ZH) == 1 {
			winM = float64(betM) * odds
		} else if GetZHDX(aAward.ZH) == 0 {
			winM = float64(betM)
		}
	case OT_321:
		if getWSDX(aAward.ZH) == 2 {
			winM = float64(betM) * odds
		}
	case OT_322:
		if getWSDX(aAward.ZH) == 1 {
			winM = float64(betM) * odds
		}

	default:
		return 0
	}

	return winM
}

func CheckNums(NumList []int, gameIndex int) error {
	if len(NumList) != 20 {
		return errors.New("号码个数不正确")
	}
	for _, n := range NumList {
		if n < 1 || n > 20 {
			return errors.New(fmt.Sprintf("号码[%d]不对。", n))
		}
	}

	return nil
}
