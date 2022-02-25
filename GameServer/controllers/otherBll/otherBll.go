package otherBll

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/GameServer/GConfig"
	"ttmyth123/GroupLottery/GameServer/GInstance"
	"ttmyth123/GroupLottery/GameServer/GInstance/GTtHint"
	"ttmyth123/GroupLottery/GameServer/OtherServer/httpGameServer"
	"ttmyth123/GroupLottery/GameServer/models"
	"ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	"ttmyth123/kit/httpKit"
	"ttmyth123/kit/pwdKit"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

type GameNameInfo struct {
	GameType int    `json:"1"`
	GameName string `json:"2"`
}

func GetGameNames() []GameNameInfo {
	o := orm.NewOrm()
	arrData := make([]GameNameInfo, 0)
	sql := fmt.Sprintf(`select * from %s`, mconst.TableName_TtGameInfo)
	o.Raw(sql).QueryRows(&arrData)
	return arrData
}

type ArticleInfo struct {
	Id          int
	Title       string //标题
	TitleImgUrl string //标题图片
	Des         string //概要（描述）
	ArticleUrl  string
	ArticleType int
	UpdatedAt   time.Time
}

func (d ArticleInfo) MarshalJSON() ([]byte, error) {
	type Alias ArticleInfo

	StrUpdatedAt := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrUpdatedAt = d.UpdatedAt.Format(timeKit.DateTimeLayout)
	} else {
		StrUpdatedAt = d.UpdatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		UpdatedAt string
	}{
		Alias:     (Alias)(d),
		UpdatedAt: StrUpdatedAt,
	})
}

func GetArticleList(ArticleType, PageIndex, PageSize int) (int, []ArticleInfo) {
	PageTotal := 0
	o := orm.NewOrm()
	arrData := make([]ArticleInfo, 0)
	sqlArgs := make([]interface{}, 0)
	sqlWhere := "where a.state=? "
	sqlArgs = append(sqlArgs, mconst.ArticleState_2_Enabled)

	if ArticleType != 0 {
		sqlWhere += " and a.article_type=?"
		sqlArgs = append(sqlArgs, ArticleType)
	}

	sqlCount := fmt.Sprintf(`select count(1) from %s a %s`, mconst.TableName_TtArticleInfo, sqlWhere)
	o.Raw(sqlCount, sqlArgs).QueryRow(&PageTotal)

	offset, _ := sqlKit.GetOffset(PageTotal, PageSize, PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, PageSize)

	sql := fmt.Sprintf(`select a.id, a.title, a.title_img_url
, a.des, a.article_type, a.updated_at from %s a %s`, mconst.TableName_TtArticleInfo, sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)

	for i := 0; i < len(arrData); i++ {
		arrData[i].TitleImgUrl = httpKit.GetImgUrl(GConfig.GetGConfig().SApiBaseP, arrData[i].TitleImgUrl)
		arrData[i].ArticleUrl = fmt.Sprintf("%s/html/articlehtml?id=%d", GConfig.GetGConfig().SApiBaseP, arrData[i].Id)
	}

	return PageTotal, arrData
}

type ServiceInfo struct {
}

func GetServiceInfo() models.TtServiceInfo {
	aServiceInfo := models.TtServiceInfo{}
	o := orm.NewOrm()
	o.QueryTable(mconst.TableName_TtServiceInfo).One(&aServiceInfo)

	return aServiceInfo
}

func ModifyPwd(Region string, userId int, OldPwd, NewPwd string) error {
	aTtGameUser, e := GInstance.GetUserRpcClient().GetUser(userId)
	if e != nil {
		return errors.New(GTtHint.GetTtHint().GetHint("用户不存在。"))
	}
	aOldPwd := pwdKit.Sha1ToStr(OldPwd)
	if aTtGameUser.Pwd != aOldPwd {
		return errors.New(GTtHint.GetTtHint().GetHint("旧密码不正确。"))
	}

	e = httpGameServer.ChangePaw(Region, userId, NewPwd)
	if e != nil {
		return errors.New(GTtHint.GetTtHint().GetHint("修改密码失败，原因：远程服务出错"))
	}

	updateData := make([]gBox.UpdateDataInfo, 0)
	newPwd1 := pwdKit.Sha1ToStr(NewPwd)
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "Pwd", Type: 0, Value: newPwd1}, gBox.UpdateDataInfo{FieldName: "Pwd2", Type: 0, Value: NewPwd})

	_, e = GInstance.GetUserRpcClient().UpdateUserInfo(userId, updateData)
	return e
}

func ChangePayInfo(UserId int) error {
	aPlayerSKXX, e := httpGameServer.GetPlayerSKXX(UserId)
	if e != nil {
		ttLog.LogError(e)
		return e
	}

	updateData := make([]gBox.UpdateDataInfo, 0)
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "VoucherFile", Type: 0, Value: aPlayerSKXX.VoucherFile})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "WXSKCodeUrl", Type: 0, Value: aPlayerSKXX.WXSKCodeUrl})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "ZFBSKCodeUrl", Type: 0, Value: aPlayerSKXX.ZFBSKCodeUrl})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "ZFBSKName", Type: 0, Value: aPlayerSKXX.ZFBSKName})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "YHName", Type: 0, Value: aPlayerSKXX.YHKInfo.YHName})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "CardNum", Type: 0, Value: aPlayerSKXX.YHKInfo.CardNum})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "YHUserName", Type: 0, Value: aPlayerSKXX.YHKInfo.Name})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "YHUserTel", Type: 0, Value: aPlayerSKXX.YHKInfo.Tel})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "Addr", Type: 0, Value: aPlayerSKXX.YHKInfo.Addr})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "Cate", Type: 0, Value: aPlayerSKXX.YHKInfo.Cate})
	updateData = append(updateData, gBox.UpdateDataInfo{FieldName: "Remark", Type: 0, Value: aPlayerSKXX.YHKInfo.Remark})
	_, e = GInstance.GetUserRpcClient().UpdateUserInfo(UserId, updateData)
	return e
}
