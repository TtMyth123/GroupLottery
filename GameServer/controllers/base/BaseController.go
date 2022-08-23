package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TtMyth123/GameServer/GConfig"
	"github.com/TtMyth123/GameServer/GInstance/GTtHint"
	"github.com/TtMyth123/GameServer/controllers/base/TtError"
	"github.com/TtMyth123/GameServer/controllers/base/enums"
	"github.com/TtMyth123/kit/beegoCacheKit"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var isDev = false

type BaseController struct {
	beego.Controller
}
type BaseSysUser struct {
	UserId int
	Name   string
}

func (this *BaseController) JsonResultCPage(code enums.JsonResultCode, msg string, PageTotal int, arrData interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	type TmpR struct {
		PageTotal int
		LastId    int
		DataList  interface{}
	}
	r := TmpR{PageTotal: PageTotal, DataList: arrData, LastId: 0}
	this.JsonResult(code, msg, r, nil)
}

func (this *BaseController) JsonResultCPageEx(code enums.JsonResultCode, msg string, PageTotal int, arrData interface{}, OtherInfo interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	type TmpR struct {
		PageTotal int
		LastId    int
		DataList  interface{}
		OtherInfo interface{}
	}
	r := TmpR{PageTotal: PageTotal, DataList: arrData, LastId: 0, OtherInfo: OtherInfo}
	this.JsonResult(code, msg, r, nil)
}

func (this *BaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}, mp interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	a := ArgsBox{}
	this.GetJsonData(&a)
	res := JsonResult{Code: code, Msg: msg, Obj: obj, MsgAdd: a.MsgAdd, Mp: mp}
	aTmp := ReplyBox{Result: true, Data: obj}
	if code != enums.JRCodeSucc {
		aTmp.Result = false
		aTmp.ErrMsg = msg
	}

	res.Obj = aTmp
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) JsonListResult(code enums.JsonResultCode, msg string, LastId int, obj interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}

	a := ArgsBox{}
	this.GetJsonData(&a)
	res := JsonResult{Code: code, Msg: msg, Obj: obj, MsgAdd: a.MsgAdd}
	aTmp := ReplyListBox{Result: true, ListData: obj, LastId: LastId}
	if code != enums.JRCodeSucc {
		aTmp.Result = false
		aTmp.ErrMsg = msg
	}

	res.Obj = aTmp
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

//func (this *BaseController) ErrJsonResult(err ...string) {
//
//	arr := make([]string,0)
//	arr = append(arr,err...)
//
//	strE := ""
//	for i:=0;i<len(arr);i++{
//		strE+= " "+arr[i]
//	}
//
//	this.JsonResult(enums.JRCodeFailed, strE,"",nil)
//}

func (this *BaseController) ErrJsonResultEx(err error) {
	e, ok := err.(*TtError.TtError)
	if ok {
		if GConfig.GetGConfig().IsI18n {
			if e.Mp != nil {
				this.JsonResult(enums.JRCodeFailed, GTtHint.GetTtHint().GetHint(e.E), "", stringKit.GetJsonStr(e.Mp))
			} else {
				this.JsonResult(enums.JRCodeFailed, GTtHint.GetTtHint().GetHint(e.E), "", nil)
			}
		}
	}

	this.JsonResult(enums.JRCodeFailed, GTtHint.GetTtHint().GetHint(err.Error()), "", nil)
}

func (this *BaseController) GetJsonData(v interface{}) error {
	strJson := this.GetString("jsonData")
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName, "jsonData:", strJson)
	e := json.Unmarshal([]byte(strJson), v)
	ttLog.LogDebug("Data:", stringKit.GetJsonStr(v))
	return e
}

func (this *BaseController) GetToken(userInfo BaseSysUser) (string, error) {
	claims := new(JWTClaims)
	claims.UserID = userInfo.UserId
	claims.UserName = userInfo.Name
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(beegoCacheKit.TimeoutOneDay)).Unix()

	return getToken(claims)
}
func (this *BaseController) Refresh(claims *JWTClaims) (string, error) {
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	claims.IssuedAt = time.Now().Unix()
	return getToken(claims)
}

func getToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(GTtHint.GetTtHint().GetHint(ErrorReason_ServerBusy))
	}
	return signedToken, nil
}

func (this *BaseController) verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(GTtHint.GetTtHint().GetHint(ErrorReason_ServerBusy))
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(GTtHint.GetTtHint().GetHint(ErrorReason_ReLogin))
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(GTtHint.GetTtHint().GetHint(ErrorReason_ReLogin))
	}
	return claims, nil
}
