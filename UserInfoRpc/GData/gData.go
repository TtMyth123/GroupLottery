package GData

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"sync"
	"time"
	game_models "ttmyth123/GroupLottery/GameServer/models"
	game_mconst "ttmyth123/GroupLottery/GameServer/models/mconst"
	"ttmyth123/GroupLottery/UserInfoRpc/GConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/GData/gBox"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/AreaConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/GInstance/GTtHint"
	"ttmyth123/GroupLottery/UserInfoRpc/OtherServer/httpGameServer"
	"ttmyth123/GroupLottery/UserInfoRpc/models"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit"
	"ttmyth123/kit/strconvEx"
	"ttmyth123/kit/ttLog"
)

var (
	lotteryGData LotteryGData
	userLock     sync.Mutex
	addUserLock  sync.Mutex

	areaLock      sync.RWMutex
	addRebateLock sync.Mutex
)

type GUserInfo struct {
	User     *models.TtGameUser
	DataLock sync.Mutex
}

type LotteryGData struct {
	userMap sync.Map
	areaMap sync.Map
}

func getArea(Area string) models.TtArea {
	areaLock.Lock()
	defer areaLock.Unlock()
	if aArea, ok := lotteryGData.areaMap.Load(Area); ok {
		aTtArea := aArea.(models.TtArea)
		return aTtArea
	} else {
		aTtArea, e := models.GetAddArea(Area)
		if e != nil {
			return aTtArea
		}

		lotteryGData.areaMap.Store(Area, aTtArea)
		return aTtArea
	}
}

func AddUser(Area, UserName, pwd, MoneyPwd string, Code, UserType int, InfoEx map[string]string) (models.TtGameUser, error) {
	addUserLock.Lock()
	defer addUserLock.Unlock()
	if MoneyPwd == "" {
		if GConfig.GetGConfig().RegDrawMoneyPwd {
			return models.TtGameUser{}, errors.New("资金密码不能空")
		}
	}
	if UserType == 0 {
		UserType = mconst.UserType_1
	}
	aNewUser, e := httpGameServer.NewPlayer(Area, UserName, pwd, Code)
	if e != nil {
		return models.TtGameUser{}, e
	}
	aArea := getArea(Area)

	aUser, e := models.AddGameUser(aArea, aNewUser.GameID, UserName, pwd, MoneyPwd, Code, aNewUser.MarketID, UserType, InfoEx)
	if e != nil {
		return aUser, errors.New(e.Error() + "aaaaa")
	}
	if _, ok := lotteryGData.userMap.Load(aUser.Id); !ok {
		aGUser := new(GUserInfo)
		aGUser.User = &aUser
		ttLog.LogDebug("new User:", aUser.Id)
		lotteryGData.userMap.Store(aUser.Id, aGUser)
	}

	return aUser, e
}
func getUser(userId int) (models.TtGameUser, error) {
	aGameUser, e := models.GetUserInfo(userId)
	return aGameUser, e
}
func getUserAndStore(UserId int) (*GUserInfo, error) {
	userLock.Lock()
	defer userLock.Unlock()
	if user, ok := lotteryGData.userMap.Load(UserId); ok {
		aUser := user.(*GUserInfo)
		return aUser, nil
	} else {
		aGameUser, e := getUser(UserId)
		if e != nil {
			return nil, e
		}
		aUser := new(GUserInfo)
		aUser.User = &aGameUser
		lotteryGData.userMap.Store(UserId, aUser)
		return aUser, e
	}
}
func GetUserByNamePwd(userName, pwd string) (*models.TtGameUser, error) {
	aUser, e := models.GetGameUserByNamePwd(userName, pwd)

	return &aUser, e
}
func GetUser(UserId int) (models.TtGameUser, error) {

	user, e := getUserAndStore(UserId)
	if e != nil {
		return models.TtGameUser{}, e
	}
	return *user.User, nil
}
func getUserResult(user *GUserInfo, e error) (models.TtGameUser, error) {
	if user == nil {
		return models.TtGameUser{}, e
	}
	return *user.User, e
}

