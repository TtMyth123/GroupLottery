package mconst

const (
	TableName_TtGameUser   = "tt_game_user"
	TableName_TtAccount    = "tt_account"
	TableName_TtRebateInfo = "tt_rebate_info"

	TableName_TtArea = "tt_area"

	TableName_TtAreaRefGame = "tt_area_ref_game"

	TableName_TtChatInfo       = "tt_chat_info"
	TableName_TtLatelyChatUser = "tt_lately_chat_user"
)

type aa struct {
	TtRebateInfo     int `json:"tt_rebate_info"`
	TtChatInfo       int `json:"tt_chat_info"`
	TtLatelyChatUser int `json:"tt_lately_chat_user"`
}
