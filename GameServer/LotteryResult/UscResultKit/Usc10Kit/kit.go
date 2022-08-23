package Usc10Kit

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/kit/lotteryKit"
	"strconv"
)

type Award struct {
	NumList []int
	GYH     int //冠亚和
	GYDS    int //冠亚单双 1:单，2:双
	GYDX    int //冠亚大小 1:小，2:大
	StrNums string
	FT_FS   int // 番摊 番数
	FT_DS   int // 番摊 单双
}

func GetGZHDX(gyh int) int {
	GYDX := 1
	if gyh > 10 {
		GYDX = 2
	}
	return GYDX
}

func GetResultGZH(NumList []int) string {
	gzh := NumList[0] + NumList[1]
	t := GetGZHDX(gzh)
	dx := "小"
	if t == 2 {
		dx = "大"
	}
	ds := "单"
	if gzh%2 == 0 {
		ds = "双"
	}

	return fmt.Sprintf("%d,%s,%s", gzh, dx, ds)
}
func GetResultFtDS(ftFS int) string {
	ResultFtDS := "单"
	if ftFS%2 == 0 {
		ResultFtDS = "双"
	}
	return ResultFtDS
}

func GetFTHS(NumList []int) int {
	GYJH := NumList[0] + NumList[1] + NumList[2]
	return GYJH
}
func GetFS(FTHS int) int {
	FT_FS := FTHS % 4
	if FT_FS == 0 {
		FT_FS = 4
	}
	return FT_FS
}

