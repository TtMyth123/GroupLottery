package mconst

const (
	TableName_TtReplyMain  = "tt_reply_main"
	TableName_TtReplySub   = "tt_reply_sub"
	TableName_TtQuickReply = "tt_quick_reply"
)

type a struct {
	TtReplySub   int    `json:"tt_reply_sub"`
	TtQuickReply string `json:"tt_quick_reply"`
}
