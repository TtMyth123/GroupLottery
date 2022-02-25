package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/UserInfoRpc/GConfig"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
	"ttmyth123/kit/httpKit"
	"ttmyth123/kit/timeKit"
)

type TtChatInfo struct {
	Id  int64
	Id2 int64 `orm:"" description:"房间聊天ID"`
	//RoomChatId int    `orm:"" description:"房间聊天ID"`
	RoomId    int       `orm:"" description:"房间ID"`
	Letter    string    `orm:"size(15)" description:"姓"`     //姓
	RoomName  string    `orm:"size(256)" description:"房间名"`  //用户名
	GameId    int64     `orm:"" description:"(游戏ID)"`        //游戏ID
	UserName  string    `orm:"size(256)" description:"用户名"`  //用户名
	GameId2   int64     `orm:"" description:"(游戏ID2)"`       //游戏ID
	UserName2 string    `orm:"size(256)" description:"用户名2"` //用户名
	Way       int       `orm:"" description:"(方向,0:客户端消息)"`
	ChatType  int       `orm:"" description:"(聊天类型.0:一般聊天内容)"`  //聊天类型.0:一般聊天内容
	Content   string    `orm:"size(4000)" description:"(聊天内容)"` //聊天内容
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	UpdateAt  time.Time `orm:"auto_now_add;type(datetime)" description:"更新时间"`
}

func (d *TtChatInfo) MarshalJSON() ([]byte, error) {
	type Alias TtChatInfo

	aContent := d.Content
	if d.ChatType == mconst.ChatType_Pic_1 {
		aContent = httpKit.GetImgUrl(GConfig.GetGConfig().BasePicRootPath, d.Content)
	}

	StrTime1 := ""
	StrTime2 := ""
	TimeLag := GConfig.GetGConfig().TimeLag
	if TimeLag == 0 {
		StrTime1 = d.CreatedAt.Format(timeKit.DateTimeLayout)
		StrTime2 = d.UpdateAt.Format(timeKit.DateTimeLayout)
	} else {
		StrTime1 = d.CreatedAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
		StrTime2 = d.UpdateAt.Add(TimeLag * time.Hour).Format(timeKit.DateTimeLayout)
	}

	return json.Marshal(&struct {
		Alias
		CreatedAt string
		UpdateAt  string
		Content   string
	}{
		Alias:     (Alias)(*d),
		CreatedAt: StrTime1,
		UpdateAt:  StrTime2,
		Content:   aContent,
	})
}

func (a *TtChatInfo) TableName() string {
	return mconst.TableName_TtChatInfo
}

func (this *TtChatInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	runeUserName := []rune(this.RoomName)
	Letter := ""
	if len(runeUserName) > 0 {
		Letter = string(runeUserName[0:1])
	}
	this.Letter = Letter
	this.CreatedAt = time.Now()
	this.UpdateAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *TtChatInfo) Update(o orm.Ormer, cols ...string) error {
	_, e := o.Update(this, cols...)
	this.UpdateAt = time.Now()
	return e
}

func (this *TtChatInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}
