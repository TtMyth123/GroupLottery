package httpGameServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
	"ttmyth123/GroupLottery/Staff/GConfig"
	"ttmyth123/GroupLottery/Staff/OtherServer/httpGameServer/httpBox"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	"ttmyth123/kit/httpClientKit"
	"ttmyth123/kit/ttLog"
)

var (
	mHttp *HttpGameServer
)

const Region = `ZGGame聊天室`

type HttpGameServer struct {
	baseArr     string
	mHttpClient *httpClientKit.HttpClient
}

func Init() {
	mHttp = new(HttpGameServer)
	mHttp.baseArr = beego.AppConfig.String("HttpGameServerUrl")
	mHttp.mHttpClient = httpClientKit.GetHttpClient("")

}

type ResultChat struct {
	Msg    string    `json:"msg"`
	Id     int64     `json:"id"`
	RoomId int       `json:"room_id"`
	Sign   string    `json:"sign"`
	Time   time.Time `json:"time"`
}

func AdminMsg(aData models.TtChatInfo) ResultChat {
	strUrl := fmt.Sprintf(`%s/Chat/%s/AdminMsg`, mHttp.baseArr, GConfig.GetGConfig().Area)

	aResultChat := ResultChat{}
	type TempS struct {
		//Region string `json:"region"`
		RoomId int    `json:"roomid"`
		Type   string `json:"type"`
		Msg    string `json:"msg"`
		Pid    int64  `json:"pid"`
	}

	//{"msg":"消息内容","type":"0","roomid":"180","pid":10018}
	t := TempS{Pid: aData.GameId, RoomId: aData.RoomId, Type: strconv.Itoa(aData.ChatType), Msg: aData.Content}

	strJ, e := json.Marshal(t)
	if e != nil {
		return aResultChat
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("AdminMsg begin url:", strUrl, "body:", string(strJ))
	var err error
	if !GConfig.GetGConfig().IsDev {
		strResultChat, err := mHttp.mHttpClient.DoRequest("POST", strUrl, nil, reader)
		if err == nil {
			err = json.Unmarshal(strResultChat, &aResultChat)
		}
	}
	ttLog.LogDebug("AdminMsg end", err, aResultChat)

	return aResultChat
}

type AdminList struct {
	CreateTime string
	GameID     string
	GameRegion string
	UserName   string
}
type AdminListResult struct {
	List []AdminList `json:"list"`
}

func GetAdminList() AdminListResult {
	strUrl := fmt.Sprintf(`%s/Chat/%s/GetAdminList`, mHttp.baseArr, GConfig.GetGConfig().Area)
	aResult := AdminListResult{}
	ttLog.LogDebug("GetAdminList begin url:", strUrl)
	var err error
	if !GConfig.GetGConfig().IsDev {
		strResult, err := mHttp.mHttpClient.DoRequest("GET", strUrl, nil, nil)
		if err == nil {
			err = json.Unmarshal(strResult, &aResult)
		}
	}
	ttLog.LogDebug("GetAdminList end", err, aResult)

	return aResult
}

type CreatePrivateToResult struct {
	Msg      string `json:"msg"`
	RoomId   int    `json:"room_id"`
	RoomName string `json:"room_name"`
	Sign     string `json:"sign"`
}

func CreatePrivateTo(GameId int64, ToGameId int64) CreatePrivateToResult {
	aCreatePrivateToResult := CreatePrivateToResult{}
	strUrl := fmt.Sprintf(`%s/Chat/%s/CreatePrivateTo`, mHttp.baseArr, GConfig.GetGConfig().Area)

	type TempS struct {
		GameId   int64 `json:"gameid"`
		ToGameId int64 `json:"togameid"`
	}

	//{"gameid":10019,"togameid":10013}
	t := TempS{GameId: GameId, ToGameId: ToGameId}

	strJ, e := json.Marshal(t)
	if e != nil {
		aCreatePrivateToResult.Msg = e.Error()
		return aCreatePrivateToResult
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("CreatePrivateTo begin url:", strUrl, "body:", string(strJ))
	var err error
	if !GConfig.GetGConfig().IsDev {
		strResultChat, err := mHttp.mHttpClient.DoRequest("GET", strUrl, nil, reader)
		if err == nil {
			err = json.Unmarshal(strResultChat, &aCreatePrivateToResult)
		}
	}
	ttLog.LogDebug("CreatePrivateTo end", err, aCreatePrivateToResult)

	return aCreatePrivateToResult
}

type CreateAdminAccountResult struct {
	Msg string `json:"msg"`
}

func CreateAdminAccount(UserName, Password, GameName string) CreateAdminAccountResult {
	aCreateAdminAccountResult := CreateAdminAccountResult{}
	strUrl := fmt.Sprintf(`%s/Chat/%s/CreateAdminAccount`, mHttp.baseArr, GConfig.GetGConfig().Area)

	type TempS struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		GameName string `json:"gameName"`
	}

	//{"gameid":10019,"togameid":10013}
	t := TempS{UserName: UserName, Password: Password, GameName: GameName}

	strJ, e := json.Marshal(t)
	if e != nil {
		aCreateAdminAccountResult.Msg = e.Error()
		return aCreateAdminAccountResult
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("CreateAdminAccount begin url:", strUrl, "body:", string(strJ))
	var err error
	if !GConfig.GetGConfig().IsDev {
		paramsHeader := make(map[string]string)
		paramsHeader["Content-Type"] = "application/json"
		strResultChat, err := mHttp.mHttpClient.DoRequest("POST", strUrl, paramsHeader, reader)
		if err == nil {
			err = json.Unmarshal(strResultChat, &aCreateAdminAccountResult)
		}
	}
	ttLog.LogDebug("CreateAdminAccount end", err, aCreateAdminAccountResult)

	return aCreateAdminAccountResult
}

type RoomInfoResult struct {
}

func GetRoomInfo(roomId int) (httpBox.RoomInfoResult, error) {
	aResult := httpBox.RoomInfoResult{}
	strUrl := fmt.Sprintf(`%s/Chat/%s/GetRoomInfo`, mHttp.baseArr, GConfig.GetGConfig().Area)

	type TempS struct {
		Id string `json:"id"`
	}
	aTempS := TempS{Id: fmt.Sprintf("%d", roomId)}

	strJ, e := json.Marshal(aTempS)
	if e != nil {
		return aResult, e
	}
	reader := strings.NewReader(string(strJ))
	ttLog.LogDebug("GetRoomInfo begin url:", strUrl, "body:", string(strJ))
	var err error
	if !GConfig.GetGConfig().IsDev {
		paramsHeader := make(map[string]string)
		paramsHeader["Content-Type"] = "application/json"
		strResultChat, err := mHttp.mHttpClient.DoRequest(httpClientKit.Method_Get, strUrl, paramsHeader, reader)
		if err == nil {
			err = json.Unmarshal(strResultChat, &aResult)
			if aResult.Msg != "" {

				return aResult, errors.New(aResult.Msg)
			}
		}
	}
	ttLog.LogDebug("GetRoomInfo end", err, aResult)

	return aResult, nil
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

	t := TempS{Region: GConfig.GetGConfig().Area, ID: UserId, Psw: pwd}

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
