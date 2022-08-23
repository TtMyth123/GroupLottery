package Usc5Kit

import (
	"errors"
	"fmt"
	"github.com/TtMyth123/kit/lotteryKit"
	"strconv"
)

type Award struct {
	NumList []int
	ZH      int //总和
	Q3BZ    int //豹子
	Q3SZ    int //顺子
	Q3DZ    int //对子
	Q3BS    int //半顺
	Q3Z6    int //杂六
	Z3BZ    int //豹子
	Z3SZ    int //顺子
	Z3DZ    int //对子
	Z3BS    int //半顺
	Z3Z6    int //杂六
	H3BZ    int //豹子
	H3SZ    int //顺子
	H3DZ    int //对子
	H3BS    int //半顺
	H3Z6    int //杂六

	StrNums string
	FT_FS   int // 番摊 番数
	FT_DS   int // 番摊 单双
}

func IsBZ(nums [3]int) int {
	if nums[0] == nums[1] && nums[1] == nums[2] {
		return 1
	}
	return 0
}

func IsSZ(nums [3]int) int {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3-1-i; j++ {
			if nums[j] > nums[j+1] {
				t := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = t
			}
		}
	}

	if nums[0] == 0 && nums[1] == 8 && nums[2] == 9 {
		return 1
	} else if nums[0] == 0 && nums[1] == 1 && nums[2] == 9 {
		return 1
	} else if nums[1]-nums[0] == 1 && nums[2]-nums[1] == 1 {
		return 1
	} else {
		return 0
	}
}
func IsDZ(isBZ int, nums [3]int) int {
	if isBZ == 1 {
		return 0
	}
	if nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2] {
		return 1
	} else {
		return 0
	}
}

func IsBS(isSZ, isDZ int, nums [3]int) int {
	return 0
}

func IsZ6(isBZ, isDZ, isSZ, isBS int) int {
	if isBZ == 1 || isDZ == 1 || isSZ == 1 || isBS == 1 {
		return 0
	}
	return 1
}

func GetZH(NumList []int) int {
	ZH := 0
	for i := 0; i < len(NumList); i++ {
		ZH += NumList[i]
	}

	return ZH
}
func GetZHDS(zh int) int {
	if zh%2 == 0 {
		return 2
	} else {
		return 1
	}
}

