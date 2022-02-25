package UserServer

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/orm"
	"ttmyth123/GroupLottery/UserInfoRpc/GData"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/AreaConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer/RpcBox"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer/chatBll"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/ttLog"
)

type UserServer struct {
}

func NewUserServer() *UserServer {
	aUserServer := new(UserServer)
	return aUserServer
}

func (this *UserServer) GetUserByNamePwd(ctx context.Context, args *RpcBox.ArgsGetUserByNamePwd, reply *RpcBox.ReplyGetUserByNamePwd) error {
	aUser, e := GData.GetUserByNamePwd(args.UserName, args.Pwd)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Result = true
	reply.Data = *aUser
	return nil
}
func (this *UserServer) GetUser(ctx context.Context, args *RpcBox.ArgsGetUser, reply *RpcBox.ReplyGetUser) error {
	aUser, e := GData.GetUser(args.UserId)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Data = aUser
	reply.Result = true
	return nil
}

func (this *UserServer) AddUser(ctx context.Context, args *RpcBox.ArgsAddUser, reply *RpcBox.ReplyAddUser) error {
	aUser, e := GData.AddUser(args.Area, args.UserName, args.Pwd, args.MoneyPwd, args.Code, args.UserType, args.InfoEx)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Data = aUser
	reply.Result = true
	return nil
}

func (this *UserServer) AddGold(ctx context.Context, args *RpcBox.ArgsAddGold, reply *RpcBox.ReplyAddGold) error {
	goldInfo := gBox.AddGoldInfo{GroupId: args.GroupId, UserId: args.UserId, Gold: args.Gold, T: args.T, Des: args.Des}
	aUser, e := GData.AddGold(goldInfo)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Data = aUser
	reply.Result = true
	return nil
}

func (this *UserServer) UpdateUserInfo(ctx context.Context, args *RpcBox.ArgsUpdateUserInfo, reply *RpcBox.ReplyUpdateUserInfo) error {
	aUser, e := GData.UpdateUserInfo(args.UserId, args.Datas)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Data = aUser
	reply.Result = true
	return nil
}

func (this *UserServer) Rebate2Gold(ctx context.Context, args *RpcBox.ArgsRebate2Gold, reply *RpcBox.ReplyRebate2Gold) error {
	aUser, e := GData.Rebate2Gold(args.UserId, args.Rebate, args.Des, args.Des2, args.DesMp)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Data = aUser
	reply.Result = true
	return nil
}

func (this *UserServer) UpdateUserAgentId(ctx context.Context, args *RpcBox.ArgsBox, reply *RpcBox.ReplyBox) error {
	GData.UpdateUserAgentId(args.UserId)
	reply.Result = true
	return nil
}
func (this *UserServer) Rebate2GoldRecord(ctx context.Context, args *RpcBox.ArgsRebate2GoldRecord, reply *RpcBox.ReplyRebate2GoldRecord) error {
	type GroupData struct {
		C     int
		MaxId int
	}

	aGroupData := GroupData{}
	sqlArgs := make([]interface{}, 0)
	arrData := make([]RpcBox.Rebate2GoldRecord, 0)

	o := orm.NewOrm()

	maxId := args.MaxId
	if args.MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtRebateInfo)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&maxId)
		if e != nil {
			ttLog.LogError(e)
		}
	}
	sqlWhere := `a.user_id=? and a.rebate_type=? and a.id <=?`
	sqlArgs = append(sqlArgs, args.UserId, mconst.Rebate_02_ToGold, maxId)

	sqlCount := fmt.Sprintf(`select count(1) c from %s a where %s`, mconst.TableName_TtRebateInfo, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		reply.ErrMsg = e.Error()
		reply.Result = false

		return e
	}

	offset, _ := sqlKit.GetOffset(aGroupData.C, args.PageSize, args.PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, args.PageSize)

	sql := fmt.Sprintf(`select a.id, -a.gold as gold, a.created_at from %s a where %s`, mconst.TableName_TtRebateInfo, sqlWhere)
	_, e = o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		reply.ErrMsg = e.Error()
		reply.Result = false

		return e
	}
	aGroupData.MaxId = maxId
	reply.DataList = arrData
	reply.GroupObj = aGroupData
	reply.PageTotal = aGroupData.C
	reply.Result = true
	return nil

}
func (this *UserServer) AddRebate(ctx context.Context, args *RpcBox.ArgsAddRebate, reply *RpcBox.ReplyRebate2Gold) error {
	aAddRebateInfo := gBox.AddRebateInfo{
		UserId:      args.UserId,
		UserSid:     args.UserSid,
		BetUserName: args.BetUserName,
		GameType:    args.GameType,
		LotteryStr:  args.LotteryStr,
		Level:       args.Level,
		OddsName:    args.OddsName,
		BetM:        args.BetM,
		T:           args.T,
	}
	aUser, e := GData.AddRebate(aAddRebateInfo)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
		return nil
	}

	reply.Data = aUser
	reply.Result = true
	return nil
}