func AddGold(goldInfo gBox.AddGoldInfo) (models.TtGameUser, error) {
	aUser, e := getUserAndStore(goldInfo.UserId)
	if e != nil {
		return getUserResult(aUser, e)
	}
	aUser.DataLock.Lock()
	defer aUser.DataLock.Unlock()

	if goldInfo.Gold < 0 {
		return getUserResult(aUser, errors.New("参数值不对"))
	}

	switch goldInfo.T {
	case mconst.Account_01_Guess:
		//竞猜
		if aUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, errors.New("余额不足"))
		}

		f, e := httpGameServer.AddMoney(goldInfo.GroupId, goldInfo.UserId, goldInfo.Gold)

		aUser.User.Gold -= goldInfo.Gold
		aUser.User.SumBet += goldInfo.Gold
		aUser.User.SumBet = kit.Decimal(aUser.User.SumBet)
		aUser.User.Gold = kit.Decimal(aUser.User.Gold)
		aUser.User.Update(nil, "Gold", "SumBet")

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}
	//case mconst.Account_10_SubGuessRebate:
	//	//下级竞猜佣金(上级)
	//
	//	//aUser.User.Gold += goldInfo.Gold
	//	aUser.User.Rebate += goldInfo.Gold
	//	aUser.User.SumRebate += goldInfo.Gold
	//	aUser.User.Update(nil, "Rebate", "SumRebate")
	//
	//	aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
	//		StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
	//		Gold: goldInfo.Gold,Des2:goldInfo.Des2, DesMp: goldInfo.DesMp,
	//	}
	//	aTtAccount.Add(nil)

	//case mconst.Account_11_SubGuessRebate:
	//	//下级竞猜佣金(下级)
	//	aUser.User.Sum2Rebate += goldInfo.Gold
	//	aUser.User.Update(nil, "Sum2Rebate")

	//case mconst.Account_12_XmGuess:
	//	//下级竞猜佣金(下级)
	//	aUser.User.SumXmBet += goldInfo.Gold
	//	aUser.User.Update(nil, "SumXmBet")

	case mconst.Account_02_Win:
		//赢得
		aUser.User.Gold += goldInfo.Gold
		aUser.User.SumWin += goldInfo.Gold

		aUser.User.SumWin = kit.Decimal(aUser.User.SumWin)
		aUser.User.Gold = kit.Decimal(aUser.User.Gold)
		aUser.User.Update(nil, "Gold", "SumWin")

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}
	case mconst.Account_14_NotOpen:
		//不开奖退款
		aUser.User.Gold += goldInfo.Gold
		aUser.User.Gold = kit.Decimal(aUser.User.Gold)
		aUser.User.Update(nil, "Gold")

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}
	case mconst.Account_03_SaveMoney:
		PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
		if e1 != nil {
			return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
		}
		if PUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, fmt.Errorf("代理商[%d],余额不足：%g", aUser.User.Pid, goldInfo.Gold))
		}
		PUser.User.Gold -= goldInfo.Gold
		PUser.User.Gold = kit.Decimal(PUser.User.Gold)
		e = PUser.User.Update(nil, "Gold")
		if e != nil {
			return getUserResult(aUser, e)
		} else {
			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_03_SaveMoney2),
					StrType:     mconst.GetAccountName(mconst.Account_03_SaveMoney2),
					Des:         fmt.Sprintf("用户%d[%s]'充值'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold),
					CurUserGold: PUser.User.Gold,
					Gold:        goldInfo.Gold,
					Des2:        fmt.Sprintf("用户%d[%s]'充值'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold), DesMp: "",
				}
				aTtAccount.Add(nil)
			}
		}

		//充值
		aUser.User.Gold += goldInfo.Gold
		aUser.User.SumSaveMoney += goldInfo.Gold
		if aUser.User.MaxSaveMoney < goldInfo.Gold {
			aUser.User.MaxSaveMoney = goldInfo.Gold

			aUser.User.MaxSaveMoney = kit.Decimal(aUser.User.MaxSaveMoney)
			aUser.User.SumSaveMoney = kit.Decimal(aUser.User.SumSaveMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumSaveMoney", "MaxSaveMoney")
		} else {
			aUser.User.SumSaveMoney = kit.Decimal(aUser.User.SumSaveMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumSaveMoney")
		}

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}

	case mconst.Account_04_DrawMoney:
		//提现
		if aUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, errors.New("余额不足"))
		}

		PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
		if e1 != nil {
			return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
		}
		PUser.User.Gold += goldInfo.Gold
		PUser.User.Gold = kit.Decimal(PUser.User.Gold)
		e = PUser.User.Update(nil, "Gold")
		if e != nil {
			return getUserResult(aUser, e)
		} else {
			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_04_DrawMoney2),
					StrType:     mconst.GetAccountName(mconst.Account_04_DrawMoney2),
					Des:         fmt.Sprintf("用户%d[%s]'提现'增加代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold),
					CurUserGold: PUser.User.Gold,
					Gold:        goldInfo.Gold,
					Des2:        fmt.Sprintf("用户%d[%s]'提现'增加代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold), DesMp: "",
				}
				aTtAccount.Add(nil)
			}
		}

		aUser.User.Gold -= goldInfo.Gold
		aUser.User.SumDrawMoney += goldInfo.Gold
		if aUser.User.MaxDrawMoney < goldInfo.Gold {
			aUser.User.MaxDrawMoney = goldInfo.Gold

			aUser.User.MaxDrawMoney = kit.Decimal(aUser.User.MaxDrawMoney)
			aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumDrawMoney", "MaxDrawMoney")
		} else {
			aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumDrawMoney")
		}

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}

	case mconst.Account_07_DrawMoneyR:
		//提现拒绝
		PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
		if e1 != nil {
			return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
		}
		if PUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, fmt.Errorf("代理商[%d],余额不足：%g", aUser.User.Pid, goldInfo.Gold))
		}
		PUser.User.Gold -= goldInfo.Gold
		PUser.User.Gold = kit.Decimal(PUser.User.Gold)
		e = PUser.User.Update(nil, "Gold")
		if e != nil {
			return getUserResult(aUser, e)
		} else {
			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_07_DrawMoneyR2),
					StrType:     mconst.GetAccountName(mconst.Account_07_DrawMoneyR2),
					Des:         fmt.Sprintf("用户%d[%s]'提现拒绝'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold),
					CurUserGold: PUser.User.Gold,
					Gold:        goldInfo.Gold,
					Des2:        fmt.Sprintf("用户%d[%s]'提现拒绝'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold), DesMp: "",
				}
				aTtAccount.Add(nil)
			}
		}

		aUser.User.Gold += goldInfo.Gold

		aUser.User.SumDrawMoney -= goldInfo.Gold
		aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
		aUser.User.Gold = kit.Decimal(aUser.User.Gold)
		aUser.User.Update(nil, "Gold", "SumDrawMoney")

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}

	case mconst.Account_08_AddMoney:
		//上分
		PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
		if e1 != nil {
			return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
		}
		if PUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, fmt.Errorf("代理商[%d],余额不足：%g", aUser.User.Pid, goldInfo.Gold))
		}
		PUser.User.Gold -= goldInfo.Gold
		PUser.User.Gold = kit.Decimal(PUser.User.Gold)
		e = PUser.User.Update(nil, "Gold")
		if e != nil {
			return getUserResult(aUser, e)
		} else {
			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_08_AddMoney2),
					StrType:     mconst.GetAccountName(mconst.Account_08_AddMoney2),
					Des:         fmt.Sprintf("用户%d[%s]'上分'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold),
					CurUserGold: PUser.User.Gold,
					Gold:        goldInfo.Gold,
					Des2:        fmt.Sprintf("用户%d[%s]'上分'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold), DesMp: "",
				}
				aTtAccount.Add(nil)
			}
		}

		aUser.User.Gold += goldInfo.Gold
		aUser.User.SumAddMoney += goldInfo.Gold
		if aUser.User.MaxAddMoney < goldInfo.Gold {
			aUser.User.SumAddMoney = kit.Decimal(aUser.User.SumAddMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.MaxAddMoney = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumAddMoney", "MaxAddMoney")
		} else {
			aUser.User.SumAddMoney = kit.Decimal(aUser.User.SumAddMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumAddMoney")
		}

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}

	case mconst.Account_09_DecMoney:
		//下分
		if aUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, errors.New("余额不足"))
		}

		PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
		if e1 != nil {
			return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
		}
		PUser.User.Gold += goldInfo.Gold
		PUser.User.Gold = kit.Decimal(PUser.User.Gold)
		e = PUser.User.Update(nil, "Gold")
		if e != nil {
			return getUserResult(aUser, e)
		} else {
			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_09_DecMoney2),
					StrType:     mconst.GetAccountName(mconst.Account_09_DecMoney2),
					Des:         fmt.Sprintf("用户%d[%s]' '增加代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold),
					CurUserGold: PUser.User.Gold,
					Gold:        goldInfo.Gold,
					Des2:        fmt.Sprintf("用户%d[%s]'下分'增加代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold), DesMp: "",
				}
				aTtAccount.Add(nil)
			}
		}

		aUser.User.Gold -= goldInfo.Gold
		aUser.User.SumDecMoney += goldInfo.Gold
		if aUser.User.MaxDecMoney < goldInfo.Gold {
			aUser.User.SumDecMoney = kit.Decimal(aUser.User.SumDecMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.MaxDecMoney = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumDecMoney", "MaxDecMoney")
		} else {
			aUser.User.SumDecMoney = kit.Decimal(aUser.User.SumDecMoney)
			aUser.User.Gold = kit.Decimal(aUser.User.Gold)
			aUser.User.Update(nil, "Gold", "SumDecMoney")
		}

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}

	case mconst.Account_05_Give:
		//赠送
		PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
		if e1 != nil {
			return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
		}
		if PUser.User.Gold < goldInfo.Gold {
			return getUserResult(aUser, fmt.Errorf("代理商[%d],余额不足：%g", aUser.User.Pid, goldInfo.Gold))
		}
		PUser.User.Gold -= goldInfo.Gold
		PUser.User.Gold = kit.Decimal(PUser.User.Gold)
		e = PUser.User.Update(nil, "Gold")
		if e != nil {
			return getUserResult(aUser, e)
		} else {
			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_05_Give2),
					StrType:     mconst.GetAccountName(mconst.Account_05_Give2),
					Des:         fmt.Sprintf("用户%d[%s]'赠送'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold),
					CurUserGold: PUser.User.Gold,
					Gold:        goldInfo.Gold,
					Des2:        fmt.Sprintf("用户%d[%s]'赠送'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, goldInfo.Gold), DesMp: "",
				}
				aTtAccount.Add(nil)
			}
		}

		aUser.User.Gold += goldInfo.Gold

		aUser.User.Gold = kit.Decimal(aUser.User.Gold)
		aUser.User.Update(nil, "Gold")

		if goldInfo.Gold != 0 {
			aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
				StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
				Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
			}
			aTtAccount.Add(nil)
		}
	}

	return getUserResult(aUser, e)
}
func UpdateUserAgentId(userId int) {
	o := orm.NewOrm()
	type T struct {
		Id int
	}
	ids := make([]int, 0)
	ids = append(ids, userId)
	i := 0
	for {
		if i >= len(ids) {
			return
		}
		arrData := make([]T, 0)
		sql := fmt.Sprintf(`select a.id from %s a where a.pid=?`, mconst.TableName_TtGameUser)
		o.Raw(sql, ids[i]).QueryRows(&arrData)

		for _, uId := range arrData {
			aUser, e := getUserAndStore(uId.Id)
			if e == nil {
				aUser.User.AgentUserId = userId
				aUser.User.Update(o, "AgentUserId")
				ids = append(ids, uId.Id)
			}
		}
		i++
	}

}

