package bll

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/Staff/GData"
	"ttmyth123/GroupLottery/Staff/GData/chat"
	"ttmyth123/GroupLottery/Staff/GInstance"
	"ttmyth123/GroupLottery/Staff/OtherServer/httpGameServer"
	"ttmyth123/GroupLottery/Staff/controllers/base/box"
	"ttmyth123/GroupLottery/Staff/controllers/bll/bllBox"
	staffModels "ttmyth123/GroupLottery/Staff/models"
	"ttmyth123/GroupLottery/Staff/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer/RpcBox"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	userConst "ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/TtErrors"
	"ttmyth123/kit/sqlKit"
	"ttmyth123/kit/timeKit"
	"ttmyth123/kit/ttLog"
)

func AddAdminChatMsg(aStaff box.BaseSysUser, RoomId int, RoomName string, ChatType int, Content string) TtErrors.TtError {
	aTtChatInfo := models.TtChatInfo{UserName: aStaff.Name,
		Way: mconst.ChatWay_Staff_1, RoomId: RoomId, GameId: aStaff.Id, ChatType: ChatType,
		RoomName: RoomName,
		Content:  Content}

	r := httpGameServer.AdminMsg(aTtChatInfo)
	switch r.Msg {
	case "":
		aTtChatInfo.Id2 = r.Id
		aTtChatInfo.CreatedAt = r.Time
		aTtChatInfo.RoomId = r.RoomId

		args := new(RpcBox.ArgsAddTtChatInfo)
		args.ChatInfo = aTtChatInfo
		aReply := GInstance.GetUserRpcClient().AddTtChatInfo(args)

		if !aReply.Result {
			ttLog.LogError(aReply.ErrMsg)
		}

		GData.GetChatManage().AddAdminChatMsg(aReply.DataInfo)
		return nil
	case "房间不存在或已解散":
		o := orm.NewOrm()
		sql := fmt.Sprintf("delete from %s where room_id=?", userConst.TableName_TtLatelyChatUser)
		//o.Raw(sql,aTtChatInfo.RoomId)
		o.Raw(sql, aTtChatInfo.RoomId).Exec()

		sql = fmt.Sprintf("delete from %s where room_id=?", userConst.TableName_TtChatInfo)
		//o.Raw(sql,aTtChatInfo.RoomId)
		o.Raw(sql, aTtChatInfo.RoomId).Exec()
		return TtErrors.NewError(r.Msg, 1)
		//aCreatePrivateToResult := httpGameServer.CreatePrivateTo(aStaff.Id,int64(GameId))
		//if aCreatePrivateToResult.Msg == "" {
		//	aTtChatInfo.RoomId = aCreatePrivateToResult.RoomId
		//	r1 := httpGameServer.AdminMsg(aTtChatInfo)
		//	if r1.Msg!= "" {
		//		return errors.New(r1.Msg)
		//	}
		//} else {
		//	return errors.New(aCreatePrivateToResult.Msg)
		//}
	}

	return TtErrors.New(r.Msg)
}

func GetChat(staffId int64, Way, Min, Max int64) chat.GroupChatMsg {
	msgs := GData.GetChatManage().GetAllChatMsg(staffId, Way, Min, Max)
	return msgs
}

func GetRoomChat(staffId int64, roomId int, Way, Min, Max int64) chat.GroupChatMsg {
	msgs := GData.GetChatManage().GetRoomChatMsg(staffId, roomId, Way, Min, Max)
	return msgs
}