func (this *UserServer) GetRebateSetConfig(ctx context.Context, args *RpcBox.ArgsGetRebateSetConfig, reply *RpcBox.ReplyGetRebateSetConfig) error {
	aRebateSet := AreaConfig.GetRebateSet(args.Area)

	reply.Data = aRebateSet
	reply.Result = true
	return nil
}
func (this *UserServer) ReRebateSetConfig(ctx context.Context, args *RpcBox.ArgsReRebateSetConfig, reply *RpcBox.ReplyReRebateSetConfig) error {
	AreaConfig.ReLoadRebateSet(0)
	reply.Result = true
	return nil
}

func (this *UserServer) GetGameUserList(ctx context.Context, args *RpcBox.ArgsGetGameUserList, reply *RpcBox.ReplyGetGameUserList) error {
	type GroupData struct {
		C int
	}

	aGroupData := GroupData{}
	sqlArgs := make([]interface{}, 0)
	arrData := make([]models.TtGameUser, 0)

	o := orm.NewOrm()

	maxId := args.MaxId
	if args.MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtGameUser)
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

	sqlCount := fmt.Sprintf(`select count(1) c from %s a where %s`, mconst.TableName_TtGameUser, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		reply.ErrMsg = e.Error()
		reply.Result = false

		return e
	}

	offset, _ := sqlKit.GetOffset(aGroupData.C, args.PageSize, args.PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, args.PageSize)

	sql := fmt.Sprintf(`select * from %s a where %s`, mconst.TableName_TtGameUser, sqlWhere)
	o.Raw(sql, sqlArgs).QueryRows(&arrData)
	reply.DataList = arrData
	reply.GroupObj = aGroupData

	return nil
}

func (this *UserServer) AddTtChatInfo(ctx context.Context, args *RpcBox.ArgsAddTtChatInfo, reply *RpcBox.ReplyAddTtChatInfo) error {
	o := orm.NewOrm()
	e := args.ChatInfo.Add(o)
	if e != nil {
		reply.Result = false
		reply.ErrMsg = e.Error()
	}

	aTtLatelyChatUser := models.TtLatelyChatUser{Id: int64(args.ChatInfo.RoomId)}
	e = o.Read(&aTtLatelyChatUser)
	if e != nil {
		aTtLatelyChatUser.Id = int64(args.ChatInfo.RoomId)
		aTtLatelyChatUser.RoomId = args.ChatInfo.RoomId
		aTtLatelyChatUser.RoomName = args.ChatInfo.RoomName
		aTtLatelyChatUser.UserName = args.ChatInfo.UserName

		aTtLatelyChatUser.Add(o)
	} else {
		aTtLatelyChatUser.Update(o, "UpdateAt")
	}

	reply.Result = true
	reply.DataInfo = args.ChatInfo
	return nil
}

func (this *UserServer) GetLatelyChatUser(ctx context.Context, args *RpcBox.ArgsGetLatelyChatUser, reply *RpcBox.ReplyGetLatelyChatUser) error {

	DataList, GroupObj, e := chatBll.GetLatelyChatUser(args)
	if e == nil {
		reply.Result = true
		reply.DataList = DataList
		reply.GroupObj = GroupObj
	}

	return e
}

