package models

import (
	"github.com/TtMyth123/Admin/models/mconst"
	"time"
)

type ChatInfo struct {
	Id        int64
	UserName  string    `orm:"description(用户名)"`  //用户名
	GameId    int64     `orm:"description(游戏ID)"` //游戏ID
	Way       int       `orm:"description(方向,0)"`
	ChatType  int       `orm:"description(聊天类型.0:一般聊天内容)"` //聊天类型.0:一般聊天内容
	Content   string    `orm:"description(聊天内容)"`          //聊天内容
	CreatedAt time.Time `orm:"auto_now_add;type(datetime);description(时间)"`
}

func (a *ChatInfo) TableName() string {
	return mconst.TableName_SysMenu
}