/**
修复聊天数据。
*/
func repairChatMsg(aTtChatInfo models.TtChatInfo) ([]models.TtChatInfo, error) {
	arrData := make([]models.TtChatInfo, 0)
	if aTtChatInfo.GameId2 == 0 || aTtChatInfo.RoomName == "" || aTtChatInfo.UserName2 == "" {
		//获取聊天信息中的客户信息
		aRoomInfo, e := httpGameServer.GetRoomInfo(aTtChatInfo.RoomId)
		if e != nil {
			return arrData, e
		}

		RoomName := aTtChatInfo.RoomName
		lenMember := len(aRoomInfo.RoomInfo.MemberInfo)
		if lenMember <= 2 {
			RoomName = aTtChatInfo.UserName
		} else {
			if RoomName == "" {
				RoomName = aRoomInfo.RoomInfo.Name
			}
		}

		for userId, userInfo := range aRoomInfo.RoomInfo.MemberInfo {
			aStaff := GData.GetStaffDataManage().GetStaffData(int64(userId))
			if aStaff != nil {
				aNewTtChatInfo := models.TtChatInfo{
					Id2:       aTtChatInfo.Id2,
					RoomId:    aTtChatInfo.RoomId,
					RoomName:  RoomName,
					ChatType:  aTtChatInfo.ChatType,
					Content:   aTtChatInfo.Content,
					Way:       aTtChatInfo.Way,
					GameId:    aTtChatInfo.GameId,
					UserName:  aTtChatInfo.UserName,
					GameId2:   int64(userId),
					UserName2: userInfo.Name,
				}

				arrData = append(arrData, aNewTtChatInfo)
			}
		}
	} else {
		aStaff := GData.GetStaffDataManage().GetStaffData(aTtChatInfo.GameId2)
		if aStaff != nil {
			aNewTtChatInfo := models.TtChatInfo{
				Id2:       aTtChatInfo.Id2,
				RoomId:    aTtChatInfo.RoomId,
				RoomName:  aTtChatInfo.RoomName,
				ChatType:  aTtChatInfo.ChatType,
				Content:   aTtChatInfo.Content,
				Way:       aTtChatInfo.Way,
				GameId:    aTtChatInfo.GameId,
				UserName:  aTtChatInfo.UserName,
				GameId2:   int64(aTtChatInfo.GameId2),
				UserName2: aTtChatInfo.UserName2,
			}

			arrData = append(arrData, aNewTtChatInfo)
		}
	}
	return arrData, nil
}

/**
客户端 给服务端发信息
*/
func AddChatMsg(Id int64, RoomId int, RoomName string, ChatType int, Content,
	CreatedAt string, GameId int, UserName string, GameId2 int, UserName2 string) {
	ttLog.LogDebug("AddChatMsg:", Id, RoomId, RoomName, ChatType, Content, CreatedAt, GameId, UserName, GameId2, UserName2)
	aCreatedAt, e := time.ParseInLocation(CreatedAt, timeKit.DateTimeLayout, time.Local)
	if e != nil {
		aCreatedAt = time.Now()
	}

	Way := mconst.ChatWay_User_0
	aTtChatInfo := models.TtChatInfo{
		Id2:       Id,
		RoomId:    RoomId,
		RoomName:  RoomName,
		ChatType:  ChatType,
		Content:   Content,
		CreatedAt: aCreatedAt,
		Way:       Way,

		GameId:    int64(GameId),
		UserName:  UserName,
		GameId2:   int64(GameId2),
		UserName2: UserName2,
	}
	arrChatInfo, e := repairChatMsg(aTtChatInfo)
	if e != nil {
		ttLog.LogError("repairChatMsg:", aTtChatInfo, e)
		return
	}

	for _, aChatInfo := range arrChatInfo {
		args := new(RpcBox.ArgsAddTtChatInfo)
		args.ChatInfo = aChatInfo
		aReply := GInstance.GetUserRpcClient().AddTtChatInfo(args)
		if !aReply.Result {
			ttLog.LogError(aReply.ErrMsg)
		}

		GData.GetChatManage().AddGameChatMsg(aReply.DataInfo)
	}
}

func GetContactList(MaxId, PageIndex, PageSize int, UserName string) ([]bllBox.ContactInfo, bllBox.GroupContactInfo) {
	arrData := make([]bllBox.ContactInfo, 0)
	aGroupData := bllBox.GroupContactInfo{}
	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	if MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, userConst.TableName_TtGameUser)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&MaxId)
		if e != nil {
			ttLog.LogError(e)
			return arrData, aGroupData
		}
	}
	sqlWhere := " where a.user_type=? "
	sqlArgs = append(sqlArgs, userConst.UserType_1)

	if UserName != "" {
		sqlWhere = sqlWhere + ` and locate(?,a.user_name)>0`
		sqlArgs = append(sqlArgs, UserName)
	}
	sqlCount := fmt.Sprintf("select count(1) c from %s a %s ", userConst.TableName_TtGameUser, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		ttLog.LogError(e)
		return arrData, aGroupData
	}

	offset, _ := sqlKit.GetOffset(aGroupData.C, PageSize, PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, PageSize)
	sql := fmt.Sprintf(`select a.id, a.user_name from %s a %s`, userConst.TableName_TtGameUser, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return arrData, aGroupData
	}

	return arrData, aGroupData
}