func (this *UserServer) GetChatRoomMsgList(ctx context.Context, args *RpcBox.ArgsGetChatRoomMsgList, reply *RpcBox.ReplyGetChatRoomMsgList) error {
	aGroupData := RpcBox.GetChatRoomMsgListGroup{MinId: args.MinId, MaxId: args.MaxId}
	sqlArgs := make([]interface{}, 0)
	arrData := make([]models.TtChatInfo, 0)

	o := orm.NewOrm()

	sqlWhere := ` and a.room_id =? `
	sqlArgs = append(sqlArgs, args.RoomId)
	if args.Way >= 0 {
		if args.MaxId != 0 {
			sqlWhere += ` and  a.id >? `
			sqlArgs = append(sqlArgs, args.MaxId)
		}

	} else {
		if args.MinId != 0 {
			sqlWhere += ` and a.id < ?`
			sqlArgs = append(sqlArgs, args.MinId)
		}
	}

	//sqlCount := fmt.Sprintf(`select count(1) c, min(a.id) min_id, max(a.id) as max_id from %s a, %s b where a.room_id=b.room_id %s`,
	//	mconst.TableName_TtChatInfo, mconst.TableName_TtLatelyChatUser, sqlWhere)
	//e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	//if e != nil {
	//	ttLog.LogError(e)
	//}

	//order by a.created_at DESC,
	if args.Way > 0 {
		sqlWhere += `order by a.id `
	} else {
		sqlWhere += `order by a.id DESC `
	}
	sqlWhere = sqlWhere + ` LIMIT ?`
	sqlArgs = append(sqlArgs, args.Count)

	sql := fmt.Sprintf(`select a.* from %s a, %s b where a.room_id=b.room_id %s`,
		mconst.TableName_TtChatInfo, mconst.TableName_TtLatelyChatUser, sqlWhere)
	c, _ := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if c > 0 {
		if aGroupData.MinId != 0 {
			if aGroupData.MinId > arrData[0].Id {
				aGroupData.MinId = arrData[0].Id
			}
		} else {
			aGroupData.MinId = arrData[0].Id
		}

		if aGroupData.MaxId < arrData[c-1].Id {
			aGroupData.MaxId = arrData[c-1].Id
		}
	}

	reply.DataList = arrData
	reply.GroupObj = aGroupData
	reply.Result = true
	return nil
}

func (this *UserServer) GetChatStaffMsgList(ctx context.Context, args *RpcBox.ArgsGetChatStaffMsgList, reply *RpcBox.ReplyGetChatStaffMsgList) error {
	aGroupData := RpcBox.GetChatStaffMsgListGroup{MinId: args.MinId, MaxId: args.MaxId}
	sqlArgs := make([]interface{}, 0)
	//arrData := make([]models.TtLatelyChatUser, 0)
	arrData := make([]models.TtChatInfo, 0)

	sqlWhere := ` a.room_id=b.room_id `
	o := orm.NewOrm()
	if args.Way >= 0 {
		if args.MaxId != 0 {
			sqlWhere += ` and a.id >? `
			sqlArgs = append(sqlArgs, args.MaxId)
		}
	} else {
		if args.MinId != 0 {
			sqlWhere += `and a.id < ?`
			sqlArgs = append(sqlArgs, args.MinId)
		}
	}

	//sqlCount := fmt.Sprintf(`select count(1) c, min(a.id) min_id, max(a.id) as max_id from %s a, %s b where %s`,
	//	mconst.TableName_TtChatInfo, mconst.TableName_TtLatelyChatUser, sqlWhere)
	//e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	//if e != nil {
	//	ttLog.LogError(e)
	//}
	if args.Way == 0 {
		sqlWhere += `order by a.id `
	} else {
		sqlWhere += `order by a.id DESC `
	}
	sqlWhere = sqlWhere + ` LIMIT ?`
	sqlArgs = append(sqlArgs, args.Count)

	sql := fmt.Sprintf(`select a.id, a.id2, a.room_id,
a.letter,b.room_name ,a.game_id,a.user_name ,a.game_id2,a.user_name2,a.way,a.chat_type,
a.content,a.created_at,a.update_at from %s a, %s b where %s`,
		mconst.TableName_TtChatInfo, mconst.TableName_TtLatelyChatUser, sqlWhere)
	c, _ := o.Raw(sql, sqlArgs).QueryRows(&arrData)

	if c > 0 {
		if aGroupData.MinId != 0 {
			if aGroupData.MinId > arrData[0].Id {
				aGroupData.MinId = arrData[0].Id
			}
		} else {
			aGroupData.MinId = arrData[0].Id
		}

		if aGroupData.MaxId < arrData[c-1].Id {
			aGroupData.MaxId = arrData[c-1].Id
		}
	}

	reply.DataList = arrData
	reply.GroupObj = aGroupData
	reply.Result = true

	return nil
}
