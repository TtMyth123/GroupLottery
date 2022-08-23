package GameClientHttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/controllers/base"
	"github.com/TtMyth123/GameServer/controllers/base/enums"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/astaxie/beego"
)

var (
	mHttp *YhClentHttpServer
)

type YhClentHttpServer struct {
	baseArr     string
	mHttpClient *httpClientKit.HttpClient
}

func init() {
	mHttp = new(YhClentHttpServer)
	//http://47.244.125.83:7799
	mHttp.baseArr = beego.AppConfig.String("YhClientHttpUrl")
	mHttp.mHttpClient = httpClientKit.GetHttpClient("")
}
func DelDrawMoney(DrawMoneyId, userId int, Excuse string) error {
	type TmpArgs struct {
		base.ArgsBox
		DrawMoneyId int
		Excuse      string
	}
	aTmpArgs := TmpArgs{}
	aTmpArgs.DrawMoneyId = DrawMoneyId
	aTmpArgs.UserId = userId
	aTmpArgs.Excuse = Excuse

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/deldrawmoney?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func AgreeSaveMoney(UserId, SaveMoneyId, AuditorId int, AuditorName string) error {
	type TmpArgs struct {
		base.ArgsBox
		SaveMoneyId int
		AuditorName string
		AuditorId   int
	}
	aTmpArgs := TmpArgs{}
	aTmpArgs.SaveMoneyId = SaveMoneyId
	aTmpArgs.AuditorId = AuditorId
	aTmpArgs.AuditorName = AuditorName
	aTmpArgs.UserId = UserId

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/agreesavemoney?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func SaveMoney(UserId, Money, AuditorId int, AuditorName string) error {
	type TmpArgs struct {
		base.ArgsBox
		Money       int
		AuditorName string
		AuditorId   int
	}
	aTmpArgs := TmpArgs{
		Money: Money, AuditorId: AuditorId, AuditorName: AuditorName,
	}
	aTmpArgs.UserId = UserId

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/savemoney?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}
func DrawMoney(UserId, Money, AuditorId int, AuditorName string) error {
	type TmpArgs struct {
		base.ArgsBox
		Money       int
		AuditorName string
		AuditorId   int
	}
	aTmpArgs := TmpArgs{
		Money: Money, AuditorId: AuditorId, AuditorName: AuditorName,
	}
	aTmpArgs.UserId = UserId

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/drawmoney?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}
func AgreeDrawMoney(SaveMoneyId, AuditorId int, AuditorName string) error {
	type TmpArgs struct {
		base.ArgsBox
		DrawMoneyId int
		AuditorName string
		AuditorId   int
	}
	aTmpArgs := TmpArgs{
		DrawMoneyId: SaveMoneyId, AuditorId: AuditorId, AuditorName: AuditorName,
	}
	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/agreedrawmoney?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func ReloadUserInfo(userId int, Fields []string) error {
	type TmpArgs struct {
		base.ArgsBox
		Fields []string
	}

	aTmpArgs := TmpArgs{}
	aTmpArgs.UserId = userId
	aTmpArgs.Fields = Fields

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/reloaduserinfo?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func OddsChange(gameType int) error {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	aTmpArgs.GameType = gameType

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/oddschange?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func RandNewAwardInfo(gameType int) error {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}
	aTmpArgs.GameType = gameType

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/randnewawardinfo?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func ReLoadRebateSet(area int) error {
	type TmpArgs struct {
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{}

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/reloadrebateset?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}

func SetAwardInfo(gameType int, LotteryAward, LotteryNum string) error {
	type TmpArgs struct {
		LotteryAward string
		LotteryNum   string
		base.ArgsBox
	}
	aTmpArgs := TmpArgs{LotteryAward: LotteryAward, LotteryNum: LotteryNum}
	aTmpArgs.GameType = gameType

	strArgs, e := json.Marshal(aTmpArgs)
	if e != nil {
		return e
	}

	strUrl := fmt.Sprintf("%s/api/setawardinfo?jsonData=%s", mHttp.baseArr, string(strArgs))
	r, e := mHttp.mHttpClient.GetBytes(strUrl)
	if e != nil {
		return e
	}
	aReplyBox := base.JsonResult{}
	e = json.Unmarshal(r, &aReplyBox)
	if e != nil {
		return e
	}

	if aReplyBox.Code != enums.JRCodeSucc {
		return errors.New(aReplyBox.Msg)
	}
	return e
}
