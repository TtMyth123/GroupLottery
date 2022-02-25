package chatBll

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer/RpcBox"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/ttLog"
)

type RoomChatInfo struct {
	MinChatId int64
	MaxChatId int64
	ArrMsg    []models.TtChatInfo
}

/**
一个房间最近几条聊天数据
*/
func GetRoomChatInfo(o orm.Ormer, RoomId, count int) RoomChatInfo {
	arrData := make([]models.TtChatInfo, 0)
	sql := fmt.Sprintf(`select * from %s a where a.room_id=? order by a.id LIMIT ?`, mconst.TableName_TtChatInfo)
	c, e := o.Raw(sql, RoomId, count).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(e)
	}

	aRoomChatInfo := RoomChatInfo{ArrMsg: arrData}
	if c > 0 {
		aRoomChatInfo.MinChatId = aRoomChatInfo.ArrMsg[0].Id
		aRoomChatInfo.MaxChatId = aRoomChatInfo.ArrMsg[c-1].Id
	}

	return aRoomChatInfo
}

type LatelyChatUser struct {
	models.TtLatelyChatUser
	ChatInfo RoomChatInfo
}

func GetLatelyChatUser(args *RpcBox.ArgsGetLatelyChatUser) (interface{}, interface{}, error) {
	type GroupData struct {
		C int
	}

	aGroupData := GroupData{}
	sqlArgs := make([]interface{}, 0)
	arrData := make([]models.TtLatelyChatUser, 0)

	o := orm.NewOrm()

	maxId := args.MaxId
	if args.MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtLatelyChatUser)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}
	sqlWhere := ` a.id <=?`
	sqlArgs = append(sqlArgs, maxId)

	if args.UserName != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.user_name)>0`
		sqlArgs = append(sqlArgs, args.UserName)
	}

	sqlCount := fmt.Sprintf(`select count(1) c from %s a where %s`, mconst.TableName_TtLatelyChatUser, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		return make([]LatelyChatUser, 0), aGroupData, e
	}

	offset, _ := sqlKit.GetOffset(aGroupData.C, args.PageSize, args.PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, args.PageSize)

	sql := fmt.Sprintf(`select * from %s a where %s`, mconst.TableName_TtLatelyChatUser, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
	}
	//aLatelyChatUser := make([]LatelyChatUser,c)
	//for i,item := range arrData{
	//	data := GetRoomChatInfo(o,item.RoomId,20)
	//	aLatelyChatUser[i].TtLatelyChatUser = item
	//	aLatelyChatUser[i].ChatInfo = data
	//}

	return arrData, aGroupData, nil
}
