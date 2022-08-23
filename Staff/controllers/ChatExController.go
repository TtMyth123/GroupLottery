package controllers

import (
	"fmt"
	"github.com/TtMyth123/Staff/GConfig"
	"github.com/TtMyth123/Staff/GInstance"
	"github.com/TtMyth123/Staff/controllers/base"
	"github.com/TtMyth123/Staff/controllers/base/enums"
	"github.com/TtMyth123/Staff/controllers/bll"
	"github.com/TtMyth123/Staff/models/mconst"
	"github.com/TtMyth123/UserInfoRpc/RpcServer/UserServer/RpcBox"
	"github.com/TtMyth123/kit"
	"github.com/TtMyth123/kit/TtErrors"
	"github.com/TtMyth123/kit/httpKit"
	"github.com/TtMyth123/kit/ttLog"
	"strings"
)

type ChatExController struct {
	base.AuthorBaseController
}

/**
SendAdminChat
*/
func (c *ChatExController) SendAdminChat() {
	curStaff := c.CurSysUserEx()
	//Way := mconst.ChatWay_Staff_1
	//GameId, _ := c.GetInt("GameId")
	RoomId, _ := c.GetInt("RoomId")
	ChatType, _ := c.GetInt("ChatType")
	Content := c.GetString("Content")
	RoomName := c.GetString("RoomName")
	e := bll.AddAdminChatMsg(curStaff.BaseSysUser, RoomId, RoomName, ChatType, Content)
	c.JsonResultEx(e, "")
}

func (c *ChatExController) GetChat() {
	Min, _ := c.GetInt64("Min")
	Max, _ := c.GetInt64("Max")
	Way, _ := c.GetInt64("Way")
	curStaff := c.CurSysUserEx()
	msgs := bll.GetChat(curStaff.Id, Way, Min, Max)
	c.JsonResultEx(nil, msgs)
}

func (c *ChatExController) GetRoomChat() {
	Min, _ := c.GetInt64("Min")
	Max, _ := c.GetInt64("Max")
	Way, _ := c.GetInt64("Way")
	RoomId, _ := c.GetInt("RoomId")
	curStaff := c.CurSysUserEx()
	msgs := bll.GetRoomChat(curStaff.Id, RoomId, Way, Min, Max)
	c.JsonResultEx(nil, msgs)
	//c.JsonResult(enums.JRCodeSucc, "", msgs)
}

func (c *ChatExController) SendPic() {
	curStaff := c.CurSysUserEx()
	ttLog.LogDebug("SendPic")
	strFilePath := ""
	imgFile, head, err := c.GetFile("File")
	if err == nil {
		filePath, _ := httpKit.UploadFile(imgFile, head, GConfig.GetGConfig().SavePicRootPath, kit.GetGuid(), "")

		if len(filePath) > 0 && filePath[0:1] == `/` {
			i := strings.LastIndex(filePath, `/`)
			str := filePath[i:]
			strFilePath = fmt.Sprintf(`%s%s`, GConfig.GetGConfig().SavePicPath, str)
		} else {
			strFilePath = filePath
		}
	}

	//Way := mconst.ChatWay_Staff_1
	//GameId,_ := c.GetInt("GameId")
	RoomId, _ := c.GetInt("RoomId")
	RoomName := c.GetString("RoomName")
	ChatType := mconst.ChatType_Pic_1
	Content := strFilePath

	e := bll.AddAdminChatMsg(curStaff.BaseSysUser, RoomId, RoomName, ChatType, Content)
	c.JsonResultEx(e, "")
	//if e!= nil {
	//	c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	//}

	//c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ChatExController) GetLatelyChatUser() {
	MaxId, _ := c.GetInt("MaxId")
	PageIndex, _ := c.GetInt("PageIndex")
	PageSize, _ := c.GetInt("PageSize")
	UserName := c.GetString("UserName")
	args := new(RpcBox.ArgsGetLatelyChatUser)
	args.MaxId = MaxId
	args.PageIndex = PageIndex
	args.PageSize = PageSize
	args.UserName = UserName

	r := GInstance.GetUserRpcClient().GetLatelyChatUser(args)

	if !r.Result {
		c.JsonResultEx(TtErrors.New(r.ErrMsg), "")
	}

	type ResultList struct {
		DataList interface{}
		GroupObj interface{}
	}
	aResultList := ResultList{
		DataList: r.DataList,
		GroupObj: r.GroupObj,
	}
	c.JsonResultEx(nil, aResultList)
}

