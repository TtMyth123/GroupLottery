package UserRpcClient

import (
	"context"
	"errors"
	"github.com/smallnest/rpcx/client"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/AreaConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/RpcConst"
	"ttmyth123/GroupLottery/UserInfoRpc/RpcServer/UserServer/RpcBox"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
)

type RpcClient struct {
	xclient client.XClient
}

func NewRpcClient(addr string) *RpcClient {
	aRpcClient := new(RpcClient)

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	aRpcClient.xclient = client.NewXClient(RpcConst.ServerName,
		client.Failtry, client.RandomSelect, d, client.DefaultOption)

	return aRpcClient
}

func (this *RpcClient) GetUser(UserId int) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsGetUser)
	args.UserId = UserId
	reply := new(RpcBox.ReplyGetUser)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetUser, &args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}

	return reply.Data, nil
}

func (this *RpcClient) AddUser(Area, UserName, Pwd, MoneyPwd string, Code, UserType int, InfoEx map[string]string) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsAddUser)
	args.Area = Area
	args.UserName = UserName
	args.Pwd = Pwd
	args.MoneyPwd = MoneyPwd
	args.Code = Code
	args.UserType = UserType
	args.InfoEx = InfoEx

	reply := new(RpcBox.ReplyAddUser)
	err := this.xclient.Call(context.Background(), RpcConst.M_AddUser, &args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}

	return reply.Data, nil
}

func (this *RpcClient) AddGold(goldInfo gBox.AddGoldInfo) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsAddGold)
	args.UserId = goldInfo.UserId
	args.Gold = goldInfo.Gold
	args.T = goldInfo.T
	args.Des = goldInfo.Des

	reply := new(RpcBox.ReplyAddGold)
	err := this.xclient.Call(context.Background(), RpcConst.M_AddGold, args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}
	return reply.Data, nil
}

func (this *RpcClient) GetUserByNamePwd(UserName, Pwd string) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsGetUserByNamePwd)
	args.UserName = UserName
	args.Pwd = Pwd

	reply := new(RpcBox.ReplyGetUserByNamePwd)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetUserByNamePwd, args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}
	return reply.Data, nil
}

func (this *RpcClient) UpdateUserInfo(userId int, infos []gBox.UpdateDataInfo) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsUpdateUserInfo)
	args.UserId = userId
	args.Datas = infos

	reply := new(RpcBox.ReplyUpdateUserInfo)
	err := this.xclient.Call(context.Background(), RpcConst.M_UpdateUserInfo, args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}
	return reply.Data, nil
}

func (this *RpcClient) Rebate2Gold(UserId int, Rebate float64, Des, Des2, DesMp string) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsRebate2Gold)
	args.UserId = UserId
	args.Rebate = Rebate
	args.Des = Des
	args.Des2 = Des2
	args.DesMp = DesMp

	reply := new(RpcBox.ReplyRebate2Gold)
	err := this.xclient.Call(context.Background(), RpcConst.M_Rebate2Gold, &args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}

	return reply.Data, nil
}
func (this *RpcClient) Rebate2GoldRecord(UserId, PageIndex, PageSize int) ([]RpcBox.Rebate2GoldRecord, interface{}, int, error) {
	args := new(RpcBox.ArgsRebate2GoldRecord)
	args.UserId = UserId
	args.PageIndex = PageIndex
	args.PageSize = PageSize

	reply := new(RpcBox.ReplyRebate2GoldRecord)
	err := this.xclient.Call(context.Background(), RpcConst.M_Rebate2GoldRecord, &args, reply)
	if err != nil {
		return nil, nil, 0, err
	}
	if !reply.Result {
		return nil, nil, 0, errors.New(reply.ErrMsg)
	}

	return reply.DataList, reply.GroupObj, reply.PageTotal, nil
}

