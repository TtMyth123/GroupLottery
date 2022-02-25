package chat

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

type ReplyManage struct {
	Msg      map[int64]map[string][]string
	userLock sync.Mutex
}

var (
	mReplyManage *ReplyManage
)

func Init() {
	mReplyManage = new(ReplyManage)
	mReplyManage.ReloadAll()
}

func GetReplyManage() *ReplyManage {
	return mReplyManage
}
func (this *ReplyManage) ReloadAll() {
	this.userLock.Lock()
	defer this.userLock.Unlock()
	this.Msg = make(map[int64]map[string][]string)

	o := orm.NewOrm()
	type Tmp struct {
		Id           int64
		StaffId      int64
		MainKey      string
		ReplyContent string
	}
	arrData := make([]Tmp, 0)
	sql := `select a.id, a.main_key,a.staff_id, b.id as sub_id, b.reply_content
from tt_reply_main a, tt_reply_sub b where a.id=b.main_id `

	o.Raw(sql).QueryRows(&arrData)
	for _, v := range arrData {
		staffMap := this.Msg[v.StaffId]
		if staffMap == nil {
			staffMap = make(map[string][]string)
		}
		arrReply := staffMap[v.MainKey]
		if arrReply == nil {
			arrReply = make([]string, 0)
		}
		arrReply = append(arrReply, v.ReplyContent)
		staffMap[v.MainKey] = arrReply

		this.Msg[v.StaffId] = staffMap
	}

}

func (this *ReplyManage) Reload(adminId int64) {
	this.userLock.Lock()
	defer this.userLock.Unlock()
	this.Msg = make(map[int64]map[string][]string)

	o := orm.NewOrm()
	type Tmp struct {
		Id           int
		StaffId      int64
		MainKey      string
		ReplyContent string
	}
	arrData := make([]Tmp, 0)
	sql := `select a.id, a.main_key,a.staff_id, b.id as sub_id, b.reply_content
from tt_reply_main a, tt_reply_sub b where a.id=b.main_id and a.staff_id=?`
	this.Msg[adminId] = make(map[string][]string)

	o.Raw(sql, adminId).QueryRows(&arrData)
	for _, v := range arrData {
		staffMap := this.Msg[v.StaffId]
		if staffMap == nil {
			staffMap = make(map[string][]string)
		}
		arrReply := staffMap[v.MainKey]
		if arrReply == nil {
			arrReply = make([]string, 0)
		}
		arrReply = append(arrReply, v.ReplyContent)
		staffMap[v.MainKey] = arrReply

		this.Msg[v.StaffId] = staffMap
	}

}
func (this *ReplyManage) GetQuickReply(adminId int64, c string) []string {
	this.userLock.Lock()
	defer this.userLock.Unlock()

	staffMap := this.Msg[adminId]
	if staffMap == nil {
		return []string{}
	} else {
		return staffMap[c]
	}
}
