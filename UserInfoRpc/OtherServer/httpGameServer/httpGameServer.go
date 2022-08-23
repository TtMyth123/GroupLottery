package httpGameServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego"
	"strings"
)

var (
	mHttp *HttpGameServer
)

//const Region = ""
type HttpGameServer struct {
	baseArrZg   string
	baseArr     string
	baseArrUser string
	mHttpClient *httpClientKit.HttpClient

	mHttpClientGold *httpClientKit.HttpClient
}

func init() {
	mHttp = new(HttpGameServer)
	//http://47.244.125.83:7799
	mHttp.baseArr = beego.AppConfig.String("HttpGameServerUrl")
	mHttp.baseArrZg = beego.AppConfig.String("HttpGameServerUrlZg")
	mHttp.baseArrUser = beego.AppConfig.String("HttpUserServerUrl")

	mHttp.mHttpClient = httpClientKit.GetHttpClient("")
	mHttp.mHttpClientGold = httpClientKit.GetHttpClient("")
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

func AddMoney(RoomID, GameID int, Money float64) (float64, error) {
	strUrl := fmt.Sprintf(`%s/Interior/AddMoney?RoomID=%d&GameID=%d&Money=%f`, mHttp.baseArrUser, RoomID, GameID, Money)
	r, e := mHttp.mHttpClientGold.GetBytes(strUrl)
	if e != nil {
		return 0, e
	}
	//{"Result":false,"Msg":"玩家不存在","Data":null}
	type TAddM struct {
		Result bool
		Msg    string
		Data   float64
	}
	aTAddM := new(TAddM)
	e = json.Unmarshal(r, aTAddM)
	if e != nil {
		return 0, e
	}
	if !aTAddM.Result {
		return 0, fmt.Errorf(aTAddM.Msg)
	}

	return aTAddM.Data, e
}

func NewPlayer(Area, UserName string, Psw string, Referrer int) (ResultNewPlayer, error) {
	aResultNewPlayer := ResultNewPlayer{}
	strUrl := fmt.Sprintf(`%s/NewPlayer`,
		mHttp.baseArr)

	type TempS struct {
		Region   string
		UserName string
		Psw      string
		Referrer int
	}

	t := TempS{Region: Area, UserName: UserName, Psw: Psw, Referrer: Referrer}

	strJ, e := json.Marshal(t)
	if e != nil {
		return aResultNewPlayer, e
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("NewPlayer begin url:", strUrl, "body:", string(strJ))
	body, err := mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
	if err != nil {
		return aResultNewPlayer, e
	}

	strBody := string(body)
	ttLog.LogDebug("NewPlayer end", strBody, err)

	e = json.Unmarshal(body, &aResultNewPlayer)
	if e != nil {
		return aResultNewPlayer, errors.New(strBody)
	}

	return aResultNewPlayer, e
}

func GetPlayerSKXX(UserId int, Region string) (PlayerSKXX, error) {
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

//http://47.244.125.83:7710/ChangePaw
func ChangePaw(UserId int, pwd, Region string) error {
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
	body, err := mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
	strBody := string(body)
	ttLog.LogDebug("ChangePaw end", strBody, err)

	return err
}