/**
大小
*/
func GetNumDX(num int) int {
	if num > 5 {
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
func GetAward(strNums string) (Award, error) {
	var NumList = lotteryKit.GetStrNum2Arr(strNums)
	GYH := NumList[0] + NumList[1]
	GYDS := 1
	if GYH%2 == 0 {
		GYDS = 2
	}
	GYDX := GetGZHDX(GYH)
	StrNums := ""
	for i := 0; i < len(NumList); i++ {
		StrNums = StrNums + "," + strconv.Itoa(NumList[i])
	}
	GYJH := GetFTHS(NumList)
	FT_FS := GetFS(GYJH)

	FT_DS := 1
	if FT_FS%2 == 0 {
		FT_DS = 2
	}

	aAward := Award{
		NumList: NumList,
		GYH:     GYH,
		GYDS:    GYDS,
		GYDX:    GYDX,
		StrNums: StrNums,
		FT_FS:   FT_FS,
		FT_DS:   FT_DS,
	}
	return aAward, nil
}

func GetWinMoney(aAward Award, oddsType, betM int, odds float64) float64 {
	winM := float64(0)
	switch oddsType {
	case OT_001:
		if aAward.GYH == 3 {
			winM = float64(betM) * odds
		}

	case OT_002:
		if aAward.GYH == 4 {
			winM = float64(betM) * odds
		}

	case OT_003:
		if aAward.GYH == 5 {
			winM = float64(betM) * odds
		}

	case OT_004:
		if aAward.GYH == 6 {
			winM = float64(betM) * odds
		}

	case OT_005:
		if aAward.GYH == 7 {
			winM = float64(betM) * odds
		}

	case OT_006:
		if aAward.GYH == 8 {
			winM = float64(betM) * odds
		}

	case OT_007:
		if aAward.GYH == 9 {
			winM = float64(betM) * odds
		}

	case OT_008:
		if aAward.GYH == 10 {
			winM = float64(betM) * odds
		}

	case OT_009:
		if aAward.GYH == 11 {
			winM = float64(betM) * odds
		}

	case OT_010:
		if aAward.GYH == 12 {
			winM = float64(betM) * odds
		}

	case OT_011:
		if aAward.GYH == 13 {
			winM = float64(betM) * odds
		}

	case OT_012:
		if aAward.GYH == 14 {
			winM = float64(betM) * odds
		}

	case OT_013:
		if aAward.GYH == 15 {
			winM = float64(betM) * odds
		}

	case OT_014:
		if aAward.GYH == 16 {
			winM = float64(betM) * odds
		}

	case OT_015:
		if aAward.GYH == 17 {
			winM = float64(betM) * odds
		}

	case OT_016:
		if aAward.GYH == 18 {
			winM = float64(betM) * odds
		}

	case OT_017:
		if aAward.GYH == 19 {
			winM = float64(betM) * odds
		}

	case OT_018:
		if aAward.GYDS == 1 {
			winM = float64(betM) * odds
		}

	case OT_019:
		if aAward.GYDS == 2 {
			winM = float64(betM) * odds
		}

	case OT_020:
		if aAward.GYDX == 2 {
			winM = float64(betM) * odds
		}

	case OT_021:
		if aAward.GYDX == 1 {
			winM = float64(betM) * odds
		}

	//冠军
	case OT_022:
		if aAward.NumList[0] == 1 {
			winM = float64(betM) * odds
		}
	case OT_023:
		if aAward.NumList[0] == 2 {
			winM = float64(betM) * odds
		}
	case OT_024:
		if aAward.NumList[0] == 3 {
			winM = float64(betM) * odds
		}
	case OT_025:
		if aAward.NumList[0] == 4 {
			winM = float64(betM) * odds
		}
	case OT_026:
		if aAward.NumList[0] == 5 {
			winM = float64(betM) * odds
		}
	case OT_027:
		if aAward.NumList[0] == 6 {
			winM = float64(betM) * odds
		}
	case OT_028:
		if aAward.NumList[0] == 7 {
			winM = float64(betM) * odds
		}
	case OT_029:
		if aAward.NumList[0] == 8 {
			winM = float64(betM) * odds
		}
	case OT_030:
		if aAward.NumList[0] == 9 {
			winM = float64(betM) * odds
		}
	case OT_031:
		if aAward.NumList[0] == 10 {
			winM = float64(betM) * odds
		}
	case OT_032:
		if aAward.NumList[0] > aAward.NumList[9] {
			winM = float64(betM) * odds
		}
	case OT_033:
		if aAward.NumList[0] < aAward.NumList[9] {
			winM = float64(betM) * odds
		}
	case OT_034:
		if GetNumDX(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_035:
		if GetNumDX(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_036:
		if GetNumDS(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_037:
		if GetNumDS(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	//亚军
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
		if aAward.NumList[1] > aAward.NumList[8] {
			winM = float64(betM) * odds
		}
	case OT_049:
		if aAward.NumList[1] < aAward.NumList[8] {
			winM = float64(betM) * odds
		}
	case OT_050:
		if GetNumDX(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_051:
		if GetNumDX(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_052:
		if GetNumDS(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_053:
		if GetNumDS(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}

	//第三名
	case OT_054:
		if aAward.NumList[2] == 1 {
			winM = float64(betM) * odds
		}
	case OT_055:
		if aAward.NumList[2] == 2 {
			winM = float64(betM) * odds
		}
	case OT_056:
		if aAward.NumList[2] == 3 {
			winM = float64(betM) * odds
		}
	case OT_057:
		if aAward.NumList[2] == 4 {
			winM = float64(betM) * odds
		}
	case OT_058:
		if aAward.NumList[2] == 5 {
			winM = float64(betM) * odds
		}
	case OT_059:
		if aAward.NumList[2] == 6 {
			winM = float64(betM) * odds
		}
	case OT_060:
		if aAward.NumList[2] == 7 {
			winM = float64(betM) * odds
		}
	case OT_061:
		if aAward.NumList[2] == 8 {
			winM = float64(betM) * odds
		}
	case OT_062:
		if aAward.NumList[2] == 9 {
			winM = float64(betM) * odds
		}
	case OT_063:
		if aAward.NumList[2] == 10 {
			winM = float64(betM) * odds
		}
	case OT_064:
		if aAward.NumList[2] > aAward.NumList[7] {
			winM = float64(betM) * odds
		}
	case OT_065:
		if aAward.NumList[2] < aAward.NumList[7] {
			winM = float64(betM) * odds
		}
	case OT_066:
		if GetNumDX(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_067:
		if GetNumDX(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_068:
		if GetNumDS(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_069:
		if GetNumDS(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}

	//第四名
	case OT_070:
		if aAward.NumList[3] == 1 {
			winM = float64(betM) * odds
		}
	case OT_071:
		if aAward.NumList[3] == 2 {
			winM = float64(betM) * odds
		}
	case OT_072:
		if aAward.NumList[3] == 3 {
			winM = float64(betM) * odds
		}
	case OT_073:
		if aAward.NumList[3] == 4 {
			winM = float64(betM) * odds
		}
	case OT_074:
		if aAward.NumList[3] == 5 {
			winM = float64(betM) * odds
		}
	case OT_075:
		if aAward.NumList[3] == 6 {
			winM = float64(betM) * odds
		}
	case OT_076:
		if aAward.NumList[3] == 7 {
			winM = float64(betM) * odds
		}
	case OT_077:
		if aAward.NumList[3] == 8 {
			winM = float64(betM) * odds
		}
	case OT_078:
		if aAward.NumList[3] == 9 {
			winM = float64(betM) * odds
		}
	case OT_079:
		if aAward.NumList[3] == 10 {
			winM = float64(betM) * odds
		}
	case OT_080:
		if aAward.NumList[3] > aAward.NumList[6] {
			winM = float64(betM) * odds
		}
	case OT_081:
		if aAward.NumList[3] < aAward.NumList[6] {
			winM = float64(betM) * odds
		}
	case OT_082:
		if GetNumDX(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_083:
		if GetNumDX(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_084:
		if GetNumDS(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_085:
		if GetNumDS(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}

	//第五名
	case OT_086:
		if aAward.NumList[4] == 1 {
			winM = float64(betM) * odds
		}
	case OT_087:
		if aAward.NumList[4] == 2 {
			winM = float64(betM) * odds
		}
	case OT_088:
		if aAward.NumList[4] == 3 {
			winM = float64(betM) * odds
		}
	case OT_089:
		if aAward.NumList[4] == 4 {
			winM = float64(betM) * odds
		}
	case OT_090:
		if aAward.NumList[4] == 5 {
			winM = float64(betM) * odds
		}
	case OT_091:
		if aAward.NumList[4] == 6 {
			winM = float64(betM) * odds
		}
	case OT_092:
		if aAward.NumList[4] == 7 {
			winM = float64(betM) * odds
		}
	case OT_093:
		if aAward.NumList[4] == 8 {
			winM = float64(betM) * odds
		}
	case OT_094:
		if aAward.NumList[4] == 9 {
			winM = float64(betM) * odds
		}
	case OT_095:
		if aAward.NumList[4] == 10 {
			winM = float64(betM) * odds
		}
	case OT_096:
		if aAward.NumList[4] > aAward.NumList[5] {
			winM = float64(betM) * odds
		}
	case OT_097:
		if aAward.NumList[4] < aAward.NumList[5] {
			winM = float64(betM) * odds
		}
	case OT_098:
		if GetNumDX(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_099:
		if GetNumDX(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_100:
		if GetNumDS(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_101:
		if GetNumDS(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}

	//第六名
	case OT_102:
		if aAward.NumList[5] == 1 {
			winM = float64(betM) * odds
		}
	case OT_103:
		if aAward.NumList[5] == 2 {
			winM = float64(betM) * odds
		}
	case OT_104:
		if aAward.NumList[5] == 3 {
			winM = float64(betM) * odds
		}
	case OT_105:
		if aAward.NumList[5] == 4 {
			winM = float64(betM) * odds
		}
	case OT_106:
		if aAward.NumList[5] == 5 {
			winM = float64(betM) * odds
		}
	case OT_107:
		if aAward.NumList[5] == 6 {
			winM = float64(betM) * odds
		}
	case OT_108:
		if aAward.NumList[5] == 7 {
			winM = float64(betM) * odds
		}
	case OT_109:
		if aAward.NumList[5] == 8 {
			winM = float64(betM) * odds
		}
	case OT_110:
		if aAward.NumList[5] == 9 {
			winM = float64(betM) * odds
		}
	case OT_111:
		if aAward.NumList[5] == 10 {
			winM = float64(betM) * odds
		}
	case OT_112:
		if GetNumDX(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_113:
		if GetNumDX(aAward.NumList[5]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_114:
		if GetNumDS(aAward.NumList[5]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_115:
		if GetNumDS(aAward.NumList[5]) == 2 {
			winM = float64(betM) * odds
		}

	//第七名
	case OT_118:
		if aAward.NumList[6] == 1 {
			winM = float64(betM) * odds
		}
	case OT_119:
		if aAward.NumList[6] == 2 {
			winM = float64(betM) * odds
		}
	case OT_120:
		if aAward.NumList[6] == 3 {
			winM = float64(betM) * odds
		}
	case OT_121:
		if aAward.NumList[6] == 4 {
			winM = float64(betM) * odds
		}
	case OT_122:
		if aAward.NumList[6] == 5 {
			winM = float64(betM) * odds
		}
	case OT_123:
		if aAward.NumList[6] == 6 {
			winM = float64(betM) * odds
		}
	case OT_124:
		if aAward.NumList[6] == 7 {
			winM = float64(betM) * odds
		}
	case OT_125:
		if aAward.NumList[6] == 8 {
			winM = float64(betM) * odds
		}
	case OT_126:
		if aAward.NumList[6] == 9 {
			winM = float64(betM) * odds
		}
	case OT_127:
		if aAward.NumList[6] == 10 {
			winM = float64(betM) * odds
		}
	case OT_128:
		if GetNumDX(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_129:
		if GetNumDX(aAward.NumList[6]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_130:
		if GetNumDS(aAward.NumList[6]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_131:
		if GetNumDS(aAward.NumList[6]) == 2 {
			winM = float64(betM) * odds
		}

	//第八名
	case OT_134:
		if aAward.NumList[7] == 1 {
			winM = float64(betM) * odds
		}
	case OT_135:
		if aAward.NumList[7] == 2 {
			winM = float64(betM) * odds
		}
	case OT_136:
		if aAward.NumList[7] == 3 {
			winM = float64(betM) * odds
		}
	case OT_137:
		if aAward.NumList[7] == 4 {
			winM = float64(betM) * odds
		}
	case OT_138:
		if aAward.NumList[7] == 5 {
			winM = float64(betM) * odds
		}
	case OT_139:
		if aAward.NumList[7] == 6 {
			winM = float64(betM) * odds
		}
	case OT_140:
		if aAward.NumList[7] == 7 {
			winM = float64(betM) * odds
		}
	case OT_141:
		if aAward.NumList[7] == 8 {
			winM = float64(betM) * odds
		}
	case OT_142:
		if aAward.NumList[7] == 9 {
			winM = float64(betM) * odds
		}
	case OT_143:
		if aAward.NumList[7] == 10 {
			winM = float64(betM) * odds
		}
	case OT_144:
		if GetNumDX(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_145:
		if GetNumDX(aAward.NumList[7]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_146:
		if GetNumDS(aAward.NumList[7]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_147:
		if GetNumDS(aAward.NumList[7]) == 2 {
			winM = float64(betM) * odds
		}

	//第九名
	case OT_150:
		if aAward.NumList[8] == 1 {
			winM = float64(betM) * odds
		}
	case OT_151:
		if aAward.NumList[8] == 2 {
			winM = float64(betM) * odds
		}
	case OT_152:
		if aAward.NumList[8] == 3 {
			winM = float64(betM) * odds
		}
	case OT_153:
		if aAward.NumList[8] == 4 {
			winM = float64(betM) * odds
		}
	case OT_154:
		if aAward.NumList[8] == 5 {
			winM = float64(betM) * odds
		}
	case OT_155:
		if aAward.NumList[8] == 6 {
			winM = float64(betM) * odds
		}
	case OT_156:
		if aAward.NumList[8] == 7 {
			winM = float64(betM) * odds
		}
	case OT_157:
		if aAward.NumList[8] == 8 {
			winM = float64(betM) * odds
		}
	case OT_158:
		if aAward.NumList[8] == 9 {
			winM = float64(betM) * odds
		}
	case OT_159:
		if aAward.NumList[8] == 10 {
			winM = float64(betM) * odds
		}
	case OT_160:
		if GetNumDX(aAward.NumList[8]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_161:
		if GetNumDX(aAward.NumList[8]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_162:
		if GetNumDS(aAward.NumList[8]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_163:
		if GetNumDS(aAward.NumList[8]) == 2 {
			winM = float64(betM) * odds
		}

	//第十名
	case OT_166:
		if aAward.NumList[9] == 1 {
			winM = float64(betM) * odds
		}
	case OT_167:
		if aAward.NumList[9] == 2 {
			winM = float64(betM) * odds
		}
	case OT_168:
		if aAward.NumList[9] == 3 {
			winM = float64(betM) * odds
		}
	case OT_169:
		if aAward.NumList[9] == 4 {
			winM = float64(betM) * odds
		}
	case OT_170:
		if aAward.NumList[9] == 5 {
			winM = float64(betM) * odds
		}
	case OT_171:
		if aAward.NumList[9] == 6 {
			winM = float64(betM) * odds
		}
	case OT_172:
		if aAward.NumList[9] == 7 {
			winM = float64(betM) * odds
		}
	case OT_173:
		if aAward.NumList[9] == 8 {
			winM = float64(betM) * odds
		}
	case OT_174:
		if aAward.NumList[9] == 9 {
			winM = float64(betM) * odds
		}
	case OT_175:
		if aAward.NumList[9] == 10 {
			winM = float64(betM) * odds
		}
	case OT_176:
		if GetNumDX(aAward.NumList[9]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_177:
		if GetNumDX(aAward.NumList[9]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_178:
		if GetNumDS(aAward.NumList[9]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_179:
		if GetNumDS(aAward.NumList[9]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_182:
		//正1
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 3 {
			winM = float64(betM)
		}
	case OT_183:
		//正2
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 4 {
			winM = float64(betM)
		}
	case OT_184:
		//正3
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 1 {
			winM = float64(betM)
		}
	case OT_185:
		//正4
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 2 {
			winM = float64(betM)
		}
	case OT_186:
		//番1
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		}
	case OT_187:
		//番2
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		}
	case OT_188:
		//番3
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		}
	case OT_189:
		//番4
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		}

	case OT_190:
		//1-2角
		if aAward.FT_FS == 1 || aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		}
	case OT_191:
		//1-4角
		if aAward.FT_FS == 1 || aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		}
	case OT_192:
		//192 //	番摊 2-3角
		if aAward.FT_FS == 2 || aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		}
	case OT_193:
		//193 //	番摊 3-4角
		if aAward.FT_FS == 3 || aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		}

	case OT_194:
		//194 //	番摊 1念2
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 2 {
			winM = float64(betM)
		}
	case OT_195:
		//if aAward.FT_N == FT_N1_3 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 3 {
			winM = float64(betM)
		}
	case OT_196:
		//if aAward.FT_N == FT_N1_4 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 4 {
			winM = float64(betM)
		}
	case OT_197:
		//if aAward.FT_N == FT_N2_1 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 1 {
			winM = float64(betM)
		}
	case OT_198:
		//if aAward.FT_N == FT_N2_3 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 3 {
			winM = float64(betM)
		}
	case OT_199:
		//if aAward.FT_N == FT_N2_4 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 4 {
			winM = float64(betM)
		}
	case OT_200:
		//if aAward.FT_N == FT_N3_1 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 1 {
			winM = float64(betM)
		}
	case OT_201:
		//if aAward.FT_N == FT_N3_2 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 2 {
			winM = float64(betM)
		}
	case OT_202:
		//if aAward.FT_N == FT_N3_4 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 4 {
			winM = float64(betM)
		}
	case OT_203:
		//if aAward.FT_N == FT_N4_1 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 1 {
			winM = float64(betM)
		}
	case OT_204:
		//if aAward.FT_N == FT_N4_2 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 2 {
			winM = float64(betM)
		}
	case OT_205:
		//if aAward.FT_N == FT_N4_3 {
		//	winM = float64(betM) * odds
		//}
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 3 {
			winM = float64(betM)
		}
	case OT_206:
		if aAward.FT_DS == FT_DS_D {
			winM = float64(betM) * odds
		}
	case OT_207:
		if aAward.FT_DS == FT_DS_S {
			winM = float64(betM) * odds
		}
	default:
		return 0
	}

	return winM
}

func CheckNums(NumList []int, gameIndex int) error {
	if len(NumList) != 10 {
		return errors.New("号码个数不正确")
	}
	mp := make(map[int]int)
	for _, n := range NumList {
		if n < 1 || n > 10 {
			return errors.New(fmt.Sprintf("号码[%d]不对。", n))
		}
		mp[n] = mp[n] + 1
	}
	for i := 1; i <= 10; i++ {
		if mp[i] > 1 {
			return errors.New(fmt.Sprintf("号码[%d]重复。", i))
		}
		if mp[i] == 0 {
			return errors.New(fmt.Sprintf("号码[%d]没有。", i))
		}
	}

	return nil
}
