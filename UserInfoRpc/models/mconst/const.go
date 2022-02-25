package mconst

type Account_Type int

const (
	Account_01_Guess   Account_Type = 1  //竞猜
	Account_02_Win     Account_Type = 2  //赢得
	Account_14_NotOpen Account_Type = 14 //不开奖退款
	//Account_10_SubGuessRebate Account_Type = 10 //下级竞猜佣金(上级)
	//Account_11_SubGuessRebate Account_Type = 11 //下级竞猜佣金(下级)
	Account_03_SaveMoney   Account_Type = 3   //充值
	Account_03_SaveMoney2  Account_Type = 103 //充值
	Account_04_DrawMoney   Account_Type = 4   //提现
	Account_04_DrawMoney2  Account_Type = 104 //提现
	Account_07_DrawMoneyR  Account_Type = 7   //"提现拒绝"
	Account_07_DrawMoneyR2 Account_Type = 107 //"提现拒绝"
	Account_08_AddMoney    Account_Type = 8   //上分
	Account_08_AddMoney2   Account_Type = 108 //上分
	Account_09_DecMoney    Account_Type = 9   //下分
	Account_09_DecMoney2   Account_Type = 109 //下分
	Account_05_Give        Account_Type = 5   //赠送
	Account_05_Give2       Account_Type = 105 //赠送
	//Account_12_XmGuess        Account_Type = 12 //洗码
	Account_13_Rebate     Account_Type = 13  //佣金转换
	Account_13_Rebate2    Account_Type = 113 //佣金转换
	Account_14_DrawMoneyR Account_Type = 14  //"拒绝"
)

const (
	Account_N_01_Guess          = "竞猜"
	Account_N_02_Win            = "赢得"
	Account_N_10_SubGuessRebate = "下级竞猜佣金(上级)"
	Account_N_11_SubGuessRebate = "下级竞猜佣金(下级)"
	Account_N_03_SaveMoney      = "充值"
	Account_N_03_SaveMoney2     = "下级充值"
	Account_N_04_DrawMoney      = "提现"
	Account_N_04_DrawMoney2     = "下级提现"
	Account_N_07_DrawMoneyR     = "提现拒绝"
	Account_N_07_DrawMoneyR2    = "下级提现拒绝"
	Account_N_08_AddMoney       = "上分"
	Account_N_08_AddMoney2      = "下级上分"
	Account_N_09_DecMoney       = "下分"
	Account_N_09_DecMoney2      = "下级下分"
	Account_N_05_Give           = "赠送"
	Account_N_05_Give2          = "下级赠送"
	Account_N_12_XmGuess        = "洗码"
	Account_N_13_Rebate         = "佣金转换"
	Account_N_13_Rebate2        = "下级佣金转换"
)

const (
	UserType_1 = 1 //一般玩家
	UserType_2 = 2 //业务玩家
	UserType_3 = 3 //游客
	UserType_4 = 4 //客服
)
const RootGameUserId = 1

const (
	UserType_N_1 = "一般玩家"
	UserType_N_2 = "业务玩家"
	UserType_N_3 = "游客"
)

const (
	User_State_0_New = 0 //新加
	User_State_1     = 1 //启动
	User_State_2     = 2 //禁用
)
