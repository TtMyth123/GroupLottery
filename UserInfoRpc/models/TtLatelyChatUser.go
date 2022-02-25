package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"ttmyth123/GroupLottery/UserInfoRpc/models/mconst"
)

type TtLatelyChatUser struct {
	Id       int64
	RoomId   int    `orm:"" description:"房间ID"`          //用户名
	RoomName string `orm:"size(256)" description:"房间名"`  //房间名
	Letter   string `orm:"size(16)" description:"姓"`     //姓
	UserName string `orm:"size(256)" description:"用户名"`  //用户名
	Sign     string `orm:"size(256)" description:"房间标示"` //房间标示
	Top      int    `orm:"" description:"置顶"`
	Status   int    `orm:"" description:"状态,0:一般,1:已解散"`

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	UpdateAt  time.Time `orm:"auto_now_add;type(datetime)" description:"更新时间"`
}

func (a *TtLatelyChatUser) TableName() string {
	return mconst.TableName_TtLatelyChatUser
}

func (this *TtLatelyChatUser) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	runeRoomName := []rune(this.RoomName)
	Letter := ""
	if len(runeRoomName) > 0 {
		Letter = string(runeRoomName[0:1])
	}
	this.CreatedAt = time.Now()
	this.UpdateAt = this.CreatedAt
	this.Letter = Letter

	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *TtLatelyChatUser) Update(o orm.Ormer, cols ...string) error {
	this.UpdateAt = time.Now()
	_, e := o.Update(this, cols...)
	return e
}

func (this *TtLatelyChatUser) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}