func UpdateUserInfo(userId int, infos []gBox.UpdateDataInfo) (models.TtGameUser, error) {
	aUser, e := getUserAndStore(userId)
	if e != nil {
		return getUserResult(aUser, e)
	}
	aUser.DataLock.Lock()
	defer aUser.DataLock.Unlock()

	arr := make([]string, 0)
	for _, item := range infos {
		switch item.FieldName {
		case "Pid":
			if item.Type == 0 {
				aUser.User.Pid = item.Value.(int)
			} else {
				aUser.User.Pid += item.Value.(int)
			}
			arr = append(arr, item.FieldName)
		case "SPids":
			if item.Type == 0 {
				aUser.User.SPids = item.Value.(string)
			} else {
				aUser.User.SPids += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "HeadImgurl":
			if item.Type == 0 {
				aUser.User.HeadImgurl = item.Value.(string)
			} else {
				aUser.User.HeadImgurl += item.Value.(string)
			}
			arr = append(arr, item.FieldName)

		case "UserName":
			if item.Type == 0 {
				aUser.User.UserName = item.Value.(string)
			} else {
				aUser.User.UserName += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Nickname":
			if item.Type == 0 {
				aUser.User.Nickname = item.Value.(string)
			} else {
				aUser.User.Nickname += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Pwd":
			if item.Type == 0 {
				aUser.User.Pwd = item.Value.(string)
			} else {
				aUser.User.Pwd += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "LowerCount":
			if item.Type == 0 {
				aUser.User.LowerCount = item.Value.(int)
			} else {
				aUser.User.LowerCount += item.Value.(int)
			}
			arr = append(arr, item.FieldName)
		case "AllLowerCount":
			if item.Type == 0 {
				aUser.User.AllLowerCount = item.Value.(int)
			} else {
				aUser.User.AllLowerCount += item.Value.(int)
			}
			arr = append(arr, item.FieldName)
		case "ReferrerCode":
			if item.Type == 0 {
				aUser.User.ReferrerCode = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.ReferrerCode += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "IsReferrer":
			if item.Type == 0 {
				aUser.User.IsReferrer = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.IsReferrer += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "State":
			if item.Type == 0 {
				aUser.User.State = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.State += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "Gold":
			if item.Type == 0 {
				aUser.User.Gold = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.Gold += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "SumXmBet":
			if item.Type == 0 {
				aUser.User.SumXmBet = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.SumXmBet += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "Sum2Rebate":
			if item.Type == 0 {
				aUser.User.Sum2Rebate = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.Sum2Rebate += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)

		case "Silver":
			if item.Type == 0 {
				aUser.User.Silver = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.Silver += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "Copper":
			if item.Type == 0 {
				aUser.User.Copper = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.Copper += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "SumSaveMoney":
			if item.Type == 0 {
				aUser.User.SumSaveMoney = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.SumSaveMoney += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "MaxSaveMoney":
			if item.Type == 0 {
				aUser.User.MaxSaveMoney = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.MaxSaveMoney += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "SumDrawMoney":
			if item.Type == 0 {
				aUser.User.SumDrawMoney = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.SumDrawMoney += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "MaxDrawMoney":
			if item.Type == 0 {

				aUser.User.MaxDrawMoney = kit.GetInterface2Float64(item.Value, 0)
			} else {
				aUser.User.MaxDrawMoney += kit.GetInterface2Float64(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "CreatedAt":
			if item.Type == 0 {
				aUser.User.CreatedAt = item.Value.(time.Time)
			}
			arr = append(arr, item.FieldName)
		case "LoginOutTime":
			if item.Type == 0 {
				aUser.User.LoginOutTime = item.Value.(time.Time)
			}
			arr = append(arr, item.FieldName)
		case "LoginTime":
			if item.Type == 0 {
				aUser.User.LoginTime = item.Value.(time.Time)
			}
			arr = append(arr, item.FieldName)
		case "VoucherFile":
			if item.Type == 0 {
				aUser.User.VoucherFile = item.Value.(string)
			} else {
				aUser.User.VoucherFile += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "WXSKCodeUrl":
			if item.Type == 0 {
				aUser.User.WXSKCodeUrl = item.Value.(string)
			} else {
				aUser.User.WXSKCodeUrl += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "YHName":
			if item.Type == 0 {
				aUser.User.YHName = item.Value.(string)
			} else {
				aUser.User.YHName += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "CardNum":
			if item.Type == 0 {
				aUser.User.CardNum = item.Value.(string)
			} else {
				aUser.User.CardNum += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "YHUserName":
			if item.Type == 0 {
				aUser.User.YHUserName = item.Value.(string)
			} else {
				aUser.User.YHUserName += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "YHUserTel":
			if item.Type == 0 {
				aUser.User.YHUserTel = item.Value.(string)
			} else {
				aUser.User.YHUserTel += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Addr":
			if item.Type == 0 {
				aUser.User.Addr = item.Value.(string)
			} else {
				aUser.User.Addr += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Cate":
			if item.Type == 0 {
				aUser.User.Cate = item.Value.(string)
			} else {
				aUser.User.Cate += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Remark":
			if item.Type == 0 {
				aUser.User.Remark = item.Value.(string)
			} else {
				aUser.User.Remark += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "ZFBSKCodeUrl":
			if item.Type == 0 {
				aUser.User.ZFBSKCodeUrl = item.Value.(string)
			} else {
				aUser.User.ZFBSKCodeUrl += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "ZFBSKName":
			if item.Type == 0 {
				aUser.User.ZFBSKName = item.Value.(string)
			} else {
				aUser.User.ZFBSKName += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Sid":
			if item.Type == 0 {
				aUser.User.Sid = item.Value.(string)
			} else {
				aUser.User.Sid += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "IdentityCard":
			if item.Type == 0 {
				aUser.User.IdentityCard = item.Value.(string)
			} else {
				aUser.User.IdentityCard += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "FullName":
			if item.Type == 0 {
				aUser.User.FullName = item.Value.(string)
			} else {
				aUser.User.FullName += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "Tel":
			if item.Type == 0 {
				aUser.User.Tel = item.Value.(string)
			} else {
				aUser.User.Tel += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "RealNameState":
			if item.Type == 0 {
				aUser.User.RealNameState = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.RealNameState += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "UserType":
			if item.Type == 0 {
				aUser.User.UserType = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.UserType += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)

		case "DrawMoneyPwd":
			if item.Type == 0 {
				aUser.User.DrawMoneyPwd = item.Value.(string)
			} else {
				aUser.User.DrawMoneyPwd += item.Value.(string)
			}
			arr = append(arr, item.FieldName)
		case "IsAgent":
			if item.Type == 0 {
				aUser.User.IsAgent = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.IsAgent += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		case "AgentUserId":
			if item.Type == 0 {
				aUser.User.AgentUserId = kit.GetInterface2Int(item.Value, 0)
			} else {
				aUser.User.AgentUserId += kit.GetInterface2Int(item.Value, 0)
			}
			arr = append(arr, item.FieldName)
		}
	}
	if len(arr) > 0 {
		aUser.User.Update(nil, arr...)
	}
	return getUserResult(aUser, e)
}

func Rebate2Gold(UserId int, Rebate float64, Des, Des2, DesMp string) (models.TtGameUser, error) {
	aUser, e := getUserAndStore(UserId)
	if e != nil {
		return getUserResult(aUser, e)
	}
	aUser.DataLock.Lock()
	defer aUser.DataLock.Unlock()

	if Rebate < 0 {
		return getUserResult(aUser, errors.New("参数值不对"))
	}
	if aUser.User.Rebate < Rebate {
		return getUserResult(aUser, errors.New("佣金不足"))
	}

	PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
	if e1 != nil {
		return getUserResult(aUser, fmt.Errorf("没有对应的代理商:%d", aUser.User.AgentUserId))
	}
	if PUser.User.Gold < Rebate {
		return getUserResult(aUser, fmt.Errorf("代理商[%d],余额不足：%g", aUser.User.Pid, Rebate))
	}
	PUser.User.Gold -= Rebate
	PUser.User.Gold = kit.Decimal(PUser.User.Gold)
	e = PUser.User.Update(nil, "Gold")
	if e != nil {
		return getUserResult(aUser, e)
	} else {
		if Rebate != 0 {
			aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_13_Rebate2),
				StrType:     mconst.GetAccountName(mconst.Account_13_Rebate2),
				Des:         fmt.Sprintf("用户%d[%s]'佣金转换'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, Rebate),
				CurUserGold: PUser.User.Gold,
				Gold:        Rebate,
				Des2:        fmt.Sprintf("用户%d[%s]'佣金转换'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, Rebate), DesMp: "",
			}
			aTtAccount.Add(nil)
		}
	}

	aUser.User.Rebate -= Rebate
	aUser.User.Gold += Rebate
	aUser.User.Gold = kit.Decimal(aUser.User.Gold)
	aUser.User.Rebate = kit.Decimal(aUser.User.Rebate)

	aUser.User.Update(nil, "Gold", "Rebate")

	aTtAccount := models.TtAccount{UserId: UserId, AccountType: int(mconst.Account_13_Rebate),
		StrType: mconst.GetAccountName(mconst.Account_13_Rebate), Des: Des, CurUserGold: aUser.User.Gold,
		Gold: Rebate, Des2: Des2, DesMp: DesMp,
	}
	aTtAccount.Add(nil)

	aTtRebateInfo := models.TtRebateInfo{UserId: UserId, RebateType: int(mconst.Rebate_02_ToGold),
		Gold:    -Rebate,
		StrType: mconst.GetRebateName(mconst.Rebate_02_ToGold), Des: mconst.GetRebateName(mconst.Rebate_02_ToGold), CurUserRebate: aUser.User.Rebate,
		Rebate: -Rebate,
	}
	aTtRebateInfo.Add(nil)

	return getUserResult(aUser, e)
}

func AddRebate(aAddRebateInfo gBox.AddRebateInfo) (models.TtGameUser, error) {
	addRebateLock.Lock()
	defer addRebateLock.Unlock()

	switch aAddRebateInfo.T {
	case mconst.Rebate_01_Guess:
		aRebateSetConfig := AreaConfig.GetRebateSet(0)
		addGuessRebate(aAddRebateInfo, aRebateSetConfig)
		return getUserResult(nil, nil)
	}

	return getUserResult(nil, nil)
}
func addGuessRebate(aAddRebateInfo gBox.AddRebateInfo, aRebateSetConfig AreaConfig.RebateSetConfig) {
	aUser, e := getUserAndStore(aAddRebateInfo.UserId)
	if e != nil {
		return
	}
	if aAddRebateInfo.Level > aRebateSetConfig.Level {
		return
	}
	aUser.DataLock.Lock()
	defer aUser.DataLock.Unlock()

	Rebate := strconvEx.Decimal(float64(aAddRebateInfo.BetM) * aRebateSetConfig.RebateRatio[aAddRebateInfo.Level])
	if Rebate == 0 {
		return
	}

	aUser.User.SumRebate += Rebate
	aUser.User.Rebate += Rebate
	aUser.User.Rebate = kit.Decimal(aUser.User.Rebate)
	aUser.User.SumRebate = kit.Decimal(aUser.User.SumRebate)
	aUser.User.Update(nil, "SumRebate", "Rebate")

	Des := fmt.Sprintf(`%s[%s][%s]%d,投注：%s %d元.返利比：%g`, game_mconst.GetGameName(aAddRebateInfo.GameType), aAddRebateInfo.LotteryStr,
		aAddRebateInfo.BetUserName, aAddRebateInfo.Level, aAddRebateInfo.OddsName, aAddRebateInfo.BetM, Rebate)

	Des2 := GTtHint.GetTtHint().GetHint(`%s[%s][%s]%d,投注：%s %d元.返利比：%g`)
	DesMp := GTtHint.GetTtHint().GetMpString(game_mconst.GetGameName(aAddRebateInfo.GameType), aAddRebateInfo.LotteryStr,
		aAddRebateInfo.BetUserName, aAddRebateInfo.Level, aAddRebateInfo.OddsName, aAddRebateInfo.BetM, Rebate)

	aTtRebateInfo := models.TtRebateInfo{UserId: aUser.User.Id, UserSid: aAddRebateInfo.UserSid, RebateType: int(aAddRebateInfo.T),
		StrType: mconst.GetRebateName(aAddRebateInfo.T), Gold: float64(aAddRebateInfo.BetM), CurUserRebate: aUser.User.Rebate, Rebate: Rebate,
		Des: Des, DesMp: DesMp, Des2: Des2}
	aTtRebateInfo.Add(nil)

	aAddRebateInfo.Level = aAddRebateInfo.Level + 1
	aAddRebateInfo.UserId = aUser.User.Pid

	addGuessRebate(aAddRebateInfo, aRebateSetConfig)
}
func getSaveMoneyResult(user *GUserInfo, aSaveMoney game_models.TtSaveMoney, e error) (models.TtGameUser, game_models.TtSaveMoney, error) {
	u, e := getUserResult(user, e)
	return u, aSaveMoney, e
}

func SaveMoney(info gBox.SaveMoneyInfo) (models.TtGameUser, game_models.TtSaveMoney, error) {
	aUser, e := getUserAndStore(info.UserId)
	if e != nil {
		return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e)
	}
	aUser.DataLock.Lock()
	defer aUser.DataLock.Unlock()

	switch info.InfoState {
	case gBox.SaveMoneyInfo_State_1_Apply:
		{
			curT := time.Now()
			aTtSaveMoney := game_models.TtSaveMoney{UserId: info.UserId, Money: info.Money, Gold: info.Money, State: mconst.SaveMoneyState_1_Apply,
				VoucherUrl: "", PayState: 0, AuditorId: 0, AuditorName: "", CreatedAt: curT, UpdatedAt: curT}
			aTtSaveMoney.Add(nil)
			return getSaveMoneyResult(aUser, aTtSaveMoney, nil)
		}
	case gBox.SaveMoneyInfo_State_2:
		{
			aSaveMoney := game_models.TtSaveMoney{Id: info.Id}
			o := orm.NewOrm()
			e := o.Read(&aSaveMoney)
			if e != nil {
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, errors.New(GTtHint.GetTtHint().GetHint("没有对应的数据")))
			}

			if info.UserId != 0 {
				if aSaveMoney.UserId != info.UserId {
					return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, errors.New(GTtHint.GetTtHint().GetHint("数据异常，用户信息不匹配")))
				}
			}

			if aSaveMoney.State >= info.State {
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, errors.New(GTtHint.GetTtHint().GetHint("数据异常,刷新后重试")))
			}

			aSaveMoney.State = info.State
			if aSaveMoney.State == mconst.SaveMoneyState_4 {
				aSaveMoney.VoucherUrl = info.VoucherUrl
				e = aSaveMoney.Update(o, "State", "UpdatedAt", "VoucherUrl")
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e)
			} else {
				e = aSaveMoney.Update(o, "State", "UpdatedAt")
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e)
			}
		}
	case gBox.SaveMoneyInfo_State_3_AgreeSaveMoney:
		{
			o := orm.NewOrm()

			aTtSaveMoney := game_models.TtSaveMoney{Id: info.Id}
			e1 := o.Read(&aTtSaveMoney)
			if e1 != nil {
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, errors.New("有问题的数据Id"))
			}

			if aTtSaveMoney.State == mconst.SaveMoneyState_5_OK || aTtSaveMoney.UserId != info.UserId {
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, errors.New("数据变动，请刷新数据"))
			}

			aTtSaveMoney.State = mconst.SaveMoneyState_5_OK
			aTtSaveMoney.AuditorId = info.AuditorId
			aTtSaveMoney.AuditorName = info.AuditorName

			e1 = aTtSaveMoney.Update(o, "State", "AuditorName", "AuditorId")
			if e1 != nil {
				return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e1)
			}

			PUser, e1 := getUserAndStore(aUser.User.AgentUserId)
			if e1 != nil {
				getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e1)
			}
			if PUser.User.Gold < info.Money {
				getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, fmt.Errorf("金额不足"))
			}
			PUser.User.Gold -= info.Money
			PUser.User.Gold = kit.Decimal(PUser.User.Gold)
			e = PUser.User.Update(nil, "Gold")
			if e != nil {
				getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e)
			} else {
				if info.Money != 0 {
					aTtAccount := models.TtAccount{UserId: PUser.User.Id, AccountType: int(mconst.Account_05_Give2),
						StrType:     mconst.GetAccountName(mconst.Account_05_Give2),
						Des:         fmt.Sprintf("用户%d[%s]'赠送'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, info.Money),
						CurUserGold: PUser.User.Gold,
						Gold:        info.Money,
						Des2:        fmt.Sprintf("用户%d[%s]'赠送'扣除代理商%d[%s]余额：%g", aUser.User.Id, aUser.User.UserName, PUser.User.Id, PUser.User.UserName, info.Money), DesMp: "",
					}
					aTtAccount.Add(nil)
				}
			}

			aUser.User.Gold += info.Money
			aUser.User.SumSaveMoney += info.Money
			if aUser.User.MaxSaveMoney < info.Money {
				aUser.User.MaxSaveMoney = info.Money

				aUser.User.MaxSaveMoney = kit.Decimal(aUser.User.MaxSaveMoney)
				aUser.User.SumSaveMoney = kit.Decimal(aUser.User.SumSaveMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				aUser.User.Update(nil, "Gold", "SumSaveMoney", "MaxSaveMoney")
			} else {
				aUser.User.SumSaveMoney = kit.Decimal(aUser.User.SumSaveMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				aUser.User.Update(nil, "Gold", "SumSaveMoney")
			}

			if info.Money != 0 {
				aTtAccount := models.TtAccount{UserId: info.UserId, AccountType: int(mconst.Account_03_SaveMoney),
					StrType: mconst.GetAccountName(mconst.Account_03_SaveMoney), Des: fmt.Sprintf("充值:%g", aTtSaveMoney.Gold),
					CurUserGold: aUser.User.Gold,
					Gold:        info.Money,
				}
				aTtAccount.Add(nil)
			}
		}

	}

	return getSaveMoneyResult(aUser, game_models.TtSaveMoney{}, e)
}

func DrawMoney(info gBox.DrawMoneyInfo) (models.TtGameUser, error) {
	aUser, e := getUserAndStore(info.UserId)
	if e != nil {
		return getUserResult(aUser, e)
	}
	aUser.DataLock.Lock()
	defer aUser.DataLock.Unlock()

	switch info.State {
	case mconst.DrawMoneyState_1_Apply:
		{
			if aUser.User.Gold < info.Money {
				return getUserResult(aUser, errors.New("余额不足"))
			}
			o := orm.NewOrm()
			aUser.User.Gold -= info.Money
			aUser.User.SumDrawMoney += info.Money
			if aUser.User.MaxDrawMoney < info.Money {
				aUser.User.MaxDrawMoney = info.Money

				aUser.User.MaxDrawMoney = kit.Decimal(aUser.User.MaxDrawMoney)
				aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				aUser.User.Update(o, "Gold", "SumDrawMoney", "MaxDrawMoney")
			} else {
				aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				aUser.User.Update(o, "Gold", "SumDrawMoney")
			}

			if info.Money != 0 {
				aTtAccount := models.TtAccount{UserId: info.UserId, AccountType: int(mconst.Account_04_DrawMoney),
					StrType: mconst.GetAccountName(mconst.Account_04_DrawMoney),
					Des:     fmt.Sprintf("提现申请扣除：%g。", info.Money), CurUserGold: aUser.User.Gold,
					Gold: info.Money, Des2: GTtHint.GetTtHint().GetHint("提现申请扣除：%g。"),
					DesMp: GTtHint.GetTtHint().GetMpString(info.Money),
				}
				aTtAccount.Add(o)
			}
			curT := time.Now()
			aDrawMoney := game_models.TtDrawMoney{
				UserId:      info.UserId,
				Gold:        info.Money,
				State:       info.State,
				AuditorId:   info.AuditorId,
				AuditorName: info.AuditorName,
				CreatedAt:   curT, UpdatedAt: curT,
			}
			aDrawMoney.Add(o)

		}
	case mconst.DrawMoneyState_2:
		{
			o := orm.NewOrm()
			goldInfo := gBox.AddGoldInfo{GroupId: info.GroupId, UserId: info.UserId, Gold: info.Money, T: mconst.Account_09_DecMoney,
				Des:  fmt.Sprintf("[%s]下分%g。", info.AuditorName, info.Money),
				Des2: fmt.Sprintf("[%s]Rút điểm%g。", info.AuditorName, info.Money),
			}

			//提现
			if aUser.User.Gold < goldInfo.Gold {
				return getUserResult(aUser, errors.New("余额不足"))
			}
			aUser.User.Gold -= goldInfo.Gold
			aUser.User.SumDrawMoney += goldInfo.Gold
			if aUser.User.MaxDrawMoney < goldInfo.Gold {
				aUser.User.MaxDrawMoney = kit.Decimal(goldInfo.Gold)
				aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				e := aUser.User.Update(o, "Gold", "SumDrawMoney", "MaxDrawMoney")
				if e != nil {
					ttLog.LogError(e)
				}
			} else {
				aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				e := aUser.User.Update(o, "Gold", "SumDrawMoney")

				if e != nil {
					ttLog.LogError(e)
				}
			}

			if goldInfo.Gold != 0 {
				aTtAccount := models.TtAccount{UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
					StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
					Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
				}
				e := aTtAccount.Add(o)
				if e != nil {
					ttLog.LogError(e)
				}
			}
			aTtDrawMoney := game_models.TtDrawMoney{UserId: info.UserId, Gold: info.Money, State: info.State,
				AuditorId: info.AuditorId, AuditorName: info.AuditorName, OrderId: info.OrderId}
			e := aTtDrawMoney.Add(o)
			if e != nil {
				ttLog.LogError(e)
			}

		}
	case mconst.DrawMoneyState_5:
		{
			if info.Id == 0 {
				o := orm.NewOrm()
				if info.OrderId == "" {
					return getUserResult(aUser, errors.New("订单号不为空"))
				}

				arrDrawMoney := make([]game_models.TtDrawMoney, 0)
				o.QueryTable(game_mconst.TableName_TtDrawMoney).Filter("OrderId", info.OrderId).Filter("UserId", info.UserId).All(&arrDrawMoney)
				for _, aDrawMoney := range arrDrawMoney {
					aDrawMoney.State = mconst.DrawMoneyState_5
					aDrawMoney.Update(o, "State", "UpdatedAt")

					goldInfo := gBox.AddGoldInfo{GroupId: aDrawMoney.GroupId, UserId: aDrawMoney.UserId, Gold: aDrawMoney.Gold, T: mconst.Account_07_DrawMoneyR,
						Des: fmt.Sprintf("[%s]提现拒绝%g。", aDrawMoney.AuditorName, aDrawMoney.Gold),
					}
					//提现拒绝
					aUser.User.Gold += goldInfo.Gold

					aUser.User.SumDrawMoney -= goldInfo.Gold
					aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
					aUser.User.Gold = kit.Decimal(aUser.User.Gold)
					aUser.User.Update(nil, "Gold", "SumDrawMoney")

					if goldInfo.Gold != 0 {
						aTtAccount := models.TtAccount{GroupId: goldInfo.GroupId, UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
							StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
							Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
						}
						aTtAccount.Add(nil)
					}
				}
			} else {
				o := orm.NewOrm()
				aDrawMoney := game_models.TtDrawMoney{Id: info.Id}
				e := o.Read(&aDrawMoney)
				if e != nil {
					return getUserResult(aUser, errors.New("没有对应的ID"))
				}
				aDrawMoney.State = mconst.DrawMoneyState_5
				aDrawMoney.Update(o, "State", "UpdatedAt")

				goldInfo := gBox.AddGoldInfo{GroupId: aDrawMoney.GroupId, UserId: aDrawMoney.UserId, Gold: aDrawMoney.Gold, T: mconst.Account_07_DrawMoneyR,
					Des:  fmt.Sprintf("[%s]下分%g。", aDrawMoney.AuditorName, aDrawMoney.Gold),
					Des2: fmt.Sprintf("[%s]Rút điểm%g。", aDrawMoney.AuditorName, aDrawMoney.Gold),
				}
				aUser.User.Gold += goldInfo.Gold

				aUser.User.SumDrawMoney -= goldInfo.Gold
				aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				aUser.User.Update(nil, "Gold", "SumDrawMoney")

				if goldInfo.Gold != 0 {
					aTtAccount := models.TtAccount{GroupId: goldInfo.GroupId, UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
						StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
						Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
					}
					aTtAccount.Add(nil)
				}

			}

		}
	case mconst.DrawMoneyState_6:
		{
			if info.Id == 0 {
				o := orm.NewOrm()
				if info.OrderId == "" {
					return getUserResult(aUser, errors.New("订单号不为空"))
				}

				arrDrawMoney := make([]game_models.TtDrawMoney, 0)
				o.QueryTable(game_mconst.TableName_TtDrawMoney).Filter("OrderId", info.OrderId).Filter("UserId", info.UserId).All(&arrDrawMoney)
				for _, aDrawMoney := range arrDrawMoney {
					aDrawMoney.State = mconst.DrawMoneyState_5
					aDrawMoney.Update(o, "State", "UpdatedAt")

					goldInfo := gBox.AddGoldInfo{GroupId: aDrawMoney.GroupId, UserId: aDrawMoney.UserId, Gold: aDrawMoney.Gold, T: mconst.Account_07_DrawMoneyR,
						Des: fmt.Sprintf("[%s]下分拒绝%g。", aDrawMoney.AuditorName, aDrawMoney.Gold),
					}
					//提现拒绝
					aUser.User.Gold += goldInfo.Gold

					aUser.User.SumDrawMoney -= goldInfo.Gold
					aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
					aUser.User.Gold = kit.Decimal(aUser.User.Gold)
					aUser.User.Update(nil, "Gold", "SumDrawMoney")

					if goldInfo.Gold != 0 {
						aTtAccount := models.TtAccount{GroupId: goldInfo.GroupId, UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
							StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
							Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
						}
						aTtAccount.Add(nil)
					}
				}
			} else {
				o := orm.NewOrm()
				aDrawMoney := game_models.TtDrawMoney{Id: info.Id}
				e := o.Read(&aDrawMoney)
				if e != nil {
					return getUserResult(aUser, errors.New("没有对应的ID"))
				}
				aDrawMoney.State = mconst.DrawMoneyState_5
				aDrawMoney.Update(o, "State", "UpdatedAt")

				goldInfo := gBox.AddGoldInfo{GroupId: aDrawMoney.GroupId, UserId: aDrawMoney.UserId, Gold: aDrawMoney.Gold, T: mconst.Account_07_DrawMoneyR,
					Des:  fmt.Sprintf("[%s]下分%g。", aDrawMoney.AuditorName, aDrawMoney.Gold),
					Des2: fmt.Sprintf("[%s]Rút điểm%g。", aDrawMoney.AuditorName, aDrawMoney.Gold),
				}
				aUser.User.Gold += goldInfo.Gold

				aUser.User.SumDrawMoney -= goldInfo.Gold
				aUser.User.SumDrawMoney = kit.Decimal(aUser.User.SumDrawMoney)
				aUser.User.Gold = kit.Decimal(aUser.User.Gold)
				aUser.User.Update(nil, "Gold", "SumDrawMoney")

				if goldInfo.Gold != 0 {
					aTtAccount := models.TtAccount{GroupId: goldInfo.GroupId, UserId: goldInfo.UserId, AccountType: int(goldInfo.T),
						StrType: mconst.GetAccountName(goldInfo.T), Des: goldInfo.Des, CurUserGold: aUser.User.Gold,
						Gold: goldInfo.Gold, Des2: goldInfo.Des2, DesMp: goldInfo.DesMp,
					}
					aTtAccount.Add(nil)
				}

			}

		}
	case mconst.DrawMoneyState_3:
		{

		}
	case mconst.DrawMoneyState_4:
		{
			if info.Id == 0 {
				o := orm.NewOrm()
				if info.OrderId == "" {
					return getUserResult(aUser, errors.New("订单号不为空"))
				}
				arrDrawMoney := make([]game_models.TtDrawMoney, 0)
				o.QueryTable(game_mconst.TableName_TtDrawMoney).Filter("OrderId", info.OrderId).All(&arrDrawMoney)
				for _, aDrawMoney := range arrDrawMoney {
					aDrawMoney.State = info.State
					aDrawMoney.Update(o, "State", "UpdatedAt")
				}
			} else {
				o := orm.NewOrm()
				aTtDrawMoney := game_models.TtDrawMoney{Id: info.Id}
				aTtDrawMoney.State = info.State
				aTtDrawMoney.Update(o, "UpdatedAt", "State")
			}
		}
	default:
		return getUserResult(aUser, errors.New("有问题的状态"))
	}

	return getUserResult(aUser, e)
}