type ReplyContentInfo struct {
	Id           int
	MainKey      string
	ReplyContent string
}
type GroupReplyContentInfo struct {
	C int
}

func GetReplyContentList(staffId int64, MaxId, PageIndex, PageSize int) ([]ReplyContentInfo, GroupReplyContentInfo) {
	arrData := make([]ReplyContentInfo, 0)
	aGroupData := GroupReplyContentInfo{}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	if MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtReplyMain)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&MaxId)
		if e != nil {
			ttLog.LogError(e)
			return arrData, aGroupData
		}
	}
	sqlWhere := " where a.staff_id=? and a.id <=?"
	sqlArgs = append(sqlArgs, staffId, MaxId)
	sqlCount := fmt.Sprintf("select count(1) c from %s a %s ", mconst.TableName_TtReplyMain, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		ttLog.LogError(e)
		return arrData, aGroupData
	}
	offset, _ := sqlKit.GetOffset(aGroupData.C, PageSize, PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, PageSize)
	sql := fmt.Sprintf(`
select a.id, a.main_key, b.reply_content from %s a LEFT JOIN (
select main_id,group_concat(reply_content separator  "\n") as reply_content from %s group by main_id
) b on (a.id=b.main_id) %s
`, mconst.TableName_TtReplyMain, mconst.TableName_TtReplySub, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return arrData, aGroupData
	}

	return arrData, aGroupData
}
func ModifyReplyContent(Id, StaffId int64, MainKey string) error {
	if Id == 0 {
		o := orm.NewOrm()
		c, e := o.QueryTable(mconst.TableName_TtReplyMain).Filter("MainKey", MainKey).Filter("StaffId", StaffId).Count()
		if e != nil {
			return e
		}
		if c > 0 {
			return errors.New(fmt.Sprintf("[%s]已存在", MainKey))
		}
		aTtReplyMain := staffModels.TtReplyMain{StaffId: StaffId, MainKey: MainKey}
		e = aTtReplyMain.Add(o)
		chat.GetReplyManage().ReloadAll()
		return e
	} else {
		o := orm.NewOrm()
		data := make([]staffModels.TtReplyMain, 0)
		c, _ := o.QueryTable(mconst.TableName_TtReplyMain).Filter("MainKey", MainKey).Filter("StaffId", StaffId).All(&data)

		if c == 1 {
			if data[0].Id != Id {
				return errors.New(fmt.Sprintf("[%s]已存在", MainKey))
			}
			data[0].MainKey = MainKey
			e := data[0].Update(o, "MainKey", "UpdateAt")
			chat.GetReplyManage().ReloadAll()
			return e
		} else if c > 1 {
			return errors.New(fmt.Sprintf("有问题的数据[%s]", MainKey))
		} else {
			aTtReplyMain := staffModels.TtReplyMain{}
			aTtReplyMain.Id = Id
			aTtReplyMain.MainKey = MainKey
			e := aTtReplyMain.Update(o, "MainKey", "UpdateAt")
			chat.GetReplyManage().ReloadAll()
			return e

		}
	}
	return nil
}
func DelReplyContent(Id int64) error {
	o := orm.NewOrm()
	_, e := o.QueryTable(mconst.TableName_TtReplyMain).Filter("Id", Id).Delete()
	chat.GetReplyManage().ReloadAll()
	if e != nil {
		return e
	}
	return nil
}

type ReplySubContentInfo struct {
	Id           int
	MainId       int
	ReplyContent string
}

