package httpGameServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/OtherServer/httpGameServer/httpBox"
	"github.com/TtMyth123/UserInfoRpc/models/mconst"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"sync"

	"github.com/TtMyth123/UserInfoRpc/models"
)

var (
	mHttp *HttpGameServer

	mapGameType2Region  map[int]map[int]models.TtArea
	gameType2RegionLock sync.RWMutex
)

const Region = `瞳瞳`

type HttpGameServer struct {
	baseArrZg     string
	baseArr       string
	baseArrResult string
	mHttpClient   *httpClientKit.HttpClient
}

func Init() {
	mHttp = new(HttpGameServer)
	//http://47.244.125.83:7799
	mHttp.baseArrZg = beego.AppConfig.String("HttpGameServerUrlZg")
	mHttp.baseArr = beego.AppConfig.String("HttpGameServerUrl")
	mHttp.baseArrResult = beego.AppConfig.String("ResultHttpUrl")
	mHttp.mHttpClient = httpClientKit.GetHttpClient("")
	readLoadGameType2Region()

}
func readLoadGameType2Region() {
	gameType2RegionLock.Lock()
	defer gameType2RegionLock.Unlock()

	mapGameType2Region = make(map[int]map[int]models.TtArea)
	o := orm.NewOrm()
	arrTtArea := make([]models.TtArea, 0)
	mpTtArea := make(map[int]models.TtArea)
	o.QueryTable(mconst.TableName_TtArea).All(&arrTtArea)
	for _, v := range arrTtArea {
		mpTtArea[v.Id] = v
	}

	mapGameType2Region = make(map[int]map[int]models.TtArea)
	arrTtAreaRefGame := make([]models.TtAreaRefGame, 0)
	o.QueryTable(mconst.TableName_TtAreaRefGame).All(&arrTtAreaRefGame)

	for _, v := range arrTtAreaRefGame {

		mapArea, ok := mapGameType2Region[v.GameType]
		if ok {
			aTtArea := mpTtArea[v.AreaId]
			mapArea[aTtArea.Id] = aTtArea
			mapGameType2Region[v.GameType] = mapArea
		} else {
			mapArea = make(map[int]models.TtArea)
			aTtArea := mpTtArea[v.AreaId]
			mapArea[aTtArea.Id] = aTtArea
			mapGameType2Region[v.GameType] = mapArea
		}
	}
}

type YHKInfo struct {
	YHName  string
	CardNum string
	Name    string
	Tel     string
	Addr    string
	Cate    string
	Remark  string
}
type PlayerSKXX struct {
	GameID       int
	VoucherFile  string
	WXSKCodeUrl  string
	YHKInfo      YHKInfo
	ZFBSKCodeUrl string
	ZFBSKName    string
}

type ResultNewPlayer struct {
	GameID   int
	MarketID int
	UserName string
}