func GetZHDX(zh int) int {
	if zh >= 23 {
		return 2
	} else {
		return 1
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
func GetFTHS(NumList []int) int {
	GYJH := NumList[0] + NumList[1] + NumList[2] + NumList[3] + NumList[4]
	return GYJH
}
func GetFS(FTHS int) int {
	FT_FS := FTHS % 4
	if FT_FS == 0 {
		FT_FS = 4
	}
	return FT_FS
}
func GetResultFtDS(ResultFtFS int) string {
	ResultFtDS := "单"
	if ResultFtFS%2 == 0 {
		ResultFtDS = "双"
	}
	return ResultFtDS
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
	ZH := GetZH(NumList)

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
	Q3BZ := IsBZ([3]int{NumList[0], NumList[1], NumList[2]})
	Q3SZ := IsSZ([3]int{NumList[0], NumList[1], NumList[2]})
	Q3DZ := IsDZ(Q3BZ, [3]int{NumList[0], NumList[1], NumList[2]})
	Q3BS := IsBS(Q3SZ, Q3DZ, [3]int{NumList[0], NumList[1], NumList[2]})
	Q3Z6 := IsZ6(Q3BZ, Q3SZ, Q3DZ, Q3BS)

	Z3BZ := IsBZ([3]int{NumList[1], NumList[2], NumList[3]})
	Z3SZ := IsSZ([3]int{NumList[1], NumList[2], NumList[3]})
	Z3DZ := IsDZ(Z3SZ, [3]int{NumList[1], NumList[2], NumList[3]})
	Z3BS := IsBS(Z3SZ, Z3DZ, [3]int{NumList[1], NumList[2], NumList[3]})
	Z3Z6 := IsZ6(Z3BZ, Z3SZ, Z3DZ, Z3BS)

	H3BZ := IsBZ([3]int{NumList[2], NumList[3], NumList[4]})
	H3SZ := IsSZ([3]int{NumList[2], NumList[3], NumList[4]})
	H3DZ := IsDZ(H3SZ, [3]int{NumList[2], NumList[3], NumList[4]})
	H3BS := IsBS(H3SZ, H3DZ, [3]int{NumList[2], NumList[3], NumList[4]})
	H3Z6 := IsZ6(H3BZ, H3SZ, H3DZ, H3BS)

	aAward := Award{
		NumList: NumList,
		ZH:      ZH,
		Q3BZ:    Q3BZ,
		Q3SZ:    Q3SZ,
		Q3DZ:    Q3DZ,
		Q3BS:    Q3BS,
		Q3Z6:    Q3Z6,
		Z3BZ:    Z3BZ,
		Z3SZ:    Z3SZ,
		Z3DZ:    Z3DZ,
		Z3BS:    Z3BS,
		Z3Z6:    Z3Z6,
		H3BZ:    H3BZ,
		H3SZ:    H3SZ,
		H3DZ:    H3DZ,
		H3BS:    H3BS,
		H3Z6:    H3Z6,
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
		if aAward.NumList[0] == 0 {
			winM = float64(betM) * odds
		}
	case OT_002:
		if aAward.NumList[0] == 1 {
			winM = float64(betM) * odds
		}
	case OT_003:
		if aAward.NumList[0] == 2 {
			winM = float64(betM) * odds
		}
	case OT_004:
		if aAward.NumList[0] == 3 {
			winM = float64(betM) * odds
		}
	case OT_005:
		if aAward.NumList[0] == 4 {
			winM = float64(betM) * odds
		}
	case OT_006:
		if aAward.NumList[0] == 5 {
			winM = float64(betM) * odds
		}
	case OT_007:
		if aAward.NumList[0] == 6 {
			winM = float64(betM) * odds
		}
	case OT_008:
		if aAward.NumList[0] == 7 {
			winM = float64(betM) * odds
		}
	case OT_009:
		if aAward.NumList[0] == 8 {
			winM = float64(betM) * odds
		}
	case OT_010:
		if aAward.NumList[0] == 9 {
			winM = float64(betM) * odds
		}
	case OT_011:
		if GetNumDS(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_012:
		if GetNumDS(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_013:
		if GetNumDX(aAward.NumList[0]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_014:
		if GetNumDX(aAward.NumList[0]) == 1 {
			winM = float64(betM) * odds
		}

	case OT_015:
		if aAward.NumList[1] == 0 {
			winM = float64(betM) * odds
		}
	case OT_016:
		if aAward.NumList[1] == 1 {
			winM = float64(betM) * odds
		}
	case OT_017:
		if aAward.NumList[1] == 2 {
			winM = float64(betM) * odds
		}
	case OT_018:
		if aAward.NumList[1] == 3 {
			winM = float64(betM) * odds
		}
	case OT_019:
		if aAward.NumList[1] == 4 {
			winM = float64(betM) * odds
		}
	case OT_020:
		if aAward.NumList[1] == 5 {
			winM = float64(betM) * odds
		}
	case OT_021:
		if aAward.NumList[1] == 6 {
			winM = float64(betM) * odds
		}
	case OT_022:
		if aAward.NumList[1] == 7 {
			winM = float64(betM) * odds
		}
	case OT_023:
		if aAward.NumList[1] == 8 {
			winM = float64(betM) * odds
		}
	case OT_024:
		if aAward.NumList[1] == 9 {
			winM = float64(betM) * odds
		}
	case OT_025:
		if GetNumDS(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_026:
		if GetNumDS(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_027:
		if GetNumDX(aAward.NumList[1]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_028:
		if GetNumDX(aAward.NumList[1]) == 1 {
			winM = float64(betM) * odds
		}

	case OT_029:
		if aAward.NumList[2] == 0 {
			winM = float64(betM) * odds
		}
	case OT_030:
		if aAward.NumList[2] == 1 {
			winM = float64(betM) * odds
		}
	case OT_031:
		if aAward.NumList[2] == 2 {
			winM = float64(betM) * odds
		}
	case OT_032:
		if aAward.NumList[2] == 3 {
			winM = float64(betM) * odds
		}
	case OT_033:
		if aAward.NumList[2] == 4 {
			winM = float64(betM) * odds
		}
	case OT_034:
		if aAward.NumList[2] == 5 {
			winM = float64(betM) * odds
		}
	case OT_035:
		if aAward.NumList[2] == 6 {
			winM = float64(betM) * odds
		}
	case OT_036:
		if aAward.NumList[2] == 7 {
			winM = float64(betM) * odds
		}
	case OT_037:
		if aAward.NumList[2] == 8 {
			winM = float64(betM) * odds
		}
	case OT_038:
		if aAward.NumList[2] == 9 {
			winM = float64(betM) * odds
		}
	case OT_039:
		if GetNumDS(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_040:
		if GetNumDS(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_041:
		if GetNumDX(aAward.NumList[2]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_042:
		if GetNumDX(aAward.NumList[2]) == 1 {
			winM = float64(betM) * odds
		}

	case OT_043:
		if aAward.NumList[3] == 0 {
			winM = float64(betM) * odds
		}
	case OT_044:
		if aAward.NumList[3] == 1 {
			winM = float64(betM) * odds
		}
	case OT_045:
		if aAward.NumList[3] == 2 {
			winM = float64(betM) * odds
		}
	case OT_046:
		if aAward.NumList[3] == 3 {
			winM = float64(betM) * odds
		}
	case OT_047:
		if aAward.NumList[3] == 4 {
			winM = float64(betM) * odds
		}
	case OT_048:
		if aAward.NumList[3] == 5 {
			winM = float64(betM) * odds
		}
	case OT_049:
		if aAward.NumList[3] == 6 {
			winM = float64(betM) * odds
		}
	case OT_050:
		if aAward.NumList[3] == 7 {
			winM = float64(betM) * odds
		}
	case OT_051:
		if aAward.NumList[3] == 8 {
			winM = float64(betM) * odds
		}
	case OT_052:
		if aAward.NumList[3] == 9 {
			winM = float64(betM) * odds
		}
	case OT_053:
		if GetNumDS(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_054:
		if GetNumDS(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_055:
		if GetNumDX(aAward.NumList[3]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_056:
		if GetNumDX(aAward.NumList[3]) == 1 {
			winM = float64(betM) * odds
		}

	case OT_057:
		if aAward.NumList[4] == 0 {
			winM = float64(betM) * odds
		}
	case OT_058:
		if aAward.NumList[4] == 1 {
			winM = float64(betM) * odds
		}
	case OT_059:
		if aAward.NumList[4] == 2 {
			winM = float64(betM) * odds
		}
	case OT_060:
		if aAward.NumList[4] == 3 {
			winM = float64(betM) * odds
		}
	case OT_061:
		if aAward.NumList[4] == 4 {
			winM = float64(betM) * odds
		}
	case OT_062:
		if aAward.NumList[4] == 5 {
			winM = float64(betM) * odds
		}
	case OT_063:
		if aAward.NumList[4] == 6 {
			winM = float64(betM) * odds
		}
	case OT_064:
		if aAward.NumList[4] == 7 {
			winM = float64(betM) * odds
		}
	case OT_065:
		if aAward.NumList[4] == 8 {
			winM = float64(betM) * odds
		}
	case OT_066:
		if aAward.NumList[4] == 9 {
			winM = float64(betM) * odds
		}
	case OT_067:
		if GetNumDS(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}
	case OT_068:
		if GetNumDS(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_069:
		if GetNumDX(aAward.NumList[4]) == 2 {
			winM = float64(betM) * odds
		}
	case OT_070:
		if GetNumDX(aAward.NumList[4]) == 1 {
			winM = float64(betM) * odds
		}

	case OT_071:
		if GetZHDS(aAward.ZH) == 1 {
			winM = float64(betM) * odds
		}
	case OT_072:
		if GetZHDS(aAward.ZH) == 2 {
			winM = float64(betM) * odds
		}
	case OT_073:
		if GetZHDX(aAward.ZH) == 2 {
			winM = float64(betM) * odds
		}
	case OT_074:
		if GetZHDX(aAward.ZH) == 1 {
			winM = float64(betM) * odds
		}

	case OT_075:
		if aAward.NumList[0] > aAward.NumList[4] {
			winM = float64(betM) * odds
		} else if aAward.NumList[0] == aAward.NumList[4] {
			winM = float64(betM)
		}
	case OT_076:
		if aAward.NumList[0] < aAward.NumList[4] {
			winM = float64(betM) * odds
		} else if aAward.NumList[0] == aAward.NumList[4] {
			winM = float64(betM)
		}
	case OT_077:
		if aAward.NumList[0] == aAward.NumList[4] {
			winM = float64(betM) * odds
		}

	case OT_078:
		if aAward.Q3BZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_079:
		if aAward.Q3SZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_080:
		if aAward.Q3DZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_081:
		if aAward.Q3BS == 1 {
			winM = float64(betM) * odds
		}
	case OT_082:
		if aAward.Q3Z6 == 1 {
			winM = float64(betM) * odds
		}

	case OT_083:
		if aAward.Z3BZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_084:
		if aAward.Z3SZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_085:
		if aAward.Z3DZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_086:
		if aAward.Z3BS == 1 {
			winM = float64(betM) * odds
		}
	case OT_087:
		if aAward.Z3Z6 == 1 {
			winM = float64(betM) * odds
		}

	case OT_088:
		if aAward.H3BZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_089:
		if aAward.H3SZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_090:
		if aAward.H3DZ == 1 {
			winM = float64(betM) * odds
		}
	case OT_091:
		if aAward.H3BS == 1 {
			winM = float64(betM) * odds
		}
	case OT_092:
		if aAward.H3Z6 == 1 {
			winM = float64(betM) * odds
		}

	case OT_108:
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 3 {
			winM = float64(betM)
		}
	case OT_109:
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 4 {
			winM = float64(betM)
		}
	case OT_110:
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 1 {
			winM = float64(betM)
		}
	case OT_111:
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS != 2 {
			winM = float64(betM)
		}

	case OT_112:
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		}
	case OT_113:
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		}
	case OT_114:
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		}
	case OT_115:
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		}

	case OT_116:
		if aAward.FT_FS == 1 || aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		}
	case OT_117:
		if aAward.FT_FS == 1 || aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		}
	case OT_118:
		if aAward.FT_FS == 2 || aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		}
	case OT_119:
		if aAward.FT_FS == 3 || aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		}

	case OT_120:
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 2 {
			winM = float64(betM)
		}
	case OT_121:
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 3 {
			winM = float64(betM)
		}
	case OT_122:
		if aAward.FT_FS == 1 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 4 {
			winM = float64(betM)
		}
	case OT_123:
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 1 {
			winM = float64(betM)
		}
	case OT_124:
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 3 {
			winM = float64(betM)
		}
	case OT_125:
		if aAward.FT_FS == 2 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 4 {
			winM = float64(betM)
		}
	case OT_126:
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 1 {
			winM = float64(betM)
		}
	case OT_127:
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 2 {
			winM = float64(betM)
		}
	case OT_128:
		if aAward.FT_FS == 3 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 4 {
			winM = float64(betM)
		}
	case OT_129:
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 1 {
			winM = float64(betM)
		}
	case OT_130:
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 2 {
			winM = float64(betM)
		}
	case OT_131:
		if aAward.FT_FS == 4 {
			winM = float64(betM) * odds
		} else if aAward.FT_FS == 3 {
			winM = float64(betM)
		}

	case OT_132:
		if aAward.FT_DS == 1 {
			winM = float64(betM) * odds
		}
	case OT_133:
		if aAward.FT_DS == 2 {
			winM = float64(betM) * odds
		}
	}

	return winM
}

func CheckNums(NumList []int, gameIndex int) error {
	if len(NumList) != 5 {
		return errors.New("号码个数不正确")
	}
	for _, n := range NumList {
		if n < 0 || n > 9 {
			return errors.New(fmt.Sprintf("号码[%d]不对。", n))
		}
	}

	return nil
}