func (c *ChatExController) GetChatRoomMsgList() {
	RoomId, _ := c.GetInt("RoomId")
	MaxId, _ := c.GetInt64("MaxId")
	MinId, _ := c.GetInt64("MinId")
	Count, _ := c.GetInt("Count")
	Way, _ := c.GetInt("Way")

	args := new(RpcBox.ArgsGetChatRoomMsgList)
	args.RoomId = RoomId
	args.MaxId = MaxId
	args.MinId = MinId
	args.Count = Count
	args.Way = Way

	r := GInstance.GetUserRpcClient().GetChatRoomMsgList(args)
	if !r.Result {
		c.JsonResultEx(TtErrors.New(r.ErrMsg), "")
	}

	type ResultList struct {
		DataList interface{}
		GroupObj interface{}
	}
	aResultList := ResultList{
		DataList: r.DataList,
		GroupObj: r.GroupObj,
	}
	c.JsonResultEx(nil, aResultList)
}

func (c *ChatExController) GetChatRoomSMsg() {
	RoomId, _ := c.GetInt("RoomId")
	MaxId, _ := c.GetInt64("MaxId")
	MinId, _ := c.GetInt64("MinId")
	Count, _ := c.GetInt("Count")
	Way, _ := c.GetInt("Way")

	args := new(RpcBox.ArgsGetChatRoomMsgList)
	args.RoomId = RoomId
	args.MaxId = MaxId
	args.MinId = MinId
	args.Count = Count
	args.Way = Way

	r := GInstance.GetUserRpcClient().GetChatRoomMsgList(args)
	if !r.Result {
		c.JsonResult(enums.JRCodeFailed, r.ErrMsg, "")
	}

	type ResultList struct {
		DataList interface{}
		GroupObj interface{}
	}
	aResultList := ResultList{
		DataList: r.DataList,
		GroupObj: r.GroupObj,
	}
	c.JsonResult(enums.JRCodeSucc, r.ErrMsg, aResultList)
}

func (c *ChatExController) GetContactList() {
	MaxId, _ := c.GetInt("MaxId")
	PageIndex, _ := c.GetInt("PageIndex")
	PageSize, _ := c.GetInt("PageSize")
	UserName := c.GetString("UserName")
	data, groupData := bll.GetContactList(MaxId, PageIndex, PageSize, UserName)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", groupData.C, data, groupData)
}

func (c *ChatExController) GetReplyContentList() {
	curStaff := c.CurSysUserEx()
	MaxId, _ := c.GetInt("MaxId")
	PageIndex, _ := c.GetInt("PageIndex")
	PageSize, _ := c.GetInt("PageSize")
	data, groupData := bll.GetReplyContentList(curStaff.Id, MaxId, PageIndex, PageSize)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", groupData.C, data, groupData)
}

func (c *ChatExController) ModifyReplyContent() {
	curStaff := c.CurSysUserEx()
	Id, _ := c.GetInt64("Id")
	MainKey := c.GetString("MainKey")
	e := bll.ModifyReplyContent(Id, curStaff.Id, MainKey)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
func (c *ChatExController) DelReplyContent() {
	Id, _ := c.GetInt64("Id")
	e := bll.DelReplyContent(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ChatExController) GetReplySubContentList() {
	MaxId, _ := c.GetInt("MaxId")
	MainId, _ := c.GetInt("MainId")
	PageIndex, _ := c.GetInt("PageIndex")
	PageSize, _ := c.GetInt("PageSize")
	data, groupData := bll.GetReplySubContentList(MainId, MaxId, PageIndex, PageSize)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", groupData.C, data, groupData)
}

func (c *ChatExController) DelReplySubContent() {
	Id, _ := c.GetInt64("Id")
	e := bll.DelReplySubContent(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ChatExController) ModifyReplySubContent() {
	Id, _ := c.GetInt64("Id")
	MainId, _ := c.GetInt64("MainId")
	ReplyContent := c.GetString("ReplyContent")
	e := bll.ModifyReplySubContent(Id, MainId, ReplyContent)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ChatExController) GetQuickReplyList() {
	curStaff := c.CurSysUserEx()
	MaxId, _ := c.GetInt("MaxId")
	PageIndex, _ := c.GetInt("PageIndex")
	PageSize, _ := c.GetInt("PageSize")
	data, groupData := bll.GetQuickReplyList(curStaff.Id, MaxId, PageIndex, PageSize)
	c.JsonPageResultGroup(enums.JRCodeSucc, "", groupData.C, data, groupData)
}

func (c *ChatExController) DelQuickReply() {
	Id, _ := c.GetInt64("Id")
	e := bll.DelQuickReply(Id)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}

func (c *ChatExController) ModifyQuickReply() {
	curStaff := c.CurSysUserEx()
	Id, _ := c.GetInt64("Id")
	ReplyC := c.GetString("ReplyC")
	e := bll.ModifyQuickReply(Id, curStaff.Id, ReplyC)
	if e != nil {
		c.JsonResult(enums.JRCodeFailed, e.Error(), "")
	}
	c.JsonResult(enums.JRCodeSucc, "", "")
}