//http://47.244.125.83:7710/ChangePaw
func ChangePaw(Region string, UserId int, pwd string) error {
	strUrl := fmt.Sprintf(`%s/ChangePaw`,
		mHttp.baseArr)

	type TempS struct {
		Region string
		ID     int
		Psw    string
	}

	t := TempS{Region: Region, ID: UserId, Psw: pwd}

	strJ, e := json.Marshal(t)
	if e != nil {
		return e
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("ChangePaw begin url:", strUrl, "body:", string(strJ))
	var err error
	if !GConfig.GetGConfig().IsDev {
		_, err = mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
	}
	ttLog.LogDebug("ChangePaw end", err)

	return err
}

func SetGameTimer(gameType int, aData httpBox.SetGameTimerBox) {
	gameType2RegionLock.RLock()
	defer gameType2RegionLock.RUnlock()
	strUrl := fmt.Sprintf(`%s/SetGameTimer`,
		mHttp.baseArrZg)

	mapRegion := mapGameType2Region[gameType]
	for _, region := range mapRegion {
		aData.Region = region.Area
		strJ, e := json.Marshal(aData)
		if e != nil {
			ttLog.LogError(e)
		}
		reader := strings.NewReader(string(strJ))
		ttLog.LogWarning("UscGameHttp SetGameTimer begin url:", strUrl, "body:", string(strJ))
		var err error
		if !GConfig.GetGConfig().IsDev {
			_, err = mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
		}
		if err != nil {
			ttLog.LogError(err)
		}

		ttLog.LogWarning("UscGameHttp SetGameTimer end", err)
	}
}

/**
封盘消息
*/
func StopBetState(gameType int, aData httpBox.StopBetStateBox) {
	gameType2RegionLock.RLock()
	defer gameType2RegionLock.RUnlock()

	strUrl := fmt.Sprintf(`%s/StopBetState`,
		mHttp.baseArrZg)
	mapRegion := mapGameType2Region[gameType]
	for _, region := range mapRegion {
		aData.Region = region.Area
		strJ, e := json.Marshal(aData)
		if e != nil {
			ttLog.LogError(e)
		}
		reader := strings.NewReader(string(strJ))
		ttLog.LogWarning("UscGameHttp StopBetState begin url:", strUrl, "body:", string(strJ))
		var err error
		if !GConfig.GetGConfig().IsDev {
			_, err = mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
		}
		if err != nil {
			ttLog.LogError(err)
		}

		ttLog.LogWarning("UscGameHttp StopBetState end", err)
	}
}

func PlayerBet(gameType int, aData httpBox.PlayerBetBox) {
	gameType2RegionLock.RLock()
	defer gameType2RegionLock.RUnlock()

	strUrl := fmt.Sprintf(`%s/PlayerBet`,
		mHttp.baseArrZg)
	mapRegion := mapGameType2Region[gameType]
	for _, region := range mapRegion {
		aData.Region = region.Area
		strJ, e := json.Marshal(aData)
		if e != nil {
			ttLog.LogError(e)
		}
		reader := strings.NewReader(string(strJ))
		ttLog.LogWarning("UscGameHttp PlayerBet begin url:", strUrl, "body:", string(strJ))
		var err error
		if !GConfig.GetGConfig().IsDev {
			_, err = mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
		}
		if err != nil {
			ttLog.LogError(err)
		}
		ttLog.LogWarning("UscGameHttp PlayerBet end", err)
	}
}

func AwardResult(gameType int, aData httpBox.AwardResultBox) {
	gameType2RegionLock.RLock()
	defer gameType2RegionLock.RUnlock()

	strUrl := fmt.Sprintf(`%s/AwardResult`,
		mHttp.baseArrZg)
	mapRegion := mapGameType2Region[gameType]
	for _, region := range mapRegion {
		aData.Region = region.Area
		strJ, e := json.Marshal(aData)
		if e != nil {
			ttLog.LogError(e)
		}
		reader := strings.NewReader(string(strJ))
		ttLog.LogWarning("UscGameHttp AwardResult begin url:", strUrl, "body:", string(strJ))
		var err error
		if !GConfig.GetGConfig().IsDev {
			_, err = mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
		}
		if err != nil {
			ttLog.LogError(err)
		}
		ttLog.LogWarning("UscGameHttp AwardResult end", err)
	}
}

func GetPlayerSKXX(UserId int) (PlayerSKXX, error) {
	aPlayerSKXX := make([]PlayerSKXX, 0)
	strUrl := fmt.Sprintf(`%s/getPlayerSKXX`,
		mHttp.baseArr)

	type TempS struct {
		Region string
		IDS    []int
		Type   int `json:"type"`
	}

	t := TempS{Region: Region, Type: 1}
	t.IDS = make([]int, 0)
	t.IDS = append(t.IDS, UserId)

	strJ, e := json.Marshal(t)
	if e != nil {
		return PlayerSKXX{}, e
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("getPlayerSKXX begin url:", strUrl, "body:", string(strJ))
	body, err := mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
	strBody := string(body)
	ttLog.LogDebug("getPlayerSKXX end", strBody, err)

	e = json.Unmarshal(body, &aPlayerSKXX)
	if len(aPlayerSKXX) == 1 {
		return aPlayerSKXX[0], e
	} else {
		return PlayerSKXX{}, errors.New("长度不对")
	}
}

func SetAwardResult(gameType int, LotteryAward, StrPeriod string) error {
	strUrl := fmt.Sprintf(`%s/ZGGame/SetAwardResult`,
		mHttp.baseArrResult)

	type TempS struct {
		Region       string
		GameType     int `json:"type"`
		LotteryAward string
		StrPeriod    string
	}

	t := TempS{Region: Region, GameType: gameType, LotteryAward: LotteryAward, StrPeriod: StrPeriod}

	strJ, e := json.Marshal(t)
	if e != nil {
		return e
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("SetAwardResult begin url:", strUrl, "body:", string(strJ))
	if !GConfig.GetGConfig().IsDev {
		body, err := mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)

		strBody := string(body)
		ttLog.LogDebug("SetAwardResult end", strBody, err)

		type TempR struct {
			Msg    string
			Result bool
		}
		aTempR := TempR{}
		e = json.Unmarshal(body, &aTempR)
		if e != nil {
			return e
		}
		if !aTempR.Result {
			return errors.New(aTempR.Msg)
		}
		return err
	} else {
		return nil
	}
}