func (this *RpcClient) AddRebate(aAddRebateInfo gBox.AddRebateInfo) (models.TtGameUser, error) {
	args := new(RpcBox.ArgsAddRebate)
	args.UserId = aAddRebateInfo.UserId
	args.UserSid = aAddRebateInfo.UserSid
	args.BetUserName = aAddRebateInfo.BetUserName
	args.GameType = aAddRebateInfo.GameType
	args.LotteryStr = aAddRebateInfo.LotteryStr
	args.Level = aAddRebateInfo.Level
	args.OddsName = aAddRebateInfo.OddsName
	args.BetM = aAddRebateInfo.BetM
	args.T = aAddRebateInfo.T

	reply := new(RpcBox.ReplyAddRebate)
	err := this.xclient.Call(context.Background(), RpcConst.M_AddRebate, args, reply)
	if err != nil {
		return models.TtGameUser{}, err
	}
	if !reply.Result {
		return models.TtGameUser{}, errors.New(reply.ErrMsg)

	}
	return reply.Data, nil
}

func (this *RpcClient) GetRebateSetConfig(area int) (AreaConfig.RebateSetConfig, error) {
	args := new(RpcBox.ArgsGetRebateSetConfig)
	args.Area = area
	reply := new(RpcBox.ReplyGetRebateSetConfig)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetRebateSetConfig, args, reply)
	if err != nil {
		return AreaConfig.GetRebateSet(area), err
	}

	return reply.Data, nil
}
func (this *RpcClient) ReRebateSetConfig(area int) error {
	args := new(RpcBox.ArgsReRebateSetConfig)
	reply := new(RpcBox.ReplyReRebateSetConfig)
	err := this.xclient.Call(context.Background(), RpcConst.M_ReRebateSetConfig, args, reply)
	if err != nil {
		return err
	}
	return nil
}

func (this *RpcClient) GetGameUserList(args *RpcBox.ArgsGetGameUserList) RpcBox.ReplyGetGameUserList {
	reply := new(RpcBox.ReplyGetGameUserList)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetGameUserList, args, reply)
	if err != nil {
		return *reply
	}
	return *reply
}

func (this *RpcClient) AddTtChatInfo(args *RpcBox.ArgsAddTtChatInfo) RpcBox.ReplyAddTtChatInfo {
	reply := new(RpcBox.ReplyAddTtChatInfo)
	err := this.xclient.Call(context.Background(), RpcConst.M_AddTtChatInfo, args, reply)
	if err != nil {
		return *reply
	}
	return *reply
}

func (this *RpcClient) GetLatelyChatUser(args *RpcBox.ArgsGetLatelyChatUser) RpcBox.ReplyGetLatelyChatUser {
	reply := new(RpcBox.ReplyGetLatelyChatUser)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetLatelyChatUser, args, reply)
	if err != nil {
		return *reply
	}
	return *reply
}

func (this *RpcClient) GetChatRoomMsgList(args *RpcBox.ArgsGetChatRoomMsgList) RpcBox.ReplyGetChatRoomMsgList {
	reply := new(RpcBox.ReplyGetChatRoomMsgList)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetChatRoomMsgList, args, reply)
	if err != nil {
		return *reply
	}
	return *reply
}

func (this *RpcClient) GetChatStaffMsgList(args *RpcBox.ArgsGetChatStaffMsgList) RpcBox.ReplyGetChatStaffMsgList {
	reply := new(RpcBox.ReplyGetChatStaffMsgList)
	err := this.xclient.Call(context.Background(), RpcConst.M_GetChatStaffMsgList, args, reply)
	if err != nil {
		return *reply
	}
	return *reply
}

func (this *RpcClient) UpdateUserAgentId(UserId int) RpcBox.ReplyBox {
	args := &RpcBox.ArgsBox{UserId: UserId}
	reply := new(RpcBox.ReplyBox)
	err := this.xclient.Call(context.Background(), RpcConst.M_UpdateUserAgentId, args, reply)
	if err != nil {
		return *reply
	}
	return *reply
}
