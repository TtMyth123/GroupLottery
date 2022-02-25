package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ChatList struct {
	Id         int64     `orm:"column(id)" description:"ID"`
	TableName1 string    `orm:"column(table_name);size(256)" description:"对应聊天表名"`
	Sign       string    `orm:"column(sign);size(256)" description:"Sign"`
	RoomId     int       `orm:"column(room_id)" description:"房间聊天ID"`
	RoomName   string    `orm:"column(room_name);size(256)" description:"房间名"`
	UserCount  int       `orm:"column(user_count)" description:"人数"`
	UserId     string    `orm:"column(user_id);size(4000)" description:"房间用户ID"`
	CreatedAt  time.Time `orm:"column(created_at); auto_now_add;type(datetime)" description:"创建时间"`
}

func (a *ChatList) TableName() string {
	return "ChatList"
}

func (this *ChatList) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.CreatedAt = time.Now()
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *ChatList) Update(o orm.Ormer, cols ...string) error {
	_, e := o.Update(this, cols...)
	return e
}

func (this *ChatList) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(&this)
	return e
}
