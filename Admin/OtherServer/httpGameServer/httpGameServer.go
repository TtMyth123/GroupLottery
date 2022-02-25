package httpGameServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/ttLog"
)

var (
	mHttp *HttpGameServer
)

const Region = "瞳瞳"

type HttpGameServer struct {
	baseArr     string
	baseArrZg   string
	mHttpClient *httpClientKit.HttpClient
}

func init() {
	mHttp = new(HttpGameServer)
	//http://47.244.125.83:7799
	mHttp.baseArr = beego.AppConfig.String("HttpGameServerUrl")
	mHttp.baseArrZg = beego.AppConfig.String("HttpGameServerUrlZg")
	mHttp.mHttpClient = httpClientKit.GetHttpClient("")
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

//http://47.244.125.83:7710/ChangePaw
func ChangePaw(UserId int, pwd string) error {
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