func GetReplySubContentList(mainId int, MaxId, PageIndex, PageSize int) ([]ReplySubContentInfo, GroupReplyContentInfo) {
	arrData := make([]ReplySubContentInfo, 0)
	aGroupData := GroupReplyContentInfo{}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	if MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtReplySub)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&MaxId)
		if e != nil {
			ttLog.LogError(e)
			return arrData, aGroupData
		}
	}
	sqlWhere := " where a.main_id=? and a.id <=?"
	sqlArgs = append(sqlArgs, mainId, MaxId)
	sqlCount := fmt.Sprintf("select count(1) c from %s a %s ", mconst.TableName_TtReplySub, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		ttLog.LogError(e)
		return arrData, aGroupData
	}
	offset, _ := sqlKit.GetOffset(aGroupData.C, PageSize, PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, PageSize)
	sql := fmt.Sprintf(` select * from %s a %s`, mconst.TableName_TtReplySub, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return arrData, aGroupData
	}

	return arrData, aGroupData
}

func DelReplySubContent(Id int64) error {
	o := orm.NewOrm()
	_, e := o.QueryTable(mconst.TableName_TtReplySub).Filter("Id", Id).Delete()
	chat.GetReplyManage().ReloadAll()
	if e != nil {
		return e
	}

	return nil
}
func ModifyReplySubContent(Id, MainId int64, ReplyContent string) error {
	o := orm.NewOrm()
	if Id == 0 {
		aTtReplySub := staffModels.TtReplySub{MainId: MainId, ReplyContent: ReplyContent}
		aTtReplySub.Id = Id
		e := aTtReplySub.Add(o)
		chat.GetReplyManage().ReloadAll()
		return e
	} else {
		aTtReplySub := staffModels.TtReplySub{MainId: MainId, ReplyContent: ReplyContent}
		aTtReplySub.Id = Id
		e := aTtReplySub.Update(o, "ReplyContent", "UpdateAt")
		chat.GetReplyManage().ReloadAll()
		return e
	}

	return nil
}

type QuickReplyInfo struct {
	Id     int64
	ReplyC string
}
type GroupQuickReplyInfo struct {
	C int
}

func GetQuickReplyList(staffId int64, MaxId, PageIndex, PageSize int) ([]QuickReplyInfo, GroupQuickReplyInfo) {
	arrData := make([]QuickReplyInfo, 0)
	aGroupData := GroupQuickReplyInfo{}

	o := orm.NewOrm()
	sqlArgs := make([]interface{}, 0)

	if MaxId == 0 {
		sqlCount := fmt.Sprintf(`select max(a.id) as maxid from %s a `, mconst.TableName_TtQuickReply)
		e := o.Raw(sqlCount, sqlArgs).QueryRow(&MaxId)
		if e != nil {
			ttLog.LogError(e)
			return arrData, aGroupData
		}
	}
	sqlWhere := " where a.staff_id=? and a.id <=?"
	sqlArgs = append(sqlArgs, staffId, MaxId)
	sqlCount := fmt.Sprintf("select count(1) c from %s a %s ", mconst.TableName_TtQuickReply, sqlWhere)
	e := o.Raw(sqlCount, sqlArgs).QueryRow(&aGroupData)
	if e != nil {
		ttLog.LogError(e)
		return arrData, aGroupData
	}
	offset, _ := sqlKit.GetOffset(aGroupData.C, PageSize, PageIndex)
	sqlWhere = sqlWhere + ` LIMIT ?,?`
	sqlArgs = append(sqlArgs, offset, PageSize)
	sql := fmt.Sprintf("select a.*  from %s a %s ", mconst.TableName_TtQuickReply, sqlWhere)
	c, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)
	if e != nil {
		ttLog.LogError(c, e)
		return arrData, aGroupData
	}

	return arrData, aGroupData
}

func DelQuickReply(Id int64) error {
	o := orm.NewOrm()
	_, e := o.QueryTable(mconst.TableName_TtQuickReply).Filter("Id", Id).Delete()
	if e != nil {
		return e
	}
	return nil
}

func ModifyQuickReply(Id, StaffId int64, ReplyC string) error {
	o := orm.NewOrm()
	if Id == 0 {
		aTtQuickReply := staffModels.TtQuickReply{StaffId: StaffId, ReplyC: ReplyC}
		aTtQuickReply.Id = Id
		e := aTtQuickReply.Add(o)
		return e
	} else {
		aTtQuickReply := staffModels.TtQuickReply{ReplyC: ReplyC}
		aTtQuickReply.Id = Id
		e := aTtQuickReply.Update(o, "ReplyC", "UpdateAt")
		return e
	}

	return nil
}
