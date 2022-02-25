package mconst

const (
	TableName_LoOddsInfo     = "lo_odds_info"
	TableName_LoBetGroupInfo = "lo_bet_group_info"
	TableName_LoBetInfo      = "lo_bet_info"
	TableName_LoAwardInfo    = "lo_award_info"
	TableName_LoSetAwardInfo = "lo_set_award_info"

	TableName_LoUserInfo2 = "lo_user_info2"

	TableName_TtSaveMoney   = "tt_save_money"
	TableName_TtDrawMoney   = "tt_draw_money"
	TableName_TtDrawSaveSet = "tt_draw_save_set"
	TableName_TtRebateSet   = "tt_rebate_set"

	TableName_TtGameInfo     = "tt_game_info"
	TableName_TtArticleInfo  = "tt_article_info"
	TableName_TtServiceInfo  = "tt_service_info"
	TableName_TtAgentPayInfo = "tt_agent_pay_info"
)

const (
	Bet_Status_1 = 1 //待开奖
	Bet_Status_2 = 2 //已兑奖
	Bet_Status_3 = 3 //不开奖退款
)
const (
	Bet_Status_N_1 = "待开奖"
	Bet_Status_N_2 = "已兑奖"
	Bet_Status_N_3 = "不开奖退款"
)

type a struct {
	TtRebateSet   int `json:"tt_rebate_set"`
	TtServiceInfo int `json:"tt_service_info"`
	TtDrawSaveSet int `json:"tt_draw_save_set"`
	TtDrawMoney   int `json:"tt_draw_money"`
	TtSaveMoney   int `json:"tt_save_money"`
	LoAwardInfo   int `json:"lo_award_info"`
}

/**
提现状态
*/
const (
	DrawMoneyState_1_Apply = 1 //申请状态
	DrawMoneyState_2       = 2 //下分
	DrawMoneyState_3       = 3 //下分完成
	DrawMoneyState_4       = 4 //下分完成——已汇款
)

/**
充值状态
*/
const (
	SaveMoneyState_1_Apply = 1 //申请状态
	SaveMoneyState_2       = 2 //付款中
	SaveMoneyState_3       = 3 //已付款
	SaveMoneyState_4       = 4 //已上传凭证
	SaveMoneyState_5_OK    = 5 //充值成功
)

const (
	RegGiveType_1 = 1 //金币
	RegGiveType_2 = 2 //积分
	RegGiveType_3 = 3 //转盘
)
const (
	RegGiveType_N_1 = "金币"
	RegGiveType_N_2 = "积分"
	RegGiveType_N_3 = "转盘"
)

const (
	SaveGiveType_1 = 1 //金币
	SaveGiveType_2 = 2 //积分
	SaveGiveType_3 = 3 //转盘
)

const (
	SaveGiveType_N_1 = "金币"
	SaveGiveType_N_2 = "积分"
	SaveGiveType_N_3 = "转盘"
)

const (
	ArticleState_1_New     = 1 //新建
	ArticleState_2_Enabled = 2 //启用
	ArticleState_3_Disable = 3 //停用
)

const (
	ArticleType_1_GG  = 1 //公告
	ArticleType_2_XW  = 2 //新闻
	ArticleType_3_WZ  = 3 //文章
	ArticleType_4_AD  = 4 //广告 (后台会控制只有三个)
	ArticleType_5_HDP = 5 //幻灯片

	ArticleType_11_QYJJ = 11 //企业简历
	ArticleType_12_QYZZ = 12 //企业资质
	ArticleType_13_QYJS = 13 //企业介绍
	ArticleType_14_QYXC = 14 //企业宣传
	ArticleType_15_Help = 15 //帮助
)
const (
	ArticleType_N_1_GG    = "公告"
	ArticleType_N_2_XW    = "新闻"
	ArticleType_N_3_WZ    = "文章"
	ArticleType_N_11_QYJJ = "企业简历"
	ArticleType_N_12_QYZZ = "企业资质"
	ArticleType_N_13_QYJS = "企业介绍"
	ArticleType_N_14_QYXC = "企业宣传"
	ArticleType_N_15_Help = "帮助"
)
