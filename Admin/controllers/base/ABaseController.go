package base

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
	"ttmyth123/GroupLottery/Admin/controllers/base/enums"
	"ttmyth123/kit/beegoCacheKit"
	"ttmyth123/kit/stringKit"
	"ttmyth123/kit/ttLog"
)

type ABaseController struct {
	beego.Controller
}
type BaseSysUser struct {
	UserId int
	Name   string
}

func (this *ABaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName)
}

func (this *ABaseController) GetJsonData(v interface{}) error {
	strJson := this.GetString("jsonData")
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName, "jsonData:", strJson)
	e := json.Unmarshal([]byte(strJson), v)
	ttLog.LogDebug("Data:", stringKit.GetJsonStr(v))
	return e
}

func (this *ABaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	res := JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonListResult(code enums.JsonResultCode, msg string, lastId int, obj interface{}) {
	listR := ListResult{LastId: lastId, ListData: obj}
	res := JsonListResult{Code: code, Msg: msg, Obj: listR}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonPageResult(code enums.JsonResultCode, msg string, PageTotal int, obj interface{}) {
	listR := PageResult{PageTotal: PageTotal, ListData: obj}
	res := JsonPageResult{Code: code, Msg: msg, Obj: listR}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) JsonPageResultGroup(code enums.JsonResultCode, msg string, PageTotal int, obj interface{}, groupObj interface{}) {
	listR := PageResult{PageTotal: PageTotal, ListData: obj, GroupData: groupObj}
	res := JsonPageResult{Code: code, Msg: msg, Obj: listR}
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *ABaseController) GetToken(userInfo BaseSysUser) (string, error) {
	claims := new(JWTClaims)
	claims.UserID = userInfo.UserId
	claims.UserName = userInfo.Name
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(beegoCacheKit.TimeoutOneDay)).Unix()

	return getToken(claims)
}
func (this *ABaseController) Refresh(claims *JWTClaims) (string, error) {
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	claims.IssuedAt = time.Now().Unix()
	return getToken(claims)
}

func getToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}

func (this *ABaseController) VerifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	return claims, nil
}
